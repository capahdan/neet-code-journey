package distinctsubsequences

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "basic example 1",
			s:    "rabbbit",
			t:    "rabbit",
			want: 3,
		},
		{
			name: "basic example 2",
			s:    "babgbag",
			t:    "bag",
			want: 5,
		},
		{
			name: "empty target",
			s:    "abc",
			t:    "",
			want: 1, // One way to form an empty subsequence
		},
		{
			name: "empty source",
			s:    "",
			t:    "abc",
			want: 0,
		},
		{
			name: "identical strings",
			s:    "abc",
			t:    "abc",
			want: 1,
		},
		{
			name: "source shorter than target",
			s:    "ab",
			t:    "abc",
			want: 0,
		},
		{
			name: "no matching characters",
			s:    "abc",
			t:    "d",
			want: 0,
		},
		{
			name: "multiple occurrences",
			s:    "aaaaa",
			t:    "aa",
			want: 10, // C(5, 2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDistinct(tt.s, tt.t); got != tt.want {
				t.Errorf("NumDistinct(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
