package swininrisingwater

import "testing"

func TestSwimInWater(t *testing.T) {
	testCases := []struct {
		grid     [][]int
		expected int
	}{
		// {
		// 	grid:     [][]int{{0, 2}, {1, 3}}, // 3
		// 	expected: 3,
		// },
		{
			grid:     [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, // 16
			expected: 16,
		},
	}

	for _, tt := range testCases {
		result := swimInWater(tt.grid)
		if result != tt.expected {
			t.Errorf("swimInWater(%v) = %d; expected %d", tt.grid, result, tt.expected)
		}
	}
}
