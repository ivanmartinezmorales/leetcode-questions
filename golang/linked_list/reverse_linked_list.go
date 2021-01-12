package lists

type SinglyLinkedList struct {
	Value int
	Next  *SinglyLinkedList
}

// ReverseLinkedList takes the head of a SinglyLinkedList and reverses it
// in place. Returns the new head.
func ReverseLinkedList(head *SinglyLinkedList) *SinglyLinkedList {
	var prev *SinglyLinkedList
	cur := head
	for cur != nil {
		n := cur.Next
		cur.Next = prev
		prev = cur
		cur = n
	}

	return prev
}
