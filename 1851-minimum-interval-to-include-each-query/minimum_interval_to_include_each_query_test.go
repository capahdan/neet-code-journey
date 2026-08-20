package minimumintervaltoincludeeachquery

import (
	"reflect"
	"testing"
)

func TestMinInterval(t *testing.T) {
	tests := []struct {
		intervals [][]int
		queries   []int
		expected  []int
	}{
		{
			[][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
			[]int{2, 3, 4, 5},
			[]int{3, 3, 1, 4},
		},
		{
			[][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
			[]int{2, 19, 5, 22},
			[]int{2, -1, 4, 6},
		},
	}

	for _, test := range tests {
		result := minInterval(test.intervals, test.queries)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, result)
		}
	}
}
