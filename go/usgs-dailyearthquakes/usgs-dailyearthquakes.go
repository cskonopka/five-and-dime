// Use the USGS URL below, open a browser and go to the following link. Note the *starttime* and *endtime* determine the date range of the results.
// https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2014-01-01&endtime=2014-01-02

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Use [JSON-to-Struct](https://mholt.github.io/json-to-go/) to create a struct from the JSON response. The modified version below focuses on the earthquake’s *Place* and *Magnitude* values.
type usgsJSON struct {
	Features []feature `json:"features"`
}

type feature struct {
	Properties Earthquake `json:"properties"`
}

type Earthquake struct {
	Place     string  `json:"place"`
	Magnitude float64 `json:"mag"`
}

func main() {
	// Get the current date.
	current := time.Now()
	currentFormat := current.Format("2006-01-02")

	// Get yesterdays date.
	yesterdayTime := time.Now().Add(-24 * time.Hour)
	yesterFormat := yesterdayTime.Format(("2006-01-02"))

	// Construct the USGS URL for the GET request using the *currentFormat* and *yesterFormat*.
	findHawaiianVolcanos := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=" + yesterFormat + "&endtime=" + currentFormat

	// Perform a GET request.
	resp, err := http.Get(findHawaiianVolcanos)
	if err != nil {
	}

	// Read the data until EOF and return the data to the body variable.
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Unmarshal the body into usgsJSON struct.
	var record usgsJSON
	json.Unmarshal(body, &record)

	// Make a slice of "Earthquake" of length 0.
	quakes := make([]Earthquake, 0)

	// Iterate over each "features" object and append the "properties" object to the quakes slice.
	for _, f := range record.Features {
		quakes = append(quakes, f.Properties)
	}

	// Iterate over the length of quakes and print out the magnitudes and location.
	for q := 0; q < len(quakes); q++ {
		fmt.Print(quakes[q].Place + " ")
		fmt.Println(quakes[q].Magnitude)
	}
}
