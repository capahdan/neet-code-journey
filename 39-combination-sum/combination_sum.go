package combinationsum

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	var dfs func(start, remain int)

	dfs = func(start, remain int) {
		if remain == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			result = append(result, comb)
			return
		}

		if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])

			// i instead of i+1 because we can reuse the same number
			dfs(i, remain-candidates[i])

			path = path[:len(path)-1]
		}

	}

	dfs(0, target)
	return result
}
