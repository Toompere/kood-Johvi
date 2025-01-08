package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Response struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

type Data struct {
	Artists []Artist
	Search  Datalist
	Details Artist
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Location     []string
	Date         []string
	Relation     map[string][]string
	Coordinates  []Coordinates
}

type Locations struct {
	Index []struct {
		Id        int
		Locations []string
	}
}

type Dates struct {
	Index []struct {
		Id    int
		Dates []string
	}
}

type Relation struct {
	Index []struct {
		DatesLocations map[string][]string
	}
}

// global variables
var links Response
var artists []Artist
var locations Locations
var dates Dates
var relation Relation
var data Data
var filterData Data

func GetData() {

	JSONtoStruct("https://groupietrackers.herokuapp.com/api", &links)
	JSONtoStruct(links.Artists, &artists)
	JSONtoStruct(links.Locations, &locations)
	JSONtoStruct(links.Dates, &dates)
	JSONtoStruct(links.Relation, &relation)

	// After unmarshaling change location and relation to a better format
	// then append everything to one struct that will be used with handler
	CleanData()
	AppendData()
	data.Artists = artists
	data.Details = artists[0]
	SearchData()
	filterData = data
}

func JSONtoStruct(url string, data interface{}) {
	response := GetResponse(url)
	err := json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func GetResponse(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return response
}

func CleanData() {
	for index, v := range locations.Index {
		tempRelation := make(map[string][]string)
		for i := range v.Locations {
			key := v.Locations[i]
			re := regexp.MustCompile("_")
			v.Locations[i] = re.ReplaceAllString(v.Locations[i], " ")
			v.Locations[i] = strings.Title(v.Locations[i])
			re = regexp.MustCompile("-([a-zA-Z]{2,3})$")
			v.Locations[i] = re.ReplaceAllStringFunc(v.Locations[i], func(s string) string {
				return strings.ToUpper(s)
			})
			re = regexp.MustCompile("-")
			v.Locations[i] = re.ReplaceAllString(v.Locations[i], " - ")
			tempRelation[v.Locations[i]] = relation.Index[index].DatesLocations[key]
		}
		relation.Index[index].DatesLocations = tempRelation
	}

}

func AppendData() {
	for i := 0; i < len(artists); i++ {
		artists[i].Location = locations.Index[i].Locations
		artists[i].Date = dates.Index[i].Dates
		artists[i].Relation = relation.Index[i].DatesLocations
	}
}
