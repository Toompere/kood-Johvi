package assets

import (
	"bufio"
	"fmt"
	"os"
)

type tetrimino [4][4]string

func TxtToTetrimino() ([]tetrimino, bool) {
	tetriminoes := []tetrimino{}
	f, err := os.Open("./" + os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	tmptetrimino := tetrimino{}
	counter := 0

	for scanner.Scan() {
		// Check for correct format
		if (len(scanner.Text()) != 4 && counter < 4) || (len(scanner.Text()) != 0 && counter == 4) {
			return tetriminoes, false
		}

		// Search for # and replace with letter
		letter := string(len(tetriminoes) + 65)
		if counter == 4 {
			tmptetrimino = tetrimino{}
			counter = 0
		} else {
			for i, v := range scanner.Text() {
				if v == '#' {
					tmptetrimino[counter][i] = letter
				} else if v == '.' {
					tmptetrimino[counter][i] = "."
				} else {
					return tetriminoes, false
				}
			}
			if counter == 3 {
				if !validTetrimino(tmptetrimino, letter) {
					return tetriminoes, false
				}
				tmptetrimino = RmEmptyRow(tmptetrimino)
				tmptetrimino = RmEmptyCol(tmptetrimino)
				tetriminoes = append(tetriminoes, tmptetrimino)
			}
			counter++
		}
	}
	return tetriminoes, true
}

func validTetrimino(tetr tetrimino, letter string) bool {
	blckcount := 0 // amount of blocks
	nextcount := 0 // amount of blocks connected to each other minus the last one
	rowcount := 0  //  amount of blocks connected to a block on the next row

	// loops through the tetrimino and counts blocks and connections
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetr[i][j] == letter {
				blckcount++
				if (i > 0 && tetr[i-1][j] == letter) || (i < 3 && tetr[i+1][j] == letter) ||
					(j > 0 && tetr[i][j-1] == letter) || (j < 3 && tetr[i][j+1] == letter) {
					if blckcount == 3 {
						tetr[i][j] = ""
					}
					if i < 3 && tetr[i+1][j] == letter {
						rowcount++
					}
					nextcount++
				}
			}
		}
		if blckcount > 0 && blckcount < 4 && rowcount == 0 {
			return false
		}
		tetr[i] = [4]string{}
	}
	if blckcount == 4 && nextcount == 3 {
		return true
	}
	return false
}

func RmEmptyRow(tetr tetrimino) tetrimino {
	for i := 0; i < 4; i++ {
		if tetr[0][i] != "." {
			return tetr
		}
	}
	// moves first row to last
	tetr[0], tetr[1], tetr[2], tetr[3] = tetr[1], tetr[2], tetr[3], tetr[0]
	return RmEmptyRow(tetr)
}

func RmEmptyCol(tetr tetrimino) tetrimino {
	for i := 0; i < 4; i++ {
		if tetr[i][0] != "." {
			return tetr
		}
	}
	for i := 0; i < 4; i++ {
		// moves first columnt to last
		tetr[i][0], tetr[i][1], tetr[i][2], tetr[i][3] = tetr[i][1], tetr[i][2], tetr[i][3], tetr[i][0]
	}
	return RmEmptyCol(tetr)
}
