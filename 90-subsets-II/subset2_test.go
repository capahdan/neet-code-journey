package subsetsii

import (
	"reflect"
	"testing"
)

func TestSubset2(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "test 1", nums: []int{1, 2, 2}, want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SubsetsII(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Subset2(%v) = %v, want %v", test.nums, got, test.want)
			}
		})
	}
}
