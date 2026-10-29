package besttimetobuystockwithcooldown

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		// {
		// 	name:   "Example 2",
		// 	prices: []int{1},
		// 	want:   0,
		// },
		// {
		// 	name:   "Example 3",
		// 	prices: []int{1, 2},
		// 	want:   1,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
