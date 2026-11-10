package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "empty matrix",
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			name:     "1x1 matrix",
			input:    [][]int{{1}},
			expected: [][]int{{1}},
		},
		{
			name:     "2x2 matrix",
			input:    [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{3, 1}, {4, 2}},
		},
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "4x4 matrix",
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "matrix with negative and repeated values",
			input: [][]int{
				{-1, 0},
				{0, -1},
			},
			expected: [][]int{
				{0, -1},
				{-1, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("rotate() = %v, want %v", tt.input, tt.expected)
			}
		})
	}
}

// Rotating four times by 90° must give back the original matrix.
func TestRotateImageFourTimesIsIdentity(t *testing.T) {
	original := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// deep copy, since rotate works in place
	m := make([][]int, len(original))
	for i := range original {
		m[i] = append([]int(nil), original[i]...)
	}

	for i := 0; i < 4; i++ {
		rotate(m)
	}

	if !reflect.DeepEqual(m, original) {
		t.Errorf("after 4 rotations got %v, want %v", m, original)
	}
}
