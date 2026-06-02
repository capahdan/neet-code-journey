package combinationsum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{name: "test 1", candidates: []int{2, 3, 6, 7}, target: 7, want: [][]int{{2, 2, 3}, {7}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CombinationSum(test.candidates, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("CombinationSum(%v, %d) = %v, want %v", test.candidates, test.target, got, test.want)
			}
		})
	}
}
