package main

import "encoding/xml"

// Fit XML into structs.
func unmarshalXML(input []byte) []Offer {
	c := Container{}
	xml.Unmarshal(input, &c)

	offers := c.Results[0].Offers

	return offers
}
