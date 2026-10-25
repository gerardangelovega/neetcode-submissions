/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var root *ListNode

	for head != nil {
		temp := head		
		head = head.Next
		temp.Next = root
		root = temp
	}

	return root
}
