package kclosestpointstoorigin

import (
	"reflect"
	"testing"
)

func TestKClosestPointsToOrigin(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{name: "test 1", points: [][]int{{1, 3}, {3, 4}, {2, -1}}, k: 2, want: [][]int{{1, 3}, {2, -1}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := kClosest(test.points, test.k)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("KClosestPointsToOrigin(%v, %d) = %v, want %v", test.points, test.k, got, test.want)
			}
		})
	}
}
