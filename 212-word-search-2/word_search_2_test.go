package wordsearch2

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{name: "Test Find Words", board: [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'}},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"oath", "eat"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findWords(test.board, test.words)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("findWords(%v, %v) = %v, want %v", test.board, test.words, got, test.want)
			}
		})
	}
}
