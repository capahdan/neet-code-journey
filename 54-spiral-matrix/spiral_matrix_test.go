package spiralmatrix

import "testing"

func TestSpiralOrder(t *testing.T) {
	testcases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "3x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "1x1 matrix",
			input: [][]int{
				{1},
			},
			expected: []int{1},
		},
		{
			name: "single row matrix",
			input: [][]int{
				{1, 2, 3, 4},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "single column matrix",
			input: [][]int{
				{1},
				{2},
				{3},
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := SpiralOrder(tc.input)
			if !equalSlices(got, tc.expected) {
				t.Errorf("spiralOrder(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

// helper to compare two integer slices
func equalSlices(a, b []int) bool {
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
