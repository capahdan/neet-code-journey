package houserobber2

// Brute force solution and it works
// func Rob(nums []int) int {

// 	if len(nums) <= 1 {
// 		return nums[0]
// 	}

// 	maxSoFar := 0
// 	l := len(nums)
// 	dp := make([]int, l)
// 	dp2 := make([]int, l)

// 	for i := l - 2; i >= 0; i-- {
// 		jump2 := 0
// 		jump3 := 0

// 		if i+2 < l {
// 			jump2 = dp[i+2]
// 		}
// 		if i+3 < l {
// 			jump3 = dp[i+3]
// 		}
// 		dp[i] = nums[i] + max(jump2, jump3)
// 		maxSoFar = max(maxSoFar, dp[i])
// 	}

// 	for i := l - 1; i > 0; i-- {
// 		jump2 := 0
// 		jump3 := 0

// 		if i+2 < l {
// 			jump2 = dp2[i+2]
// 		}
// 		if i+3 < l {
// 			jump3 = dp2[i+3]
// 		}
// 		dp2[i] = nums[i] + max(jump2, jump3)
// 		maxSoFar = max(maxSoFar, dp2[i])
// 	}

// 	res1 := max(dp[0], dp[1])
// 	res2 := max(dp2[0], dp2[1])
// 	resFinal := max(res1, res2)

// 	return max(maxSoFar, resFinal)
// }

// the intuition of this function is
// 1. we cannot rob house that adjacent to each other it means we neet to skip 1 house
// 2. because of that we use helper function to do that
// 3. since the first house and the second house is adjection we can not use the two together
//.   because of that we calculate the max value by igoring the last value in the first
//    and ignoring the first one in the second helper function

func Rob(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	return max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	prev2, prev1 := 0, 0

	for _, v := range nums {
		curr := max(prev1, prev2+v)
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
