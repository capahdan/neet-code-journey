package islandandtreasure

const INF = 2147483647

type pair struct{ r, c int }

func islandsAndTreasure(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	rows := len(grid)
	cols := len(grid[0])

	queue := make([]pair, 0, rows*cols)

	// Seed the queue with every treasure chest (multi-source BFS).
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, pair{r, c})
			}
		}
	}

	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nr, nc := curr.r+d.r, curr.c+d.c

			// Skip out-of-bounds, walls, chests, or already-visited land.
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != INF {
				continue
			}

			grid[nr][nc] = grid[curr.r][curr.c] + 1
			queue = append(queue, pair{nr, nc})
		}
	}
}
