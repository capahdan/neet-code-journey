package graphvalidtree

import "testing"

func TestGraphValidTree(t *testing.T) {
	testTable := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{

		{
			name:     "test-1",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			expected: true,
		},
		{
			name:     "test-2",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			expected: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidTree(tc.n, tc.edges)
			if result != tc.expected {
				t.Errorf("ValidTree(%d, %v) = %v; want %v", tc.n, tc.edges, result, tc.expected)
			}
		})
	}

}
