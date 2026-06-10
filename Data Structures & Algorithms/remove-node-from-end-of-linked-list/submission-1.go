/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
		if curr.Next == nil {
			break
		}
	}

	index := length - n
	var prev *ListNode

	for curr, i := head, 0; curr != nil; curr, i = curr.Next, i + 1 {
		if i == index {
			if prev == nil {
				head = curr.Next
			} else if curr.Next == nil {
				prev.Next = nil
			} else {
				prev.Next = curr.Next
			}
			break
		}
		prev = curr
	}

	return head
}
