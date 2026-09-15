package clonegraph

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {

	testCase := []struct {
		name   string
		node   *Node
		expect *Node
	}{
		{
			name: "Example 1",
			node: &Node{
				Val: 1,
				Neighbors: []*Node{
					&Node{
						Val: 2,
						Neighbors: []*Node{
							&Node{
								Val: 1,
								Neighbors: []*Node{
									&Node{
										Val:       2,
										Neighbors: []*Node{},
									},
								},
							},
						},
					},
				},
			},
			expect: &Node{
				Val: 1,
				Neighbors: []*Node{
					&Node{
						Val: 2,
						Neighbors: []*Node{
							&Node{
								Val: 1,
								Neighbors: []*Node{
									&Node{
										Val:       2,
										Neighbors: []*Node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			if got := CloneGraph(tc.node); !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("CloneGraph() = %v, want %v", got, tc.expect)
			}
		})
	}
}
