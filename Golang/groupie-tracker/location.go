package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

type Geocoord struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Name string
				Lat  float64
				Lng  float64
			}
		}
	}
}

type Coordinates struct {
	Name string
	Lat  float64
	Lng  float64
}

var wg sync.WaitGroup

func Geocode() {
	coords := make(map[string]Coordinates)

	c := make(chan Coordinates)

	for k := range data.Search.Location {
		// goroutines are used to make parsing the data faster
		// time.sleep is used because google geocoding API has a request limit 
		// of 50 requests per second 
		wg.Add(1)
		time.Sleep(25 * time.Millisecond)
		go LocationJSONtoStruct(k, c)

	}
	go func() {

		for i := 0; i < len(data.Search.Location); i++ {
			wg.Add(1)
			data := <-c
			coords[data.Name] = data
			wg.Done()
		}

	}()
	wg.Wait()
	AppendCoordinates(coords)
}

func LocationJSONtoStruct(k string, c chan Coordinates) {
	defer wg.Done()
	APIkey := "&key=AIzaSyAXYCp5amU1mjbngo1GRdB79HmtqyjDzYQ"
	mapURL := "https://maps.googleapis.com/maps/api/geocode/json?address="
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	loc := re.ReplaceAllString(k, "+")
	response := GetResponse(mapURL + loc + APIkey)
	var coordJSON Geocoord
	err := json.NewDecoder(response.Body).Decode(&coordJSON)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var tmpCoord Coordinates

	tmpCoord.Name = k
	tmpCoord.Lat = coordJSON.Results[0].Geometry.Location.Lat
	tmpCoord.Lng = coordJSON.Results[0].Geometry.Location.Lng
	c <- tmpCoord
}

func AppendCoordinates(coords map[string]Coordinates) {
	for i, artist := range artists {
		for _, location := range artist.Location {
			if len(coords[location].Name) > 0 {
				artists[i].Coordinates = append(artists[i].Coordinates, coords[location])
			}
		}
	}
}
