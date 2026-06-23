package mergeinterval

import "sort"

func MergeInterval(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		n := len(result)

		if len(result) == 0 || result[n-1][1] < interval[0] {
			result = append(result, interval)
		} else {
			if interval[1] > result[n-1][1] {
				result[n-1][1] = interval[1]
			}
		}
	}

	return result
}
