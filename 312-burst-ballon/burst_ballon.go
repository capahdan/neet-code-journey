package burstballon

type item struct {
	l int
	r int
}

func maxCoins(nums []int) int {

	arr := make([]int, len(nums)+2)
	for i := range nums {
		arr[i+1] = nums[i]
	}
	arr[0] = 1
	arr[len(arr)-1] = 1

	dp := make(map[item]int)

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}

		key := item{
			l: l,
			r: r,
		}

		if v, ok := dp[key]; ok {
			return v
		}

		dp[key] = 0
		for i := l; i <= r; i++ {
			coins := arr[l-1] * arr[i] * arr[r+1]
			coins += dfs(l, i-1) + dfs(i+1, r)
			dp[key] = max(dp[key], coins)
		}
		return dp[key]
	}

	return dfs(1, len(arr)-2)
}
