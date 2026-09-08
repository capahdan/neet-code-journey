package maximumproductsubarray

func MaxProduct(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	curMax, curMin, maxSoFar := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(curMax*num, num)
		curMin = min(curMin*num, num)
		maxSoFar = max(maxSoFar, curMax)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
