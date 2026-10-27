package longestcommonsubsequence

import "testing"

func TestLongestCommoneSubsequence(t *testing.T) {
	tests := []struct {
		name   string
		text1  string
		text2  string
		expect int
	}{
		{
			name:   "abcde and ace",
			text1:  "abcde",
			text2:  "ace",
			expect: 3,
		},
		{
			name:   "abc and adc",
			text1:  "abc",
			text2:  "adc",
			expect: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LongestCommoneSubsequence(test.text1, test.text2); got != test.expect {
				t.Errorf("LongestCommoneSubsequence(%q, %q) = %v; expected %v", test.text1, test.text2, got, test.expect)
			}
		})
	}
}
