package maximumsubarray

import (
	"testing"
)

func TestMaximumSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		{
			name:     "Mixed numbers",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "Single number",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Array with zeros",
			nums:     []int{0, 1, 0, 2, 0, 3, 0},
			expected: 6,
		},
		{
			name:     "Array with large numbers",
			nums:     []int{1000, -1000, 2000, -2000, 3000},
			expected: 3000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximumSubArray(tt.nums)
			if result != tt.expected {
				t.Errorf("MaximumSubArray(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
