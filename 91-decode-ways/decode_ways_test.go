package decodeways

import "testing"

func TestNumDecoding(t *testing.T) {
	tests := []struct {
		s      string
		wanted int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"27", 1},
	}

	for _, test := range tests {
		got := NumDecoding(test.s)
		if got != test.wanted {
			t.Errorf("NumDecoding(%q) = %d; wanted %d", test.s, got, test.wanted)
		}
	}
}
