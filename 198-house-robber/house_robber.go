package houserobber

// the intuition here is we need to sum up all money that we collect from each house
// but to make sure that we don't rob two adjacent houses

// we will use dynamic programming approach
// first we create an array that filled with calculation of current index  + max(dp[2], dp[3])
// after that we compare the value of index[0] and index[i] which one that have bigger value
// finally we return the biggest value

func Rob(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	l := len(nums)
	dp := make([]int, l)

	for i := l - 1; i >= 0; i-- {
		next2 := 0
		next3 := 0
		if i+2 < l {
			next2 = dp[i+2]
		}
		if i+3 < l {
			next3 = dp[i+3]
		}

		dp[i] = nums[i] + max(next2, next3)

	}

	return max(dp[0], dp[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
