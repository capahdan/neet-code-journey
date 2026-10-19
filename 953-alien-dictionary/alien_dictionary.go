package aliendictionary

func ForeignDictionary(words []string) string {

	// 1. collect all unique characters
	adj := make(map[byte]map[byte]bool)
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := adj[w[i]]; !ok {
				adj[w[i]] = make(map[byte]bool)
			}
		}
	}

	// 2. build edges by comparing adjacent words
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}

		// if w1 is longer than w2 but w2 is the prefix of w1
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}

		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				adj[w1[j]][w2[j]] = true
				break
			}
		}
	}

	// 3. DFS : topological sort with cycle detection
	// visited[c] == false -> finished (not in current path)
	// visited[c] == true  -> currently in the DFS path

	visited := make(map[byte]bool)
	res := []byte{}

	var dfs func(c byte) bool
	dfs = func(c byte) bool {
		if inPath, seen := visited[c]; seen {
			return inPath
		}
		visited[c] = true

		for nei := range adj[c] {
			if dfs(nei) {
				return true
			}
		}

		visited[c] = false
		res = append(res, c)
		return false
	}

	for c := range adj {
		if dfs(c) {
			return ""
		}
	}

	i, j := 0, len(res)-1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return string(res)
}
