type PriorityQueue struct {
	items [][3]int
	less func([3]int, [3]int) bool
}
func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		items: make([][3]int, 1, capacity + 1),
		less: func(a, b [3]int) bool { return a[2] < b[2] },
	}
}
func (pq *PriorityQueue) Push(value [3]int) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() [3]int {
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

func minimumSpanningTree(n int, edges [][]int) int {
	visited := make(map[int][3]int)
	adjacencies := make(map[int][][3]int, n)
	for i, _ := range adjacencies {
		adjacencies[i] = make([][3]int, 0)
	}
	for _, edge := range edges {
		adjacencies[edge[0]] = append(adjacencies[edge[0]], [3]int{ edge[0], edge[1], edge[2]})
		adjacencies[edge[1]] = append(adjacencies[edge[1]], [3]int{ edge[1], edge[0], edge[2]})
	}

	pq := NewPriorityQueue(n)
	pq.Push([3]int{ 0, 0, 0 })

	for k, v := range adjacencies {
		fmt.Println(k, v)
	}

	for len(pq.items) > 1 && len(visited) != n {
		current := pq.Pop()
		fmt.Println("loop", current)

		if _, e := visited[current[1]]; e { continue }
		visited[current[1]] = current

		for _, adjacent := range adjacencies[current[1]] {
			fmt.Println("\tloop", adjacent)
			if _, e := visited[adjacent[1]]; e { continue }
			pq.Push(adjacent)
			fmt.Println("\tpushed")
		}
	}

	if len(visited) != n { return -1 }
	res := 0
	for k, v := range visited {
		fmt.Println(k, v)
		res = res + v[2]
	}
	return res
}
