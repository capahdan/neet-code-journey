package redundantconnection

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {

	tests := []struct {
		edges  [][]int
		expect []int
	}{
		{
			edges:  [][]int{{1, 2}, {1, 3}, {2, 3}},
			expect: []int{2, 3},
		},
		{
			edges:  [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			expect: []int{1, 4},
		},
	}

	for i, tc := range tests {
		result := FindRedundantConnection(tc.edges)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, tc.expect, result)
		}
	}

}
