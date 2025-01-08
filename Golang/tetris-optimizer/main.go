package main

import (
	"fmt"
	"math"
	"os"
	"tetris-optimizer/assets"
	"time"
)

var tim = time.Now()

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Correct usage go run . input.txt")
		os.Exit(1)
	}
	// writes tetriminoes from file to slice, return false if bad format or not valid tetrimino
	tetriminoes, valid := assets.TxtToTetrimino()
	if !valid {
		fmt.Println("ERROR")
	} else {

		// start with the smallest square possbile depending on the count of tetriminoes
		size := int(math.Ceil(math.Sqrt(float64(len(tetriminoes) * 4))))
		square := assets.CreateSquare(size)

		// loop until backtracking function returns true
		for !assets.Assemble(tetriminoes, square, size) {
			size++
			square = assets.CreateSquare(size)
		}

		assets.PrintSquare(square)

	}
}
