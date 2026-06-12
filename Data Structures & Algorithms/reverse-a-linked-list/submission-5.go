/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(prev *ListNode, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}

	next := curr.Next
	curr.Next = prev

	return reverse(curr, next)
}
