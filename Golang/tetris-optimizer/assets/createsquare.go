package assets

func CreateSquare(n int) [][]string {
	square := [][]string{}
	// Fill square with dots
	for i := 0; i < n; i++ {
		square = append(square, []string{})
		for j := 0; j < n; j++ {
			square[i] = append(square[i], ".")
		}
	}
	return square
}
