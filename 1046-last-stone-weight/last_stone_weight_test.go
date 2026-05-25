package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{name: "test 1", stones: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{name: "test 2", stones: []int{1}, want: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lastStoneWeight(test.stones)
			if got != test.want {
				t.Errorf("LastStoneWeight(%v) = %d, want %d", test.stones, got, test.want)
			}
		})
	}
}
