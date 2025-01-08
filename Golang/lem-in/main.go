package main

import (
	"fmt"
	"os"
)

func main() {
	ParseInput()
	FindPaths(start, []string{})
	SortPaths()
	PrintFileContent()
	FindBestPathCombination([][]string{})
	paths = bPaths.Paths
	MoveAnts()
}

func printError(s string) {
	fmt.Printf("ERROR: %v\n", s)
	os.Exit(0)
}
