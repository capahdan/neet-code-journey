package countgoodnodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GoodNodes(root *TreeNode) int {
	return helper(root, root.Val)
}

func helper(root *TreeNode, maxValue int) int {

	if root == nil {
		return 0
	}

	result := 0
	maxValue = max(maxValue, root.Val)
	result += helper(root.Left, maxValue)
	result += helper(root.Right, maxValue)

	if root.Val >= maxValue {
		result += 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
