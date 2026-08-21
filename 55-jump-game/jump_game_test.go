package jumpgame

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{1, 0, 1, 0}, false},
		{[]int{2, 0, 0}, true},
	}

	for _, test := range tests {
		result := CanJump(test.nums)
		if result != test.expected {
			t.Errorf("CanJump(%v) = %v, expected %v", test.nums, result, test.expected)
		}
	}
}
