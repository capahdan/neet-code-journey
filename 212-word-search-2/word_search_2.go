package wordsearch2

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	return Trie{root: &TrieNode{}}
}

type TrieNode struct {
	Children [26]*TrieNode
	Word     string
}

func (this *Trie) Add(word string) {
	node := this.root
	for _, char := range word {
		index := char - 'a'
		if node.Children[index] == nil {
			node.Children[index] = &TrieNode{}
		}
		node = node.Children[index]
	}

	node.Word = word
}

func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()

	// Step 1: insert all words into the trie
	for _, word := range words {
		trie.Add(word)
	}

	// Step 2: DFS from every cell
	rows, cols := len(board), len(board[0])
	result := []string{}

	var dfs func(node *TrieNode, r, c int)
	dfs = func(node *TrieNode, r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] == '#' {
			return
		}

		ch := board[r][c]
		idx := ch - 'a'
		next := node.Children[idx]
		if next == nil {
			return
		}

		if next.Word != "" {
			result = append(result, next.Word)
			next.Word = ""
		}

		board[r][c] = '#'
		dfs(next, r+1, c)
		dfs(next, r-1, c)
		dfs(next, r, c+1)
		dfs(next, r, c-1)
		board[r][c] = ch
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(trie.root, r, c)
		}
	}

	return result
}
