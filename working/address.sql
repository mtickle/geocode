--
-- PostgreSQL database dump
--

-- Dumped from database version 15.8 (Debian 15.8-0+deb12u1)
-- Dumped by pg_dump version 16.0

-- Started on 2024-10-23 06:47:01

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3343 (class 0 OID 16391)
-- Dependencies: 215
-- Data for Name: address; Type: TABLE DATA; Schema: public; Owner: pi
--

COPY public.address (address_id, street, city, state, zip, lat, lng) FROM stdin;
59	202 HARLOW ST	Bangor	ME	04401	\N	\N
86	2276 WILTON DR	Fort Lauderdale	FL	33305	\N	\N
87	203 SOUTH WALNUT ST	Florence	AL	35630	\N	\N
88	108 CENTER POINTE DR	Clarksville	TN	37040	\N	\N
89	1800 OLD TROY RD	Union City	TN	38261	\N	\N
90	931 OLD SMITHVILLE HWY	McMinnville	TN	37110	\N	\N
91	1301 GREENE STREET	Marietta	OH	45750	\N	\N
92	602 SOUTH MICHIGAN ST	South Bend	IN	46601	\N	\N
93	500 NORTH A STREET	Richmond	IN	47374	\N	\N
94	317 SOUTH DRAKE ROAD	Kalamazoo	MI	49009	\N	\N
95	105 Amity Way	Wayne	PA	19087	\N	\N
96	305 W 45th St	New York	NY	10036	\N	\N
97	11839 Federalist Way	Fairfax	VA	22030	\N	\N
98	400 Monroe St	Hoboken	NJ	07030	\N	\N
99	101 West End Avenue	New York	NY	10023	\N	\N
100	2900 4TH AVE	Billings	MT	59101	\N	\N
101	158 N SCOTT STREET	Joliet	IL	60432	\N	\N
102	1207 NETWORK CENTRE DR	Effingham	IL	62401	\N	\N
103	3555 SOUTHERN HILLS DR6	Sioux City	IA	51106	\N	\N
106	415 VALLEY VIEW DR	Scottsbluff	NE	69361	\N	\N
4	1010 Jones Rd	Little Town	NC	27272	\N	\N
81	145 ROWAN STREET	Fayetteville	NC	28301	\N	\N
82	1420 MCCARTHY BLVD	New Bern	NC	28562	\N	\N
83	115 ENTERPRISE COURT	Greenwood	SC	29649	\N	\N
85	97 WEST OAK AVE	Panama City	FL	32401	\N	\N
104	300 E 3RD ST	North Platte	NE	69101	\N	\N
105	115 N WEBB RD	Grand Island	NE	68803	\N	\N
1	5050 Saturn Dr	Rapid City	SD	57701	\N	\N
2	4325 Grassy Field Drive	Raleigh	NC	27610	\N	\N
3	302 N. Pine Street	Wendell	NC	27591	\N	\N
56	93 NORTH 9TH STREET	Brooklyn	NY	11211	\N	\N
57	380 WESTMINSTER ST	Providence	RI	02903	\N	\N
58	177 MAIN STREET	Littleton	NH	03561	\N	\N
60	46 FRONT STREET	Waterville	ME	04901	\N	\N
61	22 SUSSEX ST	Hackensack	NJ	07601	\N	\N
62	75 OAK STREET	Patchogue	NY	11772	\N	\N
63	1 CLINTON AVE	Albany	NY	12207	\N	\N
64	7242 ROUTE 9	Plattsburgh	NY	12901	\N	\N
65	520 5TH AVE	McKeesport	PA	15132	\N	\N
66	122 W 3RD STREET	Greensburg	PA	15601	\N	\N
67	901 UNIVERSITY DR	State College	PA	16801	\N	\N
68	240 W 3RD ST	Williamsport	PA	17701	\N	\N
69	41 N 4TH ST	Allentown	PA	18102	\N	\N
70	2221 W. MARKET STREET	Pottsville	PA	17901	\N	\N
71	337 BRIGHTSEAT ROAD	Hyattsville	MD	20785	\N	\N
72	101 CHESAPEAKE BLVD	Elkton	MD	21921	\N	\N
73	2875 SABRE ST	Virginia Beach	VA	23452	\N	\N
74	324 COMMERCE ROAD	Clarksville	VA	23927	\N	\N
75	1480 EAST MAIN STREET	Wytheville	VA	24382	\N	\N
76	116 N JEFFERSON STREET	Roanoke	VA	24016	\N	\N
77	50 MCDOWELL STREET	Welch	WV	24801	\N	\N
78	146 EAST FIRST AVE	Williamson	WV	25661	\N	\N
79	1925 E MAIN ST	Albemarle	NC	28001	\N	\N
80	1013 SPRING LANE	Sanford	NC	27330	\N	\N
84	732 W 2ND ST	Tifton	GA	31793	\N	\N
\.


--
-- TOC entry 3349 (class 0 OID 0)
-- Dependencies: 214
-- Name: address_address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pi
--

SELECT pg_catalog.setval('public.address_address_id_seq', 106, true);


-- Completed on 2024-10-23 06:47:01

--
-- PostgreSQL database dump complete
--

