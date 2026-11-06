package countbit

import "testing"

func TestCountBit(t *testing.T) {
	testTable := []struct {
		n        int
		expected []int
	}{
		{
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range testTable {
		actual := countBits(tt.n)
		if !equal(actual, tt.expected) {
			t.Errorf("Input: %d, Expected: %v, Got: %v", tt.n, tt.expected, actual)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
