package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {

	testTable := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "test 2",
			s:    "cbbd",
			want: "bb",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("longestPalindrome(%v) = %v, want %v", tt.s, got, tt.want)

			}
		})
	}

}
