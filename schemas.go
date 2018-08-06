package main

type Container struct {
	Results []Results `xml:"Results"`
}

type Results struct {
	Offers []Offer `xml:"Offer"`
}

type Offer struct {
	Sellprice      string `xml:"Sellprice,attr"`
	Flightnetprice string `xml:"Flightnetprice,attr" json:"-"`
	Hotelnetprice  string `xml:"Hotelnetprice,attr" json:"-"`
	Brochurecode   string `xml:"Brochurecode,attr" json:"-"`
	Ourhtlid       string `xml:"Ourhtlid,attr" json:"-"`
	Starrating     string `xml:"Starrating,attr"`
	Boardbasis     string `xml:"Boardbasis,attr" json:"-"`
	Roomtype       string `xml:"Roomtype,attr" json:"-"`
	Resortname     string `xml:"Resortname,attr" json:"-"`
	Hotelname      string `xml:"Hotelname,attr"`
	Duration       string `xml:"Duration,attr" json:"-"`
	Inboundfltnum  string `xml:"Inboundfltnum,attr"`
	Inboundarr     string `xml:"Inboundarr,attr"`
	Inbounddep     string `xml:"Inbounddep,attr"`
	Outboundfltnum string `xml:"Outboundfltnum,attr"`
	Outboundarr    string `xml:"Outboundarr,attr" json:"-"`
	Outbounddep    string `xml:"Outbounddep,attr" json:"-"`
	Arraptname     string `xml:"Arraptname,attr" json:"-"`
	Arraptcode     string `xml:"Arraptcode,attr" json:"-"`
	Depaptname     string `xml:"Depaptname,attr" json:"-"`
	Depaptcode     string `xml:"Depaptcode,attr" json:"-"`
	Flightsuppler  string `xml:"Flightsuppler,attr" json:"-"`
	Hotelsupplier  string `xml:"Hotelsupplier,attr" json:"-"`
	Type           string `xml:"Type,attr" json:"-"`
}
