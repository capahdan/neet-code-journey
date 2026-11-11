package spiralmatrix

func SpiralOrder(matrix [][]int) []int {

	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1
	result := []int{}

	for left <= right && top <= bottom {
		// top  -> left to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// this is a condition where still only one column
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// this is a condition where still only one row
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}

	}

	return result
}
