package numbersisland

import (
	"reflect"
	"testing"
)

func TestNumberIsland(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]byte
		expect int
	}{
		{
			name: "",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberIsland(tt.grid)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("NumberIsland() = %v, want %v", got, tt.expect)
			}
		})
	}
}
