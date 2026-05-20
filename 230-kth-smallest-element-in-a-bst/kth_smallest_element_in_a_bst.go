package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root
	count := 0

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		if count == k {
			return curr.Val
		}
		curr = curr.Right
	}

	return 0
}

func inOder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := inOder(root.Left)
	mid := []int{root.Val}
	right := inOder(root.Right)

	result := append(left, mid...)
	result = append(result, right...)
	return result
}
