package missingnumber

func missingNumber(nums []int) int {
	var result int
	for i := 0; i <= len(nums); i++ {
		result ^= i
	}

	for _, num := range nums {
		result ^= num
	}

	return result
}

func missingNumberWithoutXOR(nums []int) int {
	res := len(nums)
	for i := range nums {
		res += (i - nums[i])
	}
	return res
}
