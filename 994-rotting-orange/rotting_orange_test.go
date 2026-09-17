package rottingorange

import "testing"

func TestRottinOrange(t *testing.T) {
	testTable := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 4,
		},
		{
			grid:     [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expected: -1,
		},
		{
			grid:     [][]int{{0, 2}},
			expected: 0,
		},
	}

	for _, tc := range testTable {
		if RottingOrange(tc.grid) != tc.expected {
			t.Errorf("RottingOrange(%v) = %v, want %v", tc.grid, RottingOrange(tc.grid), tc.expected)
		}
	}

}
