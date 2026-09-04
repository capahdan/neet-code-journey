package palindromicsubstring

import "testing"

func TestPalindromicSubstring(t *testing.T) {

	testTable := []struct {
		name  string
		s     string
		count int
	}{
		{
			name:  "test 1",
			s:     "abc",
			count: 3,
		},
		{
			name:  "test 2",
			s:     "aaa",
			count: 6,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := CountSubstring(tt.s)
			if got != tt.count {
				t.Errorf("CountSubstring(%v) = %d, want %d", tt.s, got, tt.count)

			}
		})
	}

}
