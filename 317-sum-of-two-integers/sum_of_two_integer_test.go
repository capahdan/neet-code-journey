package sumoftwointegers

import "testing"

func TestSumOfTwoInteger(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			a:    1,
			b:    2,
			want: 3,
		},
		{
			a:    2,
			b:    3,
			want: 5,
		},
		{
			a:    0,
			b:    0,
			want: 0,
		},
		{
			a:    -1,
			b:    1,
			want: 0,
		},
		{
			a:    -2,
			b:    -3,
			want: -5,
		},
		{
			a:    10,
			b:    20,
			want: 30,
		},
	}

	for _, tt := range tests {
		if got := GetSum(tt.a, tt.b); got != tt.want {
			t.Errorf("GetSum(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
