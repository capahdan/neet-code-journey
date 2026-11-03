package longestincreasingpathinamatrix

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "single cell",
			matrix: [][]int{{1}},
			want:   1,
		},
		{
			name:   "single row increasing",
			matrix: [][]int{{1, 2, 3, 4}},
			want:   4,
		},
		{
			name:   "single row decreasing",
			matrix: [][]int{{4, 3, 2, 1}},
			want:   4,
		},
		{
			name:   "single column increasing",
			matrix: [][]int{{1}, {2}, {3}},
			want:   3,
		},
		{
			name: "leetcode example 1",
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			want: 4, // 1 -> 2 -> 6 -> 9
		},
		{
			name: "leetcode example 2",
			matrix: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			want: 4, // 3 -> 4 -> 5 -> 6
		},
		{
			name: "all equal values",
			matrix: [][]int{
				{7, 7, 7},
				{7, 7, 7},
			},
			want: 1,
		},
		{
			name: "spiral path covers whole matrix",
			matrix: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			want: 9,
		},

		{
			name: "zeros and positives",
			matrix: [][]int{
				{0, 1},
				{2, 3},
			},
			want: 3, // 0 -> 1 -> 3 or 0 -> 2 -> 3
		},
		{
			name: "non-square matrix",
			matrix: [][]int{
				{1, 2, 3, 4},
				{8, 7, 6, 5},
			},
			want: 8,
		},
		{
			name: "multiple equal-length paths",
			matrix: [][]int{
				{1, 2},
				{2, 3},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestIncreasingPath(tt.matrix)
			if got != tt.want {
				t.Errorf("LongestIncreasingPath() = %d, want %d", got, tt.want)
			}
		})
	}
}
