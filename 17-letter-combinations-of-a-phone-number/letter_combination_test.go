package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{name: "test 1", digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := LetterCombinations(test.digits)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("LetterCombinations(%s) = %v, want %v", test.digits, got, test.want)
			}
		})
	}
}
