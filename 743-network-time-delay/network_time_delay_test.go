package networktimedelay

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	testTable := []struct {
		name     string
		times    [][]int
		n        int
		k        int
		expected int
	}{
		{"test 1", [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{"test 2", [][]int{{1, 2, 1}}, 2, 1, 1},
		{"test 3", [][]int{{1, 2, 1}}, 2, 2, -1},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			result := NetworkDelayTime(tc.times, tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
