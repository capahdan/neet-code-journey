package reversenodeinkgroup

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want *ListNode
	}{
		{name: "test 1", head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, k: 2, want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ReverseKGroup(test.head, test.k)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ReverseKGroup(%v, %d) = %v, want %v", test.head, test.k, got, test.want)
			}
		})
	}
}
