package coinchangeii

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{
			name:   "Example 1",
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
		// {
		// 	name:   "Example 2",
		// 	amount: 3,
		// 	coins:  []int{2},
		// 	want:   0,
		// },
		// {
		// 	name:   "Example 3",
		// 	amount: 10,
		// 	coins:  []int{10},
		// 	want:   1,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Change(tt.amount, tt.coins); got != tt.want {
				t.Errorf("Change() = %v, want %v", got, tt.want)
			}
		})
	}
}
