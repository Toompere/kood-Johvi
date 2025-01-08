package assets

import "fmt"

func PrintSquare(square [][]string) {
	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square); j++ {
			fmt.Printf(square[i][j])
		}
		fmt.Println()
	}
}
