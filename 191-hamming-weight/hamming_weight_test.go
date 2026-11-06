package hummingweight

import "testing"

func TestHammingWeight(t *testing.T) {
	testTable := []struct {
		n        int
		expected int
	}{
		{
			n:        11,
			expected: 3,
		},
		{
			n:        128,
			expected: 1,
		},
		{
			n:        5,
			expected: 2,
		},
		{
			n:        0,
			expected: 0,
		},
	}

	for _, tt := range testTable {
		actual := hammingWeight(tt.n)
		if actual != tt.expected {
			t.Errorf("Input: %d, Expected: %d, Got: %d", tt.n, tt.expected, actual)
		}
	}
}
