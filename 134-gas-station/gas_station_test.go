package gasstation

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	testcases := []struct {
		name   string
		gas    []int
		cost   []int
		expect int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"Example 2", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"Example 3", []int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := canCompleteCircuit(tc.gas, tc.cost)
			if result != tc.expect {
				t.Errorf("canCompleteCircuit(%v, %v) = %d; want %d", tc.gas, tc.cost, result, tc.expect)
			}
		})
	}
}
