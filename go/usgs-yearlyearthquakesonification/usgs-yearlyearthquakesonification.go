// https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2014-01-01&endtime=2014-01-02

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	// "strings"
	"time"
)

type featureCollection struct {
	Features []feature `json:"features"`
}

type feature struct {
	Properties Earthquake `json:"properties"`
}

type Earthquake struct {
	Title     string  `json:"title"`
	Magnitude float64 `json:"mag"`
}

func main() {
	file, err := os.Create("magnitudes.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done

	// Get the current time
	current := time.Now()
	currentFormat := current.Format("2006-01-02")

	yesterdayTime := time.Now().Add(-24 * time.Year)
	yesterFormat := yesterdayTime.Format(("2006-01-02"))
	findHawaiianVolcanos := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=" + yesterFormat + "&endtime=" + currentFormat

	resp, err := http.Get(findHawaiianVolcanos)
	if err != nil {
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var record featureCollection
	json.Unmarshal(body, &record)

	quakes := make([]Earthquake, 0)
	for _, f := range record.Features {
		quakes = append(quakes, f.Properties)
	}

	responseLength := len(quakes)

	for q := 0; q < responseLength; q++ {
		len3, err := file.WriteString(fmt.Sprintln(quakes[q].Magnitude))

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		_ = len3
	}
	CreateCsoundFile()
}

func CreateCsoundFile() {
	fmt.Println("creating CSD")
	fileHandle, _ := os.Create("output.csd")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, `<CsoundSynthesizer>
<CsOptions>
-odac
</CsOptions>
<CsInstruments>

sr = 44100
ksmps = 32
nchnls = 2
0dbfs = 1

instr 1

klin linseg   0, 3, 10
; aosc foscil .5, 110, 1, (klin*3.17), 1*klin, 1
aosc oscil .5, 110, 1

outs aosc, aosc

endin

</CsInstruments>
<CsScore>
f 1 0 4096 -23 "magnitudes.txt"
i 1 0 10
e
</CsScore>
</CsoundSynthesizer>

`)
	writer.Flush()

	cmd := exec.Command("csound", "-W", "-o", "SonifyData.wav", "output.csd")
	cmd.Dir = "/Users/io/Documents/_airReam/five-and-dime/go/usgs-yearlyearthquakesonification"
	cmd.Run()
}
