package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "192.168.86.58"
// 	port     = 5432
// 	user     = "pi"
// 	password = "Boomer2025"
// 	dbname   = "test"
// )

var (
	address_id int
	street     string
	city       string
	state      string
	zip        string
	lat        string
	lng        string
)

type arcgisResults struct {
	SpatialReference struct {
		Wkid       int `json:"wkid"`
		LatestWkid int `json:"latestWkid"`
	} `json:"spatialReference"`
	Candidates []struct {
		Address  string `json:"address"`
		Location struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"location"`
		Score      float64 `json:"score"`
		Attributes struct {
		} `json:"attributes"`
		// Extent struct {
		// 	Xmin float64 `json:"xmin"`
		// 	Ymin float64 `json:"ymin"`
		// 	Xmax float64 `json:"xmax"`
		// 	Ymax float64 `json:"ymax"`
		// } `json:"extent"`
	} `json:"candidates"`
}

func main() {

	//--- Get the CREDENTIALS from the env file.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	//--- Make the POSTGRES connection.
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	//log.Println("Made connection to " + host)
	if err != nil {
		//log.Println("Connection to " + host + " failed!")
		//log.Println(err)
		panic(err)
	}
	defer db.Close()

	//--- Load up a ROWS object with records that need to be geocoded.
	//--- MT: Should this be a custom function in the database?
	var sbSql strings.Builder
	sbSql.WriteString("SELECT ")
	sbSql.WriteString("		address_id, street, city, state, zip ")
	sbSql.WriteString("FROM ")
	sbSql.WriteString("		address ")
	sbSql.WriteString("WHERE ")
	sbSql.WriteString("		(lat is null and lng is null) ")
	sbSql.WriteString("order by address_id desc ")

	strSql := sbSql.String()
	rows, err := db.Query(strSql)
	if err != nil {
		//log.Println("Query failed!")
		//log.Printin(strSQL)
		//log.Println(err)
		panic(err)
	}
	defer rows.Close()

	//--- Begin a loop through all ROWS in the row object.
	for rows.Next() {
		err := rows.Scan(&address_id, &street, &city, &state, &zip)
		if err != nil {
			panic(err)
		}
		//--- Create a full_address string
		full_address := street + " " + city + " " + state + " " + zip

		//--- Now let's craft a proper URL to send to ArcGIS
		var sbArcGis strings.Builder
		sbArcGis.WriteString("https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/")
		sbArcGis.WriteString("findAddressCandidates?f=pjson&SingleLine=" + (url.QueryEscape(full_address)))
		sbArcGis.WriteString("&outFields=x,y")
		arcgisUrl := sbArcGis.String()

		//--- Fire off the API request.
		resp, err := http.Get(arcgisUrl)

		//--- Do a little dance here to get the results into the arcgisResults object
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var result arcgisResults

		//--- Error handling, please.
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		//--- Get the lat and lng from the response
		lat := fmt.Sprintf("%f", result.Candidates[0].Location.Y)
		lng := fmt.Sprintf("%f", result.Candidates[0].Location.X)

		//--- Build an UPDATE statment for each record that needs it.
		//--- MT: Also a store proc in the database?
		var sbSqlUpdate strings.Builder
		sbSqlUpdate.WriteString("UPDATE address SET ")
		sbSqlUpdate.WriteString("lat=" + lat)
		sbSqlUpdate.WriteString(", lng=" + lng)
		sbSqlUpdate.WriteString(" WHERE address_id = " + strconv.Itoa(address_id))

		//--- Run the UPDATEs against the database.
		fmt.Println(sbSqlUpdate.String())

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
