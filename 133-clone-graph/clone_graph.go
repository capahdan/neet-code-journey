package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

// the intuition that i want to share that we need to check all the graph that exist in the Node
// we need to do deep first search but the problem with breath frist search so wee

func CloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	return helper(node, visited)
}

func helper(node *Node, visited map[*Node]*Node) *Node {

	if gotNode, ok := visited[node]; ok {
		return gotNode
	}

	clonedNode := &Node{
		Val: node.Val,
	}

	visited[node] = clonedNode

	for _, curr := range node.Neighbors {
		cloneNeighbor := helper(curr, visited)
		clonedNode.Neighbors = append(clonedNode.Neighbors, cloneNeighbor)
	}

	return clonedNode
}
