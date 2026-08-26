package mergetripletstoformtargettriplet

import "testing"

func TestMergeTriplets(t *testing.T) {
	tests := []struct {
		triplets [][]int
		target   []int
		want     bool
	}{
		{
			triplets: [][]int{{2, 5, 3}, {4, 3, 1}, {1, 7, 5}},
			target:   []int{2, 7, 5},
			want:     true,
		},
		{
			triplets: [][]int{{3, 4, 5}, {4, 5, 6}}, //
			target:   []int{2, 3, 4},
			want:     false,
		},
		// {
		// 	triplets: [][]int{{5, 1, 3}, {5, 1, 5}, {5, 1, 5}},
		// 	target:   []int{5, 1, 3},
		// 	want:     true,
		// },
	}

	for _, tc := range tests {
		if got := mergeTriplets(tc.triplets, tc.target); got != tc.want {
			t.Errorf("mergeTriplets(%v, %v) = %v; want %v", tc.triplets, tc.target, got, tc.want)
		}
	}

}
