package main

const externalUrl = "http://87.102.127.86:8100/search/searchoffers.dll?page=SEARCH&platform=WEB&depart=LGW%7CSTN%7CLHR%7CLCY%7CSEN%7CLTN&countryid=1&regionid=4&areaid=9&resortid=0&depdate=15%2F08%2F2018&flex=0&adults=2&children=0&duration=7"

//const mockData = `
//<Container xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="HolidaysSearchRequest1.6.xsd">
//	<Faults/>
//	<Results count="206" searchtime="00:00.203">
//		<Offer Type="DP" Hotelsupplier="Hotel+Beds" Flightsuppler="Jet2" Depaptcode="STN" Depaptname="London%2C+Stansted+Airport" Arraptcode="TFS" Arraptname="Tenerife%2C+Sur+Int.(Reina+Sofia" Outbounddep="15/08/2018 12:35" Outboundarr="15/08/2018 19:05" Outboundfltnum="LS1663" Inbounddep="22/08/2018 18:05" Inboundarr="23/08/2018 00:20" Inboundfltnum="LS1664" Duration="7" Hotelname="Tropical" Resortname="Puerto+De+La+Cruz" Roomtype="Studio+Capacity+2" Boardbasis="SC" Starrating="3" Ourhtlid="20288" Brochurecode="HBED-4085" Hotelnetprice="62.00" Flightnetprice="362.00" Sellprice="424.35"/>
//		<Offer Type="DP" Hotelsupplier="Hotel+Beds" Flightsuppler="Jet2" Depaptcode="STN" Depaptname="London%2C+Stansted+Airport" Arraptcode="TFS" Arraptname="Tenerife%2C+Sur+Int.(Reina+Sofia" Outbounddep="15/08/2018 13:35" Outboundarr="15/08/2018 19:05" Outboundfltnum="LS1663" Inbounddep="22/08/2018 19:05" Inboundarr="23/08/2018 00:20" Inboundfltnum="LS1664" Duration="7" Hotelname="Rf+Bambi+-+Adults+Only" Resortname="Puerto+De+La+Cruz" Roomtype="Studio+Capacity+3" Boardbasis="RO" Starrating="2" Ourhtlid="10880" Brochurecode="HBED-1966" Hotelnetprice="124.00" Flightnetprice="362.00" Sellprice="486.44"/>
//		<Offer Type="DP" Hotelsupplier="Hotel+Beds" Flightsuppler="Jet2" Depaptcode="STN" Depaptname="London%2C+Stansted+Airport" Arraptcode="TFS" Arraptname="Tenerife%2C+Sur+Int.(Reina+Sofia" Outbounddep="15/08/2018 14:35" Outboundarr="15/08/2018 19:05" Outboundfltnum="LS1663" Inbounddep="22/08/2018 20:05" Inboundarr="23/08/2018 00:20" Inboundfltnum="LS1664" Duration="7" Hotelname="Rf+Bambi+-+Adults+Only" Resortname="Puerto+De+La+Cruz" Roomtype="Studio+Capacity+3" Boardbasis="RO" Starrating="2" Ourhtlid="10880" Brochurecode="HBED-1966" Hotelnetprice="124.00" Flightnetprice="362.00" Sellprice=""/>
//	</Results>
//</Container>
//`

var filters = map[string]string{
	"earliest_departure_time": "Outbounddep",
	"earliest_return_time":    "Inbounddep",
	"max_price":               "Sellprice",
	"min_price":               "Sellprice",
	"star_rating":             "Starrating",
}

var userOutput = map[string]bool{
	"Sellprice":      true,
	"Starrating":     true,
	"Hotelname":      true,
	"Inboundfltnum":  true,
	"Outboundfltnum": true,
	"Inboundarr":     true,
	"Inbounddep":     true,
}

const apiDateFormat = "02/01/2006 15:04"
const userDateFormat = "15:04"
