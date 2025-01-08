package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Room struct {
	Name    string
	CoordX  int
	CoordY  int
	Links   []string
	Ant     int
	Visited bool
}

var lines []string
var linenum int
var antCount int
var rooms map[string]Room
var start, end string

func ParseInput() {
	if len(os.Args) < 2 {
		printError(`Usage: "go run . filename.txt"`)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rooms = make(map[string]Room)

	for scanner.Scan() {
		line := CheckComment(scanner)
		switch {
		case IsAntCount(line):
		case IsRoom(line, ""):
		case line == "##start":
			scanner.Scan()
			if !IsRoom(CheckComment(scanner), "start") {
				printError("invalid data format, line " + fmt.Sprint(linenum))
			}
		case line == "##end":
			scanner.Scan()
			if !IsRoom(CheckComment(scanner), "end") {
				printError("invalid data format, line " + fmt.Sprint(linenum))
			}
		case IsLink(line):
		default:
			printError("invalid data format, line " + fmt.Sprint(linenum))
		}
	}
	if len(start) == 0 {
		printError("invalid data format, no start room found ")
	}
	if len(end) == 0 {
		printError("invalid data format, no end room found ")
	}
}

func CheckComment(scanner *bufio.Scanner) string {
	line := scanner.Text()
	linenum++
	lines = append(lines, line)
	re := regexp.MustCompile(`^#[^#]`)
	if match := re.FindAllString(line, 1); match != nil || line == "#"{
		scanner.Scan()
		return CheckComment(scanner)
	}
	return line
}

func IsAntCount(line string) bool {
	if n, err := strconv.Atoi(line); err == nil {
		if antCount != 0 || n < 1 {
			printError("invalid data format, incorrect number of ants")
		} else {
			antCount = n
			return true
		}
	}
	return false
}

func IsRoom(line, arg string) bool {
	re := regexp.MustCompile(`^([^\s]+)\s(\d+)\s(\d+)$`)
	var tmproom Room
	if match := re.FindStringSubmatch(line); match != nil {
		if _, exists := rooms[match[1]]; exists || match[1][0] == 'l' || match[1][0] == 'L' {
			printError("invalid data format, multiple rooms with identical name")
		}
		tmproom.Name = match[1]
		tmproom.CoordX, _ = strconv.Atoi(match[2])
		tmproom.CoordY, _ = strconv.Atoi(match[3])
		rooms[match[1]] = tmproom
		switch arg {
		case "start":
			if len(start) > 0 {
				printError("invalid data format, multiple start rooms")
			}
			start = tmproom.Name
		case "end":
			if len(end) > 0 {
				printError("invalid data format, multiple end rooms")
			}
			end = tmproom.Name
		}
		return true
	}
	return false
}

func IsLink(line string) bool {
	re := regexp.MustCompile(`^([^\s]+)-([^\s]+)$`)
	var tmplink Room
	if match := re.FindStringSubmatch(line); match != nil {
		if _, exists := rooms[match[1]]; !exists {
			printError("invalid data format, path leads to invalid room, line " + fmt.Sprint(linenum))
		}
		if _, exists := rooms[match[2]]; !exists {
			printError("invalid data format, path leads to invalid room, line " + fmt.Sprint(linenum))
		}
		tmplink = rooms[match[1]]
		tmplink.Links = append(tmplink.Links, match[2])
		rooms[match[1]] = tmplink
		tmplink = rooms[match[2]]
		tmplink.Links = append(tmplink.Links, match[1])
		rooms[match[2]] = tmplink
		return true
	}
	return false
}

func PrintFileContent() {
	for _, v := range lines {
		fmt.Println(v)
	}
}
