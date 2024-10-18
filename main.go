package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.86.58"
	port     = 5432
	user     = "pi"
	password = "Boomer2025"
	dbname   = "test"
)

var (
	street string
	city   string
	state  string
	zip    string
	lat    string
	lng    string
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

	//--------------------------------------------------------------
	//--- POSTGRES CONNECTION
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//--------------------------------------------------------------

	//--------------------------------------------------------------
	//--- THIS WILL BRING BACK ROWS THAT NEED TO BE GEOCODED.
	//--- IN THE FUTURE, LETS LOOK FOR ROWS WITH NULL LAT AND LNG
	rows, err := db.Query("SELECT street, city, state, zip  FROM address order by address_id desc")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//--------------------------------------------------------------

	for rows.Next() {
		err := rows.Scan(&street, &city, &state, &zip)
		if err != nil {
			panic(err)
		}
		full_address := street + " " + city + " " + state + " " + zip

		//-----------------------------------------------------------------------
		//--- GEOCODE HERE
		arcgisUrl := "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates?f=pjson&SingleLine=" + (url.QueryEscape(full_address)) + "&outFields=x,y"

		resp, err := http.Get(arcgisUrl)

		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var result arcgisResults

		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		fmt.Println("---------------------")
		fmt.Println("Address:" + full_address)
		fmt.Println("Lng: " + fmt.Sprintf("%f", result.Candidates[0].Location.X))
		fmt.Println("Lat: " + fmt.Sprintf("%f", result.Candidates[0].Location.Y))
		fmt.Println("---------------------")

		//var result arcgisResults

		// fmt.Println(arcgisUrl)
		// resp, err := http.Get(arcgisUrl)
		// fmt.Println(resp.Body)

		// if err != nil {
		// 	fmt.Println("No response from request")
		// }

		// //--- Load all the JSON as text into "body"
		// defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body) // response body is []byte

		// //--- Unmarshal the JSON into the "result" value
		// var result arcgisResults

		// //fmt.Println(result)
		// if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer

		// 	fmt.Println("Can not unmarshal JSON")
		// }
		// //-----------------------------------------------------------------------

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
