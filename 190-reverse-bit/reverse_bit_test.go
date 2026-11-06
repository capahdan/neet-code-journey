package reversebit

import "testing"

func TestReverseBit(t *testing.T) {
	testTable := []struct {
		n        int
		expected int
	}{
		{
			n:        43261596,
			expected: 964176192,
		},
	}

	for _, tt := range testTable {
		actual := reverseBit(tt.n)
		if actual != tt.expected {
			t.Errorf("Input: %d, Expected: %d, Got: %d", tt.n, tt.expected, actual)
		}
	}
}
