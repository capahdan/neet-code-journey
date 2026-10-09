package mincosttoconnectallnodes

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	testCases := []struct {
		points   [][]int
		expected int
	}{
		// {
		// 	points:   [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, // 20
		// 	expected: 20,
		// },
		{
			points:   [][]int{{3, 12}, {-2, 5}, {-4, 1}}, // 18
			expected: 18,
		},
	}

	for _, tt := range testCases {
		result := minCostConnectPoints(tt.points)
		if result != tt.expected {
			t.Errorf("minCostConnectPoints(%v) = %d; expected %d", tt.points, result, tt.expected)
		}
	}
}
