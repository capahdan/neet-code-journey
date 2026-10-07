package reconstructitenary

import "container/heap"

type StringHeap []string

func (s StringHeap) Len() int {
	return len(s)
}
func (s StringHeap) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s StringHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s *StringHeap) Push(x interface{}) {
	*s = append(*s, x.(string))
}
func (s *StringHeap) Pop() interface{} {
	old := *s
	n := len(old)
	item := old[n-1]
	*s = old[:n-1]
	return item
}

// the solution works because we use min heap to explore the graph
// after that we then pop the heap and using dfs to explore every possible solution
// finnaly we append the route in the result

func FindItenary(tickets [][]string) []string {
	graph := make(map[string]*StringHeap)
	for _, t := range tickets {
		if graph[t[0]] == nil {
			graph[t[0]] = &StringHeap{}
		}
		heap.Push(graph[t[0]], t[1])
	}

	route := []string{}

	var dfs func(string)
	dfs = func(airport string) {
		h := graph[airport]
		for h != nil && h.Len() > 0 {
			nxt := heap.Pop(h).(string)
			dfs(nxt)
		}
		route = append(route, airport)
	}

	dfs("JFK")

	for i := 0; i < len(route)/2; i++ {
		route[i], route[len(route)-i-1] = route[len(route)-i-1], route[i]
	}

	return route
}
