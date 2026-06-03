package combinationsumii

import (
	"reflect"
	"testing"
)

func TestCombinationSumII(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{name: "test 1", candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CombinationSumII(test.candidates, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("CombinationSumII(%v, %d) = %v, want %v", test.candidates, test.target, got, test.want)
			}
		})
	}
}
