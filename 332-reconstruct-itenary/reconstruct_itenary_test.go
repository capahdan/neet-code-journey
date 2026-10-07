package reconstructitenary

import "testing"

func TestFindItenary(t *testing.T) {
	testTable := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		// {
		// 	name:     "test 1",
		// 	tickets:  [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
		// 	expected: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		// },
		{
			name:     "test 2",
			tickets:  [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			expected: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		// {
		// 	name:     "test 3",
		// 	tickets:  [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
		// 	expected: []string{"JFK", "NRT", "JFK", "KUL"},
		// },
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			result := FindItenary(tc.tickets)
			if !areEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func areEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
