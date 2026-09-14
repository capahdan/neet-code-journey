package maxareaofisland

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {

	testCases := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{"Example 1", [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}, 4},
		{"Example 2", [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}, 0},
		{"Example 3", [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}, 20},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxAreaOfIsland(tc.grid); got != tc.expect {
				t.Errorf("MaxAreaOfIsland() = %v, want %v", got, tc.expect)
			}
		})
	}

}
