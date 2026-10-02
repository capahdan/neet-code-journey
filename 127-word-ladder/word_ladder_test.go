package wordladder

import "testing"

func TestLadderLenght(t *testing.T) {
	tests := []struct {
		startWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
	}

	for _, test := range tests {
		if result := LadderLenght(test.startWord, test.endWord, test.wordList); result != test.expected {
			t.Errorf("LadderLenght(%s, %s, %v) = %d; expected %d", test.startWord, test.endWord, test.wordList, result, test.expected)
		}
	}

}
