package nonoverlappinginterval

import "sort"

func NonOverlappingInterval(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	overLapCount := 0
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			// Overlap: remove current interval (it has a later end since we sorted by end)
			overLapCount++
			// prevEnd stays the same — we keep the earlier-ending interval

		} else {
			// No overlap: update prevEnd to current interval's end
			prevEnd = intervals[i][1]
		}
	}

	return overLapCount
}
