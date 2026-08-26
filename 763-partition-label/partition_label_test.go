package partitionlabel

import "testing"

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			s:    "eccbbbbdec",
			want: []int{10},
		},
	}

	for _, tc := range tests {
		got := partitionLabels(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("partitionLabels(%q) = %v; want %v", tc.s, got, tc.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
