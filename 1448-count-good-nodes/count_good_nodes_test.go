package countgoodnodes

import "testing"

func TestCountGoodNodes(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val: 4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 5},
				},
			},
			want: 4,
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{Val: 2},
					},
				},
			},
			want: 3,
		},
		{
			name: "example 3",
			root: &TreeNode{Val: 1},
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GoodNodes(test.root)
			if got != test.want {
				t.Errorf("GoodNodes(%v) = %d, want %d", test.root, got, test.want)
			}
		})
	}
}
