package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	i, n := 0, len(intervals)

	// 1. insert all interval that the last value is less than new Interval index 0
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 2. Try to merge all the value that can be merge

	for i < n && intervals[i][1] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	// 4. insert all the remaining interval
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// the intuition for the solution is we have 3 step

// Example 1:

// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
// Example 2:

// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
