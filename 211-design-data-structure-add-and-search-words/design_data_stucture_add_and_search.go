package designdatastructureaddandsearchwords

type WordDictionary struct {
	Root *TrieNode
}

type TrieNode struct {
	Children [26]*TrieNode
	IsEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &TrieNode{},
	}
}

func (this *WordDictionary) Insert(word string) {

	node := this.Root
	for _, ch := range word {
		idx := ch - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &TrieNode{}
		}
		node = node.Children[idx]
	}
	node.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.Root, word, 0)
}
func dfs(node *TrieNode, word string, i int) bool {

	// CASE 1 — we ran out of letters
	if i == len(word) {
		return node.IsEnd // only true if a real word ends here
	}

	ch := rune(word[i])

	// CASE 2 — the current letter is a dot
	if ch == '.' {
		for _, child := range node.Children {
			if child != nil && dfs(child, word, i+1) {
				return true // ONE branch succeeded — that's enough
			}
		}
		return false // no branch matched
	}

	// CASE 3 — normal letter
	idx := ch - 'a'
	if node.Children[idx] == nil {
		return false // dead end
	}
	return dfs(node.Children[idx], word, i+1)
}
