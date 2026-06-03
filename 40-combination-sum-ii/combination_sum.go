package combinationsumii

import "sort"

func CombinationSumII(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var dfs func(start int, remain int)

	dfs = func(start int, remain int) {
		if remain == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			dfs(i+1, remain-candidates[i])
			path = path[:len(path)-1]
		}

	}

	dfs(0, target)
	return result
}
