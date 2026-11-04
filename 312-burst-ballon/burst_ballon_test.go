package burstballon

import "testing"

func Test_maxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test Case 1",
			nums: []int{3, 1, 5, 8},
			want: 167,
		},
		{
			name: "Test Case 2",
			nums: []int{1, 5},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
