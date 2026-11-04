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

func networkDelayTime(times [][]int, n int, k int) int {
	mem := n + n / 2
    adj, shortest := make(map[int][][2]int, mem), make(map[int]int)

	for i := range n { adj[i + 1], shortest[i + 1] = make([][2]int, 0, n), -1 }
	for _, time := range times { adj[time[0]] = append(adj[time[0]], [2]int{ time[1], time[2] }) }

	pq := NewPriorityQueue(mem)
	pq.Push([2]int{ k, 0 })

	for len(pq.items) > 1 {
		current := pq.Pop()

		if v, _ := shortest[current[0]]; v != -1 { continue }
		shortest[current[0]] = current[1]

		for _, adjacent := range adj[current[0]] {
			if v, _ := shortest[adjacent[0]]; v != -1 { continue }
			pq.Push([2]int{ adjacent[0], current[1] + adjacent[1] })
		}
	}

	res := math.MinInt
	for _, v := range shortest {
		if v == -1 { return -1 }
		res = max(res, v)
	}

	return res
}
