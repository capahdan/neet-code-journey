package subsetsii

import "sort"

func SubsetsII(nums []int) [][]int {
	var result [][]int
	var backTrack func(start int, curr []int)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	backTrack = func(start int, curr []int) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		result = append(result, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			backTrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backTrack(0, []int{})
	return result
}
