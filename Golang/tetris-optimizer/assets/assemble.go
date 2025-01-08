package assets

func Assemble(tetriminoes []tetrimino, square [][]string, size int) bool {
	// All tetriminoes have been placed in the square, return true
	if len(tetriminoes) == 0 {
		return true
	}

	// Loop through square and search for emtpy spots, then place tetrimino if possible
	// Backtracking until true
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if square[x][y] == "." {
				if CheckSq(tetriminoes[0], square, x, y, size) {
					WriteTetr(tetriminoes[0], square, x, y)
					if Assemble(tetriminoes[1:], square, size) {
						return true
					}
					RemvTetr(tetriminoes[0], square, x, y)
				}
			}
		}
	}
	// If no possible solutions return false to main to increase square size
	return false
}
