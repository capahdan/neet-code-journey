package courseschedulle

import "testing"

func TestCanFinish(t *testing.T) {
	testTable := []struct {
		name          string
		numCourse     int
		prerequisites [][]int
		expected      bool
	}{
		// {
		// 	name:          "success case",
		// 	numCourse:     2,
		// 	prerequisites: [][]int{{1, 0}},
		// 	expected:      true,
		// },
		// {
		// 	name:          "fail case",
		// 	numCourse:     2,
		// 	prerequisites: [][]int{{1, 0}, {0, 1}},
		// 	expected:      false,
		// },
		{
			name:          "success case 2",
			numCourse:     3,
			prerequisites: [][]int{{2, 1}, {1, 0}},
			expected:      true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			result := CanFinish(tc.numCourse, tc.prerequisites)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}

}
