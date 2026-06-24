package meetingroomii

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      int
	}{
		{
			name: "example 1",
			intervals: []Interval{
				{start: 0, end: 30},
				{start: 5, end: 10},
				{start: 15, end: 20},
			},
			want: 2,
		},
		{
			name: "example 2",
			intervals: []Interval{
				{start: 7, end: 10},
				{start: 2, end: 4},
			},
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minMeetingRooms(test.intervals)
			if got != test.want {
				t.Errorf("minMeetingRooms(%v) = %d, want %d", test.intervals, got, test.want)
			}
		})
	}
}
