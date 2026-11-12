package setmatrixzero

import "math"

type item struct {
	r int
	c int
}

func SetZeroes(matrix [][]int) {

	rows := len(matrix)
	cols := len(matrix[0])
	const marker = math.MaxInt64

	var setToNegative func(r, c int)
	setToNegative = func(r, c int) {

		for i := 0; i < cols; i++ {
			if matrix[r][i] != 0 {
				matrix[r][i] = marker
			}
		}
		for i := 0; i < rows; i++ {

			if matrix[i][c] != 0 {
				matrix[i][c] = marker
			}

		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				setToNegative(i, j)
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 || matrix[i][j] == marker {
				matrix[i][j] = 0
			}
		}
	}

}
