package serializedeserializebinarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec { return Codec{} }

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var sb strings.Builder
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			sb.WriteString("#,")
		} else {
			sb.WriteString(strconv.Itoa(node.Val) + ",")
			queue = append(queue, node.Left, node.Right)
		}
	}
	return sb.String()
}

func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	parts := strings.Split(strings.TrimRight(data, ","), ",")
	root := &TreeNode{Val: mustInt(parts[0])}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(parts) {
		curr := queue[0]
		queue = queue[1:]

		if parts[i] != "#" {
			curr.Left = &TreeNode{Val: mustInt(parts[i])}
			queue = append(queue, curr.Left)
		}
		i++
		if i < len(parts) && parts[i] != "#" {
			curr.Right = &TreeNode{Val: mustInt(parts[i])}
			queue = append(queue, curr.Right)
		}
		i++
	}
	return root
}

func mustInt(s string) int { v, _ := strconv.Atoi(s); return v }
