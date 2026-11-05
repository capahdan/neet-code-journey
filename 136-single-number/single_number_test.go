package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	testTable := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tt := range testTable {
		actual := SingleNumber(tt.nums)
		if actual != tt.expected {
			t.Errorf("Input: %v, Expected: %d, Got: %d", tt.nums, tt.expected, actual)
		}
	}
}
