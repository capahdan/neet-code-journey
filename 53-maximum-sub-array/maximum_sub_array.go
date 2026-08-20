package maximumsubarray

func MaximumSubArray(nums []int) int {
	bestSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		bestSum = max(bestSum, currSum)
	}

	return bestSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
