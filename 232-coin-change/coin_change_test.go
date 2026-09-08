package coinchange

import "testing"

func TestCoinchange(t *testing.T) {

	TestTable := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Test Case 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		// {
		// 	name:     "Test Case 2",
		// 	coins:    []int{2},
		// 	amount:   3,
		// 	expected: -1,
		// },
		// {
		// 	name:     "Test Case 3",
		// 	coins:    []int{1},
		// 	amount:   0,
		// 	expected: 0,
		// },
	}

	for _, tc := range TestTable {
		t.Run(tc.name, func(t *testing.T) {
			if got := coinChange(tc.coins, tc.amount); got != tc.expected {
				t.Errorf("coinchange(%v, %d) = %d; want %d", tc.coins, tc.amount, got, tc.expected)
			}
		})
	}

}
