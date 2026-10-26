package cheapestflightwithinkstop

import "testing"

func TestFindCheapestPrice(t *testing.T) {

	testCases := []struct {
		n        int
		flights  [][]int
		src      int
		dst      int
		k        int
		expected int
	}{

		// {
		// 	n:        3,
		// 	flights:  [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
		// 	src:      0,
		// 	dst:      2,
		// 	k:        1,
		// 	expected: 200,
		// },
		{
			n:        3,
			flights:  [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:      0,
			dst:      2,
			k:        0,
			expected: 500,
		},
	}

	for _, tt := range testCases {
		result := FindCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
		if result != tt.expected {
			t.Errorf("FindCheapestPrice(%d, %v, %d, %d, %d) = %d; expected %d", tt.n, tt.flights, tt.src, tt.dst, tt.k, result, tt.expected)
		}
	}
}
