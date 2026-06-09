package generateparentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "test 1", n: 3, want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generateParenthesis(test.n)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("generateParenthesis(%d) = %v, want %v", test.n, got, test.want)
			}
		})
	}
}
