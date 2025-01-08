package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Check for correct arg count
	if len(os.Args) != 2 {
		fmt.Println(`Correct usage: go run . "word"`)
		os.Exit(1)
	}
	// Open banner file 
	file, err := os.Open("standard.txt")
	checkErr(err)
	defer file.Close()

	// Get []string where each element is a line from txt file
	letters := txtToSlice(file)
	
	// Get words and charcount from user input
	words, charcount := inputToSlice()

	// Print out the words
	printLetters(words, letters, charcount)	
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func inputToSlice() ([]string, int) {
	// Split words to slices by \n 
	re := regexp.MustCompile(`\\n`)
	words := re.Split(os.Args[1], -1)
	// Char counter for correct printing of newlines
	charcount := 0
	for _, v := range words {
		charcount += len(v)
	}
	return words, charcount
}

func printLetters(words, letters []string, charcount int)  {
	// Check to avoid printing line to empty input
	if len(os.Args[1]) > 0 {
		// Range through words and print out first line of every letter, then second etc.
		for _, v := range words {
			if len(v) > 0 {
				for i := 1; i < 9; i++ {
					for _, vi := range v {
						fmt.Print(letters[(int(vi)-32)*9+i])
					}
					fmt.Println()
				}
				
			} else {
				// Check to avoid printing extra newline when input consists of only \n
				if charcount > 0{
					fmt.Println()
				} else {
					charcount++
				}
				
			}
		}
	}
}

func txtToSlice(file *os.File) []string {
	// Init scanner for reading input file
	scanner := bufio.NewScanner(file)
	// Store lines to []string
	letters := []string{}
	for scanner.Scan() {
		letters = append(letters, scanner.Text())
	}
	return letters
}
