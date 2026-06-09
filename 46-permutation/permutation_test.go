package permutation

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "test 1", nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := permute(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("permute(%v) = %v, want %v", test.nums, got, test.want)
			}
		})
	}
}
