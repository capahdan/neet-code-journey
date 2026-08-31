package mincostclimbingstairs

// this function works because we transverse the minClimbing stairs from the end
// first thing that we do is to set an array with the length of the cost array + 1
// that because it will be the sign for the end of the stairs also because we will need an extra position
// after that we will try to add the value of current index and add with minimum value of
// the first and the second positon of the dp array
// after that we will sum up and then we compare index 0 and index 1 which one that have minimum value

func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	dp := make([]int, n+1)
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		next1 := dp[i+1]
		next2 := 0
		if i+2 <= n {
			next2 = dp[i+2]
		}
		dp[i] = cost[i] + min(next1, next2)
	}
	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
