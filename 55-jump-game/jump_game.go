package jumpgame

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

func CanJump(nums []int) bool {
	dis := 0

	for i := 0; i < len(nums) && i <= dis; i++ {
		if nums[i]+i > dis {
			dis = nums[i] + i
		}

		if dis >= len(nums)-1 {
			return true
		}
	}

	return false
}
