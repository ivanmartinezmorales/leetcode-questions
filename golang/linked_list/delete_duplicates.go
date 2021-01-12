package lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteDuplicates removes duplicated members from a sorted linked list
// and returns the sorted list in order.
func DeleteDuplicates(head *ListNode) *ListNode {
	s := &ListNode{Next: head}
	p := s

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head.Next = head.Next.Next
			}
			p.Next = head.Next
		} else {
			p = head
		}
		head = head.Next
	}
	return s.Next
}
