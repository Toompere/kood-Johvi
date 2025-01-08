package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Datalist struct {
	Name         map[string][]int
	Members      map[string][]int
	Location     map[string][]int
	FirstAlbum   map[string][]int
	CreationDate map[string][]int
	CreationList []int
}

var search Datalist

func SearchData() {
	// Creates a map where all artist ids are appended which match the same keywords
	nameMap := make(map[string][]int)
	membersMap := make(map[string][]int)
	locationMap := make(map[string][]int)
	albumMap := make(map[string][]int)
	creationMap := make(map[string][]int)
	var creationList []int
	for i := 0; i < len(artists); i++ {
		nameMap[artists[i].Name] = append(nameMap[artists[i].Name], artists[i].Id)
		albumMap[artists[i].FirstAlbum] = append(albumMap[artists[i].FirstAlbum], artists[i].Id)
		creationMap[fmt.Sprint(artists[i].CreationDate)] = append(creationMap[fmt.Sprint(artists[i].CreationDate)], artists[i].Id)
		creationList = append(creationList, artists[i].CreationDate)
		for _, v := range artists[i].Members {
			membersMap[v] = append(membersMap[v], artists[i].Id)
		}
	}
	for i := 0; i < len(locations.Index); i++ {
		for _, val := range locations.Index[i].Locations {
			locationMap[val] = append(locationMap[val], locations.Index[i].Id)
		}

	}
	sort.Ints(creationList)

	search.Name = nameMap
	search.Members = membersMap
	search.Location = locationMap
	search.FirstAlbum = albumMap
	search.CreationDate = creationMap
	search.CreationList = creationList
	data.Search = search
}

func SearchResult(srch string) Data {
	var resultData Data
	resultData.Search = search
	var match []int
	re := regexp.MustCompile(`(.+) - (.+)$`)
	searchWord := re.ReplaceAllString(srch, "$1")
	category := re.ReplaceAllString(srch, "$2")
	//If a entry from datalist is selected in the search bar, it only finds match from one category
	// otherwise finds matches from all categories
	switch category {
	case "artist/band":
		match = FindMatch(search.Name, searchWord)
	case "member":
		match = FindMatch(search.Members, searchWord)
	case "location":
		match = FindMatch(search.Location, searchWord)
	case "first album":
		match = FindMatch(search.FirstAlbum, searchWord)
	case "creation date":
		match = FindMatch(search.CreationDate, searchWord)
	default:
		match = FindMatch(search.Name, srch)
		match = append(match, FindMatch(search.Members, srch)...)
		match = append(match, FindMatch(search.Location, srch)...)
		match = append(match, FindMatch(search.FirstAlbum, srch)...)
		match = append(match, FindMatch(search.CreationDate, srch)...)
	}

	resultData.Artists = AppendResult(RemoveDuplicate(match))

	return resultData
}

func AppendResult(id []int) []Artist {
	var resultArtists []Artist
	for _, v := range id {
		resultArtists = append(resultArtists, artists[v-1])
	}
	return resultArtists
}

func FindMatch(data map[string][]int, srch string) []int {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	srch = re.ReplaceAllString(srch, " ")
	var matchID []int
	if srch == " " {
		return matchID
	}
	for k, v := range data {
		key := re.ReplaceAllString(k, " ")
		if strings.Contains(strings.ToLower(key), strings.ToLower(srch)) {
			matchID = append(matchID, v...)
		}
	}
	return matchID
}

// This function is needed so when searching for something
// it would not sugest same location for example multiple times
func RemoveDuplicate(matchid []int) []int {
	allIds := make(map[int]bool)
	ids := []int{}
	for _, id := range matchid {
		if _, value := allIds[id]; !value {
			allIds[id] = true
			ids = append(ids, id)
		}
	}
	sort.Ints(ids)
	return ids
}
