// Given the head of a singly linked list, reverse the list, and return the reversed list.

package datastructures

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	return reverseListImpl(nil, head)
}

func reverseListImpl(prev, cur *ListNode) *ListNode {
	if cur == nil {
		return prev
	}
	next := cur.Next
	cur.Next = prev
	return reverseListImpl(cur, next)
}
