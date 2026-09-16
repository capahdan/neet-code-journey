package islandandtreasure

import (
	"reflect"
	"testing"
)

func TestIslandsAndTreasure(t *testing.T) {
	testTable := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			expected: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			matrix: [][]int{
				{0, -1},
				{INF, INF},
			},
			expected: [][]int{
				{0, -1},
				{1, 2},
			},
		},
	}

	for _, tc := range testTable {
		islandsAndTreasure(tc.matrix)
		if !reflect.DeepEqual(tc.matrix, tc.expected) {
			t.Errorf("islandsAndTreasure() = %v, want %v", tc.matrix, tc.expected)
		}
	}
}
