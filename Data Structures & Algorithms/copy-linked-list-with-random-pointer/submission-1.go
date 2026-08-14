/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	copied := make(map[*Node]*Node)

	root := &Node{ Val: head.Val }
	copied[head] = root

	curr := head.Next

	for curr != nil {
		root.Next = &Node{
			Val: curr.Val,
		}

		copied[curr] = root.Next

		root = root.Next
		curr = curr.Next
	}

	root = copied[head]
	curr = head

	for curr != nil {
		if n, e := copied[curr.Random]; e {
			root.Random = n
		}
		root = root.Next
		curr = curr.Next
	}


	return copied[head];
}
