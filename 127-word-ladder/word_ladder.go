package wordladder

// what did i learn from this problem is we need to do bfs by level so that
// we don't process newly added data in the queue
// we could do this by maintaining size of queue at the start of inner loop

// futhermore, i learn that we can use pattern for finding the neighbors data

func LadderLenght(startWord, endWord string, wordList []string) int {

	found := false

	for _, word := range wordList {
		if word == endWord {
			found = true
			break
		}
	}

	if !found {
		return 0
	}

	wordList = append(wordList, startWord)
	nei := make(map[string][]string)
	for _, word := range wordList {
		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]
			nei[pattern] = append(nei[pattern], word)
		}
	}

	visited := make(map[string]bool)
	visited[startWord] = true
	queue := []string{startWord}
	result := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endWord {
				return result
			}

			for j := range curr {
				pattern := curr[:j] + "*" + curr[j+1:]
				for _, neighbor := range nei[pattern] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}

		}

		result += 1
	}

	return 0
}
