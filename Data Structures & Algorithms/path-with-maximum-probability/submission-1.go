type PriorityQueue struct {
	items [][2]float64
	less func([2]float64, [2]float64) bool
}
func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		items: make([][2]float64, 1, capacity + 1),
		less: func(a, b [2]float64) bool { return a[1] > b[1] },
	}
}
func (pq *PriorityQueue) Push(value [2]float64) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() [2]float64 {
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

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	mem := n + n / 2
	adj := make(map[int][][2]float64, mem)
	probabilities := make(map[int]float64, mem)

	for i := range n {
		adj[i] = make([][2]float64, 0, n)
	}
	for i, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], [2]float64{ float64(edge[1]), succProb[i]})
		adj[edge[1]] = append(adj[edge[1]], [2]float64{ float64(edge[0]), succProb[i]})
	}

	for k, v := range adj {
		fmt.Println(k, v)
	}
	
	pq := NewPriorityQueue(mem)
	pq.Push([2]float64{ float64(start_node), 1.0 })

	for len(pq.items) > 1 {
		current := pq.Pop()

		if _, e := probabilities[int(current[0])]; e {
			continue
		}
		probabilities[int(current[0])] = current[1]

		for _, neighbor := range adj[int(current[0])] {
			if _, e := probabilities[int(neighbor[0])]; e {
				continue
			}
			pq.Push([2]float64{ neighbor[0], current[1] * neighbor[1] })
		}
	}

	for k, v := range probabilities {
		fmt.Println(k, v)
	}

	if v, e := probabilities[end_node]; e { return v } else { return 0.0 }
}