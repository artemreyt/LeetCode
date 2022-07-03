// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

package datastructures

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, current *ListNode

	for list1 != nil || list2 != nil {
		var candidate *ListNode
		if list1 != nil && list2 != nil && list1.Val < list2.Val || list2 == nil {
			candidate = list1
			list1 = list1.Next
		} else {
			candidate = list2
			list2 = list2.Next
		}

		if head == nil {
			head = candidate
			current = head
		} else {
			current.Next = candidate
			current = current.Next
		}

	}
	return head
}
