// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer.

package datastructures

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleBruteForce(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	for head != nil && !seen[head] {
		seen[head] = true
		head = head.Next
	}
	return head != nil
}

// Floyd's cycle detecting algorithm or tortoise-hare algorithm
// but it's still has speed compexity O(n) though better space O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	tortoise, hare := head, head.Next
	for hare != nil && hare.Next != nil && tortoise != hare {
		hare = hare.Next.Next
		tortoise = tortoise.Next
	}
	return tortoise == hare
}
