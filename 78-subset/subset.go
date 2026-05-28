package subset

func Subset(nums []int) [][]int {

	var result [][]int
	var backTrack func(start int, curr []int)
	backTrack = func(start int, curr []int) {

		temp := make([]int, len(curr))
		copy(temp, curr)
		result = append(result, temp)

		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i]) // choose
			backTrack(i+1, curr)         // explore
			curr = curr[:len(curr)-1]    // unchoose
		}
	}

	backTrack(0, []int{})
	return result
}
