package designdatastructureaddandsearchwords

import "testing"

func TestWordDictionary_Search(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		search string
		want   bool
	}{
		{
			name:   "exact match",
			words:  []string{"bad", "dad", "mad"},
			search: "bad",
			want:   true,
		},
		{
			name:   "no match",
			words:  []string{"bad", "dad", "mad"},
			search: "pad",
			want:   false,
		},
		{
			name:   "wildcard suffix",
			words:  []string{"bad", "dad", "mad"},
			search: ".ad",
			want:   true,
		},
		{
			name:   "wildcard prefix and suffix",
			words:  []string{"bad", "dad", "mad"},
			search: "b..",
			want:   true,
		},
		{
			name:   "prefix not a word",
			words:  []string{"apple"},
			search: "app",
			want:   false,
		},
		{
			name:   "longer than stored word",
			words:  []string{"bad"},
			search: "bade",
			want:   false,
		},
		{
			name:   "empty dictionary",
			words:  nil,
			search: "bad",
			want:   false,
		},
		{
			name:   "all wildcards",
			words:  []string{"abc"},
			search: "...",
			want:   true,
		},
		{
			name:   "wildcard no match",
			words:  []string{"abc"},
			search: "a.d",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd := Constructor()
			for _, word := range tt.words {
				wd.Insert(word)
			}
			if got := wd.Search(tt.search); got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.search, got, tt.want)
			}
		})
	}
}

func TestWordDictionary_LeetCodeExample(t *testing.T) {
	wd := Constructor()
	wd.Insert("bad")
	wd.Insert("dad")
	wd.Insert("mad")

	cases := []struct {
		word string
		want bool
	}{
		{"pad", false},
		{"bad", true},
		{".ad", true},
		{"b..", true},
	}

	for _, c := range cases {
		if got := wd.Search(c.word); got != c.want {
			t.Errorf("Search(%q) = %v, want %v", c.word, got, c.want)
		}
	}
}
