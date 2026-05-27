package subset

import (
	"reflect"
	"testing"
)

func TestSubset(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "test 1", nums: []int{1, 2, 3}, want: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Subset(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Subset(%v) = %v, want %v", test.nums, got, test.want)
			}
		})
	}
}
