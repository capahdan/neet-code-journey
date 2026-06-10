package wordsearch

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{name: "test 1", board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ABCCED", want: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := exist(test.board, test.word)
			if got != test.want {
				t.Errorf("exist(%v, %v) = %v, want %v", test.board, test.word, got, test.want)
			}
		})
	}
}
