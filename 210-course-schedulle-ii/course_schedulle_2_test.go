package courseschedulleii

import "testing"

func TestCourseSchedulle(t *testing.T) {
	test := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			name:          "test-1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			name:          "test-2",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			name:          "test-3",
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      []int{0},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := FindOrder(tc.numCourses, tc.prerequisites)
			if !compareSlice(result, tc.expected) {
				t.Errorf("FindOrder(%d, %v) = %v; want %v", tc.numCourses, tc.prerequisites, result, tc.expected)
			}
		})
	}
}

func compareSlice(a, b []int) bool {
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
