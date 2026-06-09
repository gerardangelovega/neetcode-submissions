/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]bool)

	for curr := head; curr != nil; curr = curr.Next {
		if m[curr] {
			return true
		}

		m[curr] = true
	}

	return false
}
