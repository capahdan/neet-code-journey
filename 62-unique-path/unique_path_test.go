package uniquepath

import "testing"

func TestUniqueQueue(t *testing.T) {
	testTable := []struct {
		m        int
		n        int
		expected int
	}{
		// {
		// 	m:        2,
		// 	n:        2,
		// 	expected: 2,
		// },
		{
			m:        3,
			n:        4,
			expected: 10,
		},
	}
	for _, testData := range testTable {
		actual := UniquePath(testData.m, testData.n)
		if actual != testData.expected {
			t.Errorf("Expected %d, got %d", testData.expected, actual)
		}
	}
}
