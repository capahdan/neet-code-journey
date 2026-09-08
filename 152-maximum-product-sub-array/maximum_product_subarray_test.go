package maximumproductsubarray

import "testing"

func TestMaximumProductSubarray(t *testing.T) {

	testTable := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
	}

	for _, tt := range testTable {
		res := MaxProduct(tt.nums)
		if res != tt.expected {
			t.Errorf("got %d, expected %d", res, tt.expected)
		}
	}
}
