package main

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

// Validate single offer:
// - check if it has all necessary fields,
// - check if the user's command is valid
// - check if the offer matches the user's criteria.
func validateOffer(offer Offer, query url.Values) (bool, error) {
	for key := range userOutput {
		r := reflect.ValueOf(&offer)
		if v := reflect.Indirect(r).FieldByName(key); v.String() == "" {
			return false, nil
		}
	}

	for _, key := range filters {
		r := reflect.ValueOf(&offer)
		if v := reflect.Indirect(r).FieldByName(key); v.String() == "" {
			return false, nil
		}
	}

	for key := range query {
		if key == "debug" {
			continue
		}
		if _, ok := filters[key]; !ok {
			return false, errors.New(fmt.Sprintf("Unknown parameter: %s", key))
		}

		desired := query.Get(key)

		r := reflect.ValueOf(&offer)
		current := reflect.Indirect(r).FieldByName(filters[key])

		switch key {
		case "default":
			continue
		case "star_rating":
			currentValue, _ := strconv.Atoi(current.String())
			desiredValue, err := strconv.Atoi(desired)
			if err != nil {
				return false, errors.New(fmt.Sprintf("%s is not valid value for %s", desired, key))
			}
			if currentValue < desiredValue {
				return false, nil
			}

		case "min_price":
			currentValue, _ := strconv.ParseFloat(current.String(), 64)
			desiredValue, err := strconv.ParseFloat(desired, 64)
			if err != nil {
				return false, errors.New(fmt.Sprintf("%s is not valid value for %s", desired, key))
			}
			if currentValue < desiredValue {
				return false, nil
			}
		case "max_price":
			currentValue, _ := strconv.ParseFloat(current.String(), 64)
			desiredValue, err := strconv.ParseFloat(desired, 64)
			if err != nil {
				return false, errors.New(fmt.Sprintf("%s is not valid value for %s", desired, key))
			}
			if currentValue > desiredValue {
				return false, nil
			}
		case "earliest_departure_time", "earliest_return_time":
			currentValue, _ := time.Parse(apiDateFormat, current.String())
			desiredValue, err := time.Parse(userDateFormat, desired)
			if err != nil {
				return false, errors.New(fmt.Sprintf("Invalid time format for `%s`.", key))
			}
			if currentValue.Hour() < desiredValue.Hour() ||
				(currentValue.Hour() == desiredValue.Hour() && currentValue.Minute() < desiredValue.Minute()) {
				return false, nil
			}
		}
	}

	return true, nil
}
