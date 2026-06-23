package nonoverlappinginterval

import (
	"testing"
)

func TestNonOverlappingInterval(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  int
	}{
		// {[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		// {[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		// {[][]int{{1, 2}, {2, 3}}, 0},
		// {[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		// {[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, 2},
	}

	for _, test := range tests {
		result := NonOverlappingInterval(test.intervals)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}
