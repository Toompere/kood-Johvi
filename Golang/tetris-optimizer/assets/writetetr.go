package assets

func WriteTetr(tetr tetrimino, square [][]string, x, y int) {
	count := 0
	for i := 0; i < 4; i++ {
		if tetr[0][i] != "." {
			y -= i
			for j := 0; j < 4; j++ {
				for k := 0; k < 4; k++ {
					if tetr[j][k] != "." {
						square[x+j][y+k] = tetr[j][k]
						count++
						if count == 4 {
							return
						}
					}
				}
			}
		}
	}
}
