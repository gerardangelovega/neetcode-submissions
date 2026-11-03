type PriorityQueue struct {
	items [][2]int
	less func([2]int, [2]int) bool
}
func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue {
		items: make([][2]int, 1, capacity + 1),
		less: func(a, b [2]int) bool { return a[1] < b[1] },
	}
}
func (pq *PriorityQueue) Push(value [2]int) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() [2]int {
	value := pq.items[1]
	pq.items[1] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.Down(1)
	return value
}
func (pq *PriorityQueue) Up(i int) {
	for ; i > 1 && pq.less(pq.items[i], pq.items[i >> 1]); i = i >> 1 {
		pq.items[i], pq.items[i >> 1] = pq.items[i >> 1], pq.items[i]
	}
}
func (pq *PriorityQueue) Down(i int) {
	n := len(pq.items)
	for (i << 1) < n {
		l, r, s := (i << 1), (i << 1) | 1, i

		if l < n && pq.less(pq.items[l], pq.items[s]) { s = l }
		if r < n && pq.less(pq.items[r], pq.items[s]) { s = r }
		if s == i { break }

		pq.items[i], pq.items[s] = pq.items[s], pq.items[i]
		i = s
	}
}

func shortestPath(n int, edges [][]int, src int) map[int]int {
	mem := n + n / 2
	adj, res := make(map[int][][2]int, mem), make(map[int]int, mem)

	for i := range n { adj[i], res[i] = make([][2]int, 0, mem), -1 }
	for _, e := range edges { adj[e[0]] = append(adj[e[0]], [2]int{ e[1], e[2] }) }

	pq := NewPriorityQueue(mem)
	pq.Push([2]int{ src, 0 })

	for len(pq.items) > 1 {
		p := pq.Pop()

		if v, _ := res[p[0]]; v != -1 { continue }
		res[p[0]] = p[1]

		for _, n := range adj[p[0]] {
			if v, _ := res[n[0]]; v != -1 { continue }
			pq.Push([2]int{ n[0], p[1] + n[1] }) 
		}
	}

	return res
}