package rottingorange

// this function is need to count the number of minute is needed to make all the orange rotted
// 2 = rotted
// 1 = fresh
// 0 = non transverable
// we will use breadFirstSearch because we need to process all the rotted orange per level

type pair struct {
	row int
	col int
}

func RottingOrange(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var minutes int
	var freshOrange int
	var queue []pair
	rows := len(grid)
	cols := len(grid[0])

	directions := []pair{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				freshOrange++
			} else if grid[i][j] == 2 {
				queue = append(queue, pair{
					row: i,
					col: j,
				})
			}
		}
	}

	if freshOrange == 0 {
		return 0
	}

	for len(queue) > 0 && freshOrange > 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				newRow := curr.row + dir.row
				newCol := curr.col + dir.col

				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || grid[newRow][newCol] != 1 {
					continue
				}

				grid[newRow][newCol] = 2
				queue = append(queue, pair{
					row: newRow,
					col: newCol,
				})
				freshOrange--
			}

		}
		minutes++
	}

	if freshOrange > 0 {
		return -1
	}
	return minutes
}
