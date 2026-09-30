package numberofcomponentinundirectedgraph

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		n      int
		edges  [][]int
		wanted int
	}{
		{n: 5, edges: [][]int{{0, 1}, {1, 2}, {3, 4}}, wanted: 2},
		{n: 5, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, wanted: 1},
	}

	for _, tt := range tests {
		got := CountComponents(tt.n, tt.edges)
		if got != tt.wanted {
			t.Errorf("CountComponents(%d, %v) = %d; wanted %d", tt.n, tt.edges, got, tt.wanted)
		}
	}
}
