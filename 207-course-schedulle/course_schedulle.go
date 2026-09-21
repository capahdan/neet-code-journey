package courseschedulle

//  the inttution behind this function is that we need to make
// a graph and find if there is a cycle in it. If there is a cycle in it then we can't finish all the courses.
// in this case the graph is directed and the edges are from the prerequisite to the course.

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// map each course to its prereq list
	preMap := make([][]int, numCourses)
	for _, p := range prerequisites {
		crs, pre := p[0], p[1]
		preMap[crs] = append(preMap[crs], pre)
	}

	// visitSet = all courses along the current DFS path
	visitSet := make(map[int]bool)

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if visitSet[crs] {
			return false // already on my path: loop
		}
		if len(preMap[crs]) == 0 {
			return true // no prereqs left: safe
		}

		visitSet[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visitSet, crs) // leaving this path
		preMap[crs] = nil     // mark as fully checked
		return true
	}

	for crs := range numCourses {
		if !dfs(crs) {
			return false
		}
	}
	return true
}
