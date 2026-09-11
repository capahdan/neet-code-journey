package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	testTable := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
	}

	for _, test := range testTable {
		actual := WordBreak(test.s, test.wordDict)
		if actual != test.expected {
			t.Errorf("WordBreak(%q, %v) = %v; expected %v", test.s, test.wordDict, actual, test.expected)
		}
	}

}
