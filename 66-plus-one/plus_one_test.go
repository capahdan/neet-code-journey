package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {

	testTable := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			input:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			input:    []int{9},
			expected: []int{1, 0},
		},
		{
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
	}

	for _, tt := range testTable {

		res := PlusOne(tt.input)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("PlusOne(%v) = %v; expected %v", tt.input, res, tt.expected)
		}

	}

}
