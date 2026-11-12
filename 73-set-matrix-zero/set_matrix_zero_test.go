package setmatrixzero

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name:   "test case 1",
			input:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expect: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			SetZeroes(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expect) {
				t.Errorf("SetZeroes() = %v, want %v", tt.input, tt.expect)
			}
		})
	}
}
