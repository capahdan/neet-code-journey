package longestincreasingsubsequence

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	testTable := []struct {
		nums     []int
		expected int
	}{
		// {
		// 	nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
		// 	expected: 4,
		// },
		{
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		// {
		// 	nums:     []int{7, 7, 7, 7, 7, 7, 7},
		// 	expected: 1,
		// },
	}

	for _, test := range testTable {
		actual := LengthOfLIS(test.nums)
		if actual != test.expected {
			t.Errorf("LengthOfLIS(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}

}
