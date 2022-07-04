// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

package datastructures

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
