/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	prev := head.Next
	curr := head
	curr.Next = nil

	for prev != nil {
		temp := prev
		prev = prev.Next
		temp.Next = curr
		curr = temp
	}

	return curr
}
