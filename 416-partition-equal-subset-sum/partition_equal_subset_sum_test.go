package partitionequalsubsetsum

import "testing"

func TestCanPartition(t *testing.T) {
	testTable := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			nums:     []int{2, 2, 3, 5},
			expected: false,
		},
	}

	for _, test := range testTable {
		actual := CanPartition(test.nums)
		if actual != test.expected {
			t.Errorf("CanPartition(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}
