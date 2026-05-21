package serializedeserializebinarytree

import (
	"reflect"
	"testing"
)

func TestSerializeDeserializeBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{name: "Test Serialize", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: "1,2,3,#,#,#,#,"},
		{name: "Test Deserialize", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: "1,2,3,#,#,#,#,"},
	}

	codec := Constructor()

	for _, test := range tests {

		if test.name == "Test Serialize" {
			t.Run(test.name, func(t *testing.T) {
				got := codec.serialize(test.root)
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("Serialize(%s) = %v, want %v", test.want, got, test.want)
				}
			})
		}

		if test.name == "Test Deserialize" {
			t.Run(test.name, func(t *testing.T) {
				got := codec.deserialize(test.want)
				if !reflect.DeepEqual(got, test.root) {
					t.Errorf("Deserialize(%s) = %v, want %v", test.want, got, test.root)
				}
			})
		}
	}
}
