package reversenodeinkgroup

// ListNode is a node in the linked list
// Val is the value of the node
// Next is the next node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseKGroup reverses the linked list in k-groups
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	if count < k {
		return head
	}

	prev := &ListNode{}
	curr = head
	next := &ListNode{}

	for i := 0; i < k; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head.Next = ReverseKGroup(next, k)

	return prev
}
