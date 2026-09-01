package houserobber

import "testing"

func TestHouseRobber(t *testing.T) {
	testTable := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			input:    []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			input:    []int{2, 1, 1, 2},
			expected: 4,
		},
		{
			input:    []int{1, 2},
			expected: 2,
		},
		{
			input:    []int{1},
			expected: 1,
		},
		{
			input:    []int{1, 1, 1, 1, 1},
			expected: 3,
		},
	}

	for _, tt := range testTable {
		result := Rob(tt.input)
		if result != tt.expected {
			t.Errorf("Input: %v, Expected: %d, Got: %d", tt.input, tt.expected, result)
		}
	}

}
