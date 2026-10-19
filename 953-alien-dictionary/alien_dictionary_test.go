package aliendictionary

import (
	"testing"
)

func TestAlienDictionary(t *testing.T) {
	tests := []struct {
		words    []string
		expected string
	}{
		// {[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		// {[]string{"z", "x"}, "zx"},
		{[]string{"z", "x", "z"}, ""},
		// {[]string{"abc", "ab"}, ""},
	}

	for _, test := range tests {
		actual := ForeignDictionary(test.words)
		if actual != test.expected {
			t.Errorf("Input: %v, Expected: %s, Got: %s", test.words, test.expected, actual)
		}
	}
}
