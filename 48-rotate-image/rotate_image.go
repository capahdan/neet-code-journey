package rotateimage

func rotate(matrix [][]int) {

	stact := []int{}

	for i := range matrix {
		for j := range matrix {
			stact = append(stact, matrix[i][j])
		}
	}

	idx := 0

	for i := range matrix {
		x := len(matrix) - 1
		for j := range matrix[i] {
			matrix[j][x-i] = stact[idx]
			idx++
		}
	}

}

func rotateImage(matrix [][]int) {
	l, r := 0, len(matrix)-1

	for l < r {
		for i := 0; i < r-l; i++ {
			top, bottom := l, r

			// save top-left
			topLeft := matrix[top][l+i]

			// bottom-left -> top-left
			matrix[top][l+i] = matrix[bottom-i][l]

			// bottom-right -> bottom-left
			matrix[bottom-i][l] = matrix[bottom][r-i]

			// top-right -> bottom-right
			matrix[bottom][r-i] = matrix[top+i][r]

			// saved top-left -> top-right
			matrix[top+i][r] = topLeft
		}

		l++
		r--
	}
}
