package swininrisingwater

// import "container/heap"

// type item struct {
// 	x        int
// 	y        int
// 	distance int
// }

// type minHeap []item

// func (m minHeap) Len() int           { return len(m) }
// func (m minHeap) Less(i, j int) bool { return m[i].distance < m[j].distance }
// func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

// func (m *minHeap) Push(x any) {
// 	*m = append(*m, x.(item))
// }

// func (m *minHeap) Pop() any {
// 	old := *m
// 	n := len(old)
// 	x := old[n-1]
// 	*m = old[:n-1]
// 	return x
// }

// func swimInWater(grid [][]int) int {

// 	rows := len(grid)
// 	cols := len(grid[0])

// 	graph := make(map[[2]int][]item)

// 	for i := range grid {
// 		for j := range grid {
// 			key := [2]int{i, j}
// 			if i-1 > 0 {
// 				graph[key] = append(graph[key], item{
// 					x:        i - 1,
// 					y:        j,
// 					distance: grid[i-1][j],
// 				})
// 			}

// 			if i+1 < rows {
// 				graph[key] = append(graph[key], item{
// 					x:        i + 1,
// 					y:        j,
// 					distance: grid[i+1][j],
// 				})
// 			}

// 			if j-1 > 0 {
// 				graph[key] = append(graph[key], item{
// 					x:        i,
// 					y:        j - 1,
// 					distance: grid[i][j-1],
// 				})
// 			}

// 			if j+1 < cols {
// 				graph[key] = append(graph[key], item{
// 					x:        i,
// 					y:        j + 1,
// 					distance: grid[i][j+1],
// 				})
// 			}
// 		}
// 	}

// 	visited := make(map[[2]int]int)
// 	h := &minHeap{}

// 	heap.Push(h, item{
// 		x:        0,
// 		y:        0,
// 		distance: grid[0][0],
// 	})

// 	lastX := len(grid) - 1
// 	lastY := len(grid[0]) - 1

// 	for h.Len() > 0 {
// 		cur := heap.Pop(h).(item)

// 		key := [2]int{cur.x, cur.y}
// 		if _, ok := visited[key]; ok {
// 			continue
// 		}

// 		visited[key] = cur.distance
// 		if cur.x == lastX && cur.y == lastY {
// 			break
// 		}

// 		for _, node := range graph[key] {
// 			nodeKey := [2]int{node.x, node.y}
// 			if _, ok := visited[nodeKey]; !ok {
// 				heap.Push(h, item{
// 					x:        node.x,
// 					y:        node.y,
// 					distance: node.distance,
// 				})
// 			}
// 		}

// 	}

// 	var result int
// 	for _, value := range visited {
// 		result = max(result, value)
// 	}

// 	return result
// }
import (
	"container/heap"
)

type item struct {
	t, r, c int
}

type minHeap []item

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].t < m[j].t }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(item)) }
func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	h := &minHeap{{grid[0][0], 0, 0}}

	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true

	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.r == n-1 && cur.c == n-1 {
			return cur.t
		}

		for _, d := range dirs {
			neiR, neiC := cur.r+d[0], cur.c+d[1]
			if neiR < 0 || neiR == n || neiC < 0 || neiC >= n || visited[[2]int{neiR, neiC}] {
				continue
			}

			visited[[2]int{neiR, neiC}] = true
			heap.Push(h, item{
				t: max(cur.t, grid[neiR][neiC]),
				r: neiR,
				c: neiC,
			})
		}

	}

	return -1
}
