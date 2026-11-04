type Data struct {
	index int	
	probability float64
}

type PriorityQueue struct {
	items []Data
	less func(Data, Data) bool
}
func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		items: make([]Data, 1, capacity + 1),
		less: func(a, b Data) bool { return a.probability > b.probability },
	}
}
func (pq *PriorityQueue) Push(value Data) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() Data {
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
	adj := make([][]Data, n)
	probabilities := make([]float64, n)

	for i := range n {
		adj[i] = make([]Data, 0)
		probabilities[i] = -1.0
	}
	for i, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], Data{ edge[1], succProb[i]})
		adj[edge[1]] = append(adj[edge[1]], Data{ edge[0], succProb[i]})
	}

	pq := NewPriorityQueue(n)
	pq.Push(Data{ start_node, 1.0 })

	for len(pq.items) > 1 {
		current := pq.Pop()
		if current.index == end_node { return current.probability }

		if probabilities[current.index] != -1.0 { continue }
		probabilities[current.index] = current.probability

		for _, neighbor := range adj[current.index] {
			if probabilities[neighbor.index] != -1.0 { continue }
			pq.Push(Data{ neighbor.index, current.probability * neighbor.probability })
		}
	}

	return 0.0
}