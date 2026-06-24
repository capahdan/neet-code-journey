package meetingroomii

import "sort"

type Interval struct {
	start int
	end   int
}

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	n := len(intervals)
	start := make([]int, n)
	end := make([]int, n)

	for i, interval := range intervals {
		start[i] = interval.start
		end[i] = interval.end
	}

	sort.Ints(start)
	sort.Ints(end)

	res, count := 0, 0
	s, e := 0, 0

	for s < n {
		if start[s] < end[e] {
			s++
			count++
		} else {
			e++
			count--
		}
		if count > res {
			res = count
		}
	}

	return res
}
