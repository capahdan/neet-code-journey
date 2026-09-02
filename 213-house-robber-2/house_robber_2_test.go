package houserobber2

import "testing"

func TestHouseRobber(t *testing.T) {
	testTable := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 3, 2},
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			nums:     []int{1, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{1, 2, 1, 1},
			expected: 3,
		},
	}

	for _, tt := range testTable {
		result := Rob(tt.nums)
		if result != tt.expected {
			t.Errorf("Input: %v, Expected: %d, Got: %d", tt.nums, tt.expected, result)
		}
	}
}
