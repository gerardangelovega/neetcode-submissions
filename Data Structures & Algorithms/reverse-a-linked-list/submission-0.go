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
	
	list := &ListNode{
		Val: head.Val,
		Next: nil,
	}

   for curr := head.Next; curr != nil; curr = curr.Next {
		node := &ListNode{
			Val: list.Val,
			Next: list.Next,
		}
		list = &ListNode{
			Val: curr.Val,
			Next: node,
		}
   }

   return list
}
