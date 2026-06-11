package palindromepartitioning

import (
	"reflect"
	"testing"
)

func TestPalindromePartitioning(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{name: "test 1", s: "aab", want: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{name: "test 2", s: "a", want: [][]string{{"a"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := partition(test.s)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("partition(%s) = %v, want %v", test.s, got, test.want)
			}
		})
	}
}
