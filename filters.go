package main

import (
	"math"
	"net/url"
	"strconv"
)

// Generate summary with:
// - highest price
// - lowest price
// - average price.
func generateSummary(offers []Offer) map[string]interface{} {
	result := map[string]interface{}{
		"summary": map[string]float64{
			"most_expensive_price": math.Inf(-1),
			"cheapest_price":       math.Inf(1),
			"average_price":        0,
		},
		"offers": offers,
	}

	resultSummary, _ := result["summary"].(map[string]float64)

	for _, offer := range offers {
		price, _ := strconv.ParseFloat(offer.Sellprice, 64)
		resultSummary["most_expensive_price"] = math.Max(resultSummary["most_expensive_price"], price)
		resultSummary["cheapest_price"] = math.Min(resultSummary["cheapest_price"], price)
		resultSummary["average_price"] += price
	}

	if resultSummary["average_price"] != 0 {
		resultSummary["average_price"] /= float64(len(offers))
		resultSummary["average_price"] = math.Round(resultSummary["average_price"]*100) / 100
	} else {
		resultSummary["most_expensive_price"] = 0.0
		resultSummary["cheapest_price"] = 0.0
	}

	return result
}

// Main filtering function.
func filterData(query url.Values) (map[string]interface{}, error) {
	input := retrieveData()
	offers := unmarshalXML(input)
	//offers := unmarshalXML([]byte(mockData))

	i := 0
	for _, offer := range offers {
		ok, err := validateOffer(offer, query)
		if err != nil {
			return make(map[string]interface{}), err
		}
		if ok {
			offers[i] = offer
			i++
		}
	}
	offers = offers[:i]

	return generateSummary(offers), nil
}
