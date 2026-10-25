type Heap struct {
	items []*ListNode
	cmp func(*ListNode, *ListNode) bool
}
func NewHeap(capacity int) Heap {
	return Heap{
		items: make([]*ListNode, 1, capacity + 1),
		cmp: func(a, b *ListNode) bool {
			return a.Val < b.Val
		},
	}
}
func (h *Heap) Push(node *ListNode) {
	if len(h.items) == 1 {
		h.items = append(h.items, node)
		return
	}
	h.items = append(h.items, node)
	h.Up(len(h.items)-1)
}
func (h *Heap) Pop() *ListNode {
	if len(h.items) == 1 {
		return nil
	}
	if len(h.items) == 2 {
		node := h.items[1]
		h.items = h.items[:1]
		return node
	}
	node := h.items[1]
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.Down(1)
	return node
}
func (h *Heap) Up(i int) {
	for i > 1 && h.cmp(h.items[i], h.items[i >> 1]) {
		h.items[i], h.items[i >> 1] = h.items[i >> 1], h.items[i]
		i = i >> 1
	}
}
func (h *Heap) Down(i int) {
	n := len(h.items)
	for (i << 1) < n {
		l, r, s := (i << 1), (i << 1) | 1, i
		if l < n && h.cmp(h.items[l], h.items[s]) { s = l }
		if r < n && h.cmp(h.items[r], h.items[s]) { s = r }
		if s == i { break }

		h.items[i], h.items[s] = h.items[s], h.items[i]
		i = s
	}
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	heap := NewHeap(4)
	for _, list := range lists {
		for list != nil {
			heap.Push(list)
			list = list.Next
		}
	}

	sentinel := &ListNode{}
	current := sentinel

	for node := heap.Pop(); node != nil; node = heap.Pop() {
		current.Next = node
		current = current.Next
	}

	return sentinel.Next
}