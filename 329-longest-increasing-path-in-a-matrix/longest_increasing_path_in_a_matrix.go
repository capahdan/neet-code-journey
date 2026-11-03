package longestincreasingpathinamatrix

type item struct {
	i int
	j int
}

func LongestIncreasingPath(matrix [][]int) int {

	rows := len(matrix)
	cols := len(matrix[0])
	cache := make(map[item]int)

	var dfs func(prevValue, i, j int) int
	dfs = func(prevValue, i, j int) int {
		if i < 0 || i >= rows || j < 0 || j >= cols || prevValue >= matrix[i][j] {
			return 0
		}

		key := item{
			i: i,
			j: j,
		}

		if val, ok := cache[key]; ok {
			return val
		}

		currValue := matrix[i][j]

		a := dfs(currValue, i+1, j)
		b := dfs(currValue, i-1, j)
		c := dfs(currValue, i, j+1)
		d := dfs(currValue, i, j-1)

		cache[key] = max(a, b, c, d) + 1
		return cache[key]
	}

	result := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result = max(result, dfs(-1, i, j))
		}
	}

	return result
}
