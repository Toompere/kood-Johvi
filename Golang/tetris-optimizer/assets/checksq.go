package assets

func CheckSq(tetr tetrimino, square [][]string, x, y, size int) bool {
	count := 0
	for i := 0; i < 4; i++ { 
		// This changes y if current tetrimino doesnt strt from upper left corner
		if tetr[0][i] != "." {
			if y-i < 0 {
				return false
			}
			y -= i
			for j := 0; j < 4; j++ {
				for k := 0; k < 4; k++ {
					if tetr[j][k] != "." {
						if x+j >= size || y+k >= size || square[x+j][y+k] != "." {
							return false
						}
						count++
						if count == 4 {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
