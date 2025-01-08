package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func FilterData(crMin, crMax, faMin, faMax, location string, members []string) {
	cMin, err := strconv.Atoi(crMin)
	if err != nil {
		panic(err)
	}

	cMax, err := strconv.Atoi(crMax)
	if err != nil {
		panic(err)
	}

	fMin, err := strconv.Atoi(faMin)
	if err != nil {
		panic(err)
	}

	fMax, err := strconv.Atoi(faMax)
	if err != nil {
		panic(err)
	}

	var cdMatch, faMatch, memberMatch, locationMatch []int

	for _, v := range data.Search.CreationList {
		if v >= cMin && v <= cMax {
			cdMatch = append(cdMatch, data.Search.CreationDate[fmt.Sprint(v)]...)
		}
	}

	re := regexp.MustCompile(`(\d{4})`)
	for _, v := range data.Artists {
		faDate, err := strconv.Atoi(re.FindString(v.FirstAlbum))
		if err != nil {
			panic(err)
		}
		if faDate >= fMin && faDate <= fMax {
			faMatch = append(faMatch, v.Id)
		}
	}

	if len(members) == 0 {
		memberMatch = cdMatch
	} else {
		for _, v := range members {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			for _, artist := range data.Artists {
				if len(artist.Members) == num {
					memberMatch = append(memberMatch, artist.Id)
				}
			}
		}
	}

	if len(location) == 0 {
		locationMatch = cdMatch
	} else {
		locationMatch = FindMatch(search.Location, location)
	}

	filterData.Artists = AppendResult(FindFilterMatch(cdMatch, faMatch, memberMatch, locationMatch))
}

func FindFilterMatch(cd, fa, member, location []int) []int {
	cd = RemoveDuplicate(cd)
	fa = RemoveDuplicate(fa)
	member = RemoveDuplicate(member)
	location = RemoveDuplicate(location)
	match := []int{}

	// Only appends ids that matches all of the filter values
	for _, cdVal := range cd {
		for _, faVal := range fa {
			if cdVal == faVal {
				for _, memb := range member {
					if cdVal == memb {
						for _, loc := range location {
							if cdVal == loc {
								match = append(match, cdVal)
							} else if cdVal < loc {
								break
							}
						}
					} else if cdVal < memb {
						break
					}
				}
			} else if cdVal < faVal {
				break
			}
		}
	}
	return match
}
