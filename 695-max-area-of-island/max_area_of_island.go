package maxareaofisland

func MaxAreaOfIsland(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	var count func(r, c int) int
	count = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0
		a := count(r+1, c)
		b := count(r-1, c)
		e := count(r, c-1)
		d := count(r, c+1)
		return 1 + a + b + e + d
	}

	var maxArea int

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				result := count(r, c)
				maxArea = max(maxArea, result)
			}
		}
	}

	return maxArea
}
