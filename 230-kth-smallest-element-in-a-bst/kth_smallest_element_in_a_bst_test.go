package kthsmallestelementinabst

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{name: "test 1", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}, k: 1, want: 1},
		// {name: "test 2", root: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}, k: 3, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := KthSmallest(test.root, test.k)
			if got != test.want {
				t.Errorf("KthSmallest(%v, %d) = %d, want %d", test.root, test.k, got, test.want)
			}
		})
	}
}
