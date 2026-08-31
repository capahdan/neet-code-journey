package mincostclimbingstairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "example1",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "example2",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.cost); got != tt.expected {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.expected)
			}
		})
	}
}
