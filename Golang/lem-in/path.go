package main

import (
	"sort"
)

type QuickPath struct {
	NumMoves float64
	Paths    [][]string
}

var paths [][]string
var bPaths QuickPath

func FindPaths(room string, path []string) bool {
	if room == end {
		return true
	}
	if rooms[room].Visited {
		return false
	}
	tmproom := rooms[room]
	tmppath := append(path, room)

	for _, v := range tmproom.Links {
		tmproom.Visited = true
		rooms[room] = tmproom
		if FindPaths(v, tmppath) {
			tmppath = append([]string(nil), tmppath...)
			if CheckMultipleEnds(tmppath) {
				tmppath = append(tmppath, end)
				paths = append(paths, tmppath)
			}
		}
		tmproom.Visited = false
		rooms[room] = tmproom
	}

	return false
}

func CheckMultipleEnds(path []string) bool {
	pathlen := len(path) - 1
	for i, v := range path {
		if i < pathlen && v == end {
			return false
		}
	}
	return true
}

func SortPaths() {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	if len(paths) == 0 {
		printError("invalid data format, no paths found")
	}
}

func FindBestPathCombination(qPath [][]string) {
	tmppaths := qPath
	for _, path := range paths {
		if ComparePaths(path, tmppaths) {
			tmppaths = append(tmppaths, path)
			if num := CalculateMoves(tmppaths); num < bPaths.NumMoves || bPaths.Paths == nil {
				bPaths.Paths = tmppaths
				bPaths.NumMoves = num

			}
			FindBestPathCombination(tmppaths)
			tmppaths = qPath
		}
	}
}

// Formula for finding how many moves it takes depending on the paths used is
// (number of ants + length of the path of all paths - (number of paths * 2)) / number of paths
func CalculateMoves(p [][]string) float64 {
	sum := antCount - (len(p) * 2)
	for _, v := range p {
		sum += len(v)
	}
	return float64(sum) / float64(len(p))
}

func ComparePaths(p []string, ps [][]string) bool {
	for _, v := range ps {
		for _, vi := range v {
			if (vi != start && vi != end) || (len(v) == 2 && len(p) == 2) {
				for _, vj := range p {
					if vi == vj {
						return false
					}
				}
			}

		}
	}
	return true
}
