package handofstraight

import "testing"

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		name      string
		hand      []int
		groupSize int
		want      bool
	}{
		{
			name:      "example 1 - can form consecutive groups",
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			want:      true,
		},
		// {
		// 	name:      "example 2 - cannot form consecutive groups",
		// 	hand:      []int{1, 2, 3, 4, 5},
		// 	groupSize: 4,
		// 	want:      false,
		// },
		// {
		// 	name:      "length not divisible by groupSize",
		// 	hand:      []int{1, 2, 3, 4, 5},
		// 	groupSize: 2,
		// 	want:      false,
		// },
		// {
		// 	name:      "group size of 1 always works",
		// 	hand:      []int{5, 3, 1, 4, 2},
		// 	groupSize: 1,
		// 	want:      true,
		// },
		// {
		// 	name:      "single group exactly consecutive",
		// 	hand:      []int{1, 2, 3},
		// 	groupSize: 3,
		// 	want:      true,
		// },
		// {
		// 	name:      "gap prevents last group from being consecutive",
		// 	hand:      []int{1, 2, 3, 4, 6, 7, 8, 10},
		// 	groupSize: 4,
		// 	want:      false,
		// },
		// {
		// 	name:      "negative numbers can form groups",
		// 	hand:      []int{-2, -1, 0, 1, 2, 3},
		// 	groupSize: 3,
		// 	want:      true,
		// },
		// {
		// 	name:      "many duplicates forming multiple groups",
		// 	hand:      []int{1, 1, 2, 2, 3, 3},
		// 	groupSize: 3,
		// 	want:      true,
		// },
		// {
		// 	name:      "duplicate count exceeds available run",
		// 	hand:      []int{1, 1, 1, 2, 2, 3, 3},
		// 	groupSize: 3,
		// 	want:      false,
		// },
		// {
		// 	name:      "group size larger than hand",
		// 	hand:      []int{1, 2, 3},
		// 	groupSize: 4,
		// 	want:      false,
		// },
		// {
		// 	name:      "empty hand",
		// 	hand:      []int{},
		// 	groupSize: 1,
		// 	want:      true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNStraightHand(tt.hand, tt.groupSize)
			if got != tt.want {
				t.Errorf("isNStraightHand(%v, %d) = %v, want %v", tt.hand, tt.groupSize, got, tt.want)
			}
		})
	}
}
