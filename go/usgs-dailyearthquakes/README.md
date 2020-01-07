# usgs-dailyearthquakes

<h3 align="center">
  <img height="255" width="253" src="https://i.ibb.co/19W631Z/usgs-dailyearthquakes-250x250.png"/>
</h3>

### Get daily earthquake data from the USGS using Go. Use the USGS earthquake API endpoint to get a list of daily earthquake locations and magnitudes.*

- [Dev.to](https://dev.to/cskonopka/get-daily-earthquake-data-from-the-usgs-using-go-45gl)

## Walkthrough

Recently I started an earthquake sonification project and the first step is acquiring daily earthquake magnitudes from the USGS.

Use the USGS URL below, open a browser and go to the following link. Note the *starttime* and *endtime* determine the date range of the results.
```
https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2014-01-01&endtime=2014-01-02
```

A large JSON response is returned, but for the purposes of the post a shorter version is provided. The goal is to access the *Place* and *Magnitude* values from the *properties* object inside the *features* object.
``` javascript
{
	"type": "FeatureCollection",
	"metadata": {
		"generated": 1578386362000,
		"url": "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2014-01-01&endtime=2014-01-02",
		"title": "USGS Earthquakes",
		"status": 200,
		"api": "1.8.1",
		"count": 324
	},
	"features": [{
		"type": "Feature",
		"properties": {
			"mag": 1.29,
			"place": "10km SSW of Idyllwild, CA",
			"time": 1388620296020,
			"updated": 1457728844428,
			"tz": -480,
			"url": "https://earthquake.usgs.gov/earthquakes/eventpage/ci11408890",
			"detail": "https://earthquake.usgs.gov/fdsnws/event/1/query?eventid=ci11408890&format=geojson",
			"felt": null,
			"cdi": null,
			"mmi": null,
			"alert": null,
			"status": "reviewed",
			"tsunami": 0,
			"sig": 26,
			"net": "ci",
			"code": "11408890",
			"ids": ",ci11408890,",
			"sources": ",ci,",
			"types": ",cap,focal-mechanism,general-link,geoserve,nearby-cities,origin,phase-data,scitech-link,",
			"nst": 39,
			"dmin": 0.067290000000000003,
			"rms": 0.089999999999999997,
			"gap": 51,
			"magType": "ml",
			"type": "earthquake",
			"title": "M 1.3 - 10km SSW of Idyllwild, CA"
		},
		"geometry": {
			"type": "Point",
			"coordinates": [-116.7776667, 33.663333299999998, 11.007999999999999]
		},
		"id": "ci11408890"
	}]
}
```

Use [JSON-to-Struct](https://mholt.github.io/json-to-go/) to create a struct from the JSON response. A modified version is added below that focuses on the earthquake's *Place* and *Magnitude*.
``` go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

func main(){
     // insert code here
}
```

Get the current date.
```go
current := time.Now()
currentFormat := current.Format("2006-01-02")
```

Get yesterdays date.
```go
yesterdayTime := time.Now().Add(-24 * time.Hour)
yesterFormat := yesterdayTime.Format(("2006-01-02"))
```

Construct the USGS URL for the GET request using the *currentFormat* and *yesterFormat*. 
```go
findHawaiianVolcanos := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=" + yesterFormat + "&endtime=" + currentFormat
```

Perform a GET request. 
```go
resp, err := http.Get(findHawaiianVolcanos)
if err != nil {}
```

Read the data until EOF and return the data to the body variable.
```go
body, err := ioutil.ReadAll(resp.Body)
defer resp.Body.Close()
```

Unmarshal the body into usgsJSON struct.
```go
var record usgsJSON
json.Unmarshal(body, &record)
```

Make a slice of "Earthquake" of length 0.
```go
quakes := make([]Earthquake, 0)
```

Iterate over each "features" object and append the "properties" object to the quakes slice. 
```go 
for _, f := range record.Features {
	quakes = append(quakes, f.Properties)
}
```

Iterate over the length of quakes and print out the magnitudes and location.
```go
for q := 0; q < len(quakes); q++ {
	fmt.Print(quakes[q].Place + " ")
	fmt.Println(quakes[q].Magnitude)
}
```

Build the program.
```go
go build usgs-dailyearthquakes.go
```

List the earthquakes.
![gif](https://i.ibb.co/ftZnDHk/usgs-dailyearthquakes.gif) 
