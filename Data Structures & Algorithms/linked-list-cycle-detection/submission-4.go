/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	for s, f := head, head; f != nil && f.Next != nil; {
		s, f = s.Next, f.Next.Next
		if s == f { return true }
	}
	return false
}
