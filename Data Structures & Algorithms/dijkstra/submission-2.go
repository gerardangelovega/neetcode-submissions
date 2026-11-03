type Edge [3]int
type Path [2]int

type PriorityQueue struct {
	items []Path
	less func(Path, Path) bool
}
func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{
		items: make([]Path, 1),
		less: func(a, b Path) bool {
			return a[1] < b[1]
		},
	}
}
func (pq *PriorityQueue) Push(path Path) {
	if len(pq.items) < 2 {
		pq.items = append(pq.items, path)
		return
	}
	pq.items = append(pq.items, path)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() Path {
	if len(pq.items) < 2 {
		return Path{-1, -1}
	}
	if len(pq.items) < 3 {
		path := pq.items[1]
		pq.items = pq.items[:1]
		return path
	}
	path := pq.items[1]
	pq.items[1] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.Down(1)
	return path
}
func (pq *PriorityQueue) Top() Path {
	if len(pq.items) < 2 {
		return Path{-1, -1}
	}
	return pq.items[1]
}
func (pq *PriorityQueue) Up(i int) {
	for i > 1 && pq.less(pq.items[i], pq.items[i >> 1]) {
		pq.items[i], pq.items[i >> 1] = pq.items[i >> 1], pq.items[i]
		i = i >> 1
	}
}
func (pq *PriorityQueue) Down(i int) {
	n := len(pq.items)
	for (i << 1) < n {
		l := (i << 1)
		r := (i << 1) | 1
		s :=  i

		if l < n && pq.less(pq.items[l], pq.items[s]) { s = l }
		if r < n && pq.less(pq.items[r], pq.items[s]) { s = r }
		if s == i { break }

		pq.items[i], pq.items[s] = pq.items[s], pq.items[i]
		i = s
	}
}

func shortestPath(n int, edges [][]int, src int) map[int]int {
	adj := make(map[int][]Edge)
	for i := range n {
		adj[i] = make([]Edge, 0)
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], Edge{ edge[0], edge[1], edge[2] })
	}

	shortest := make(map[int]int)
	for i := range n {
		shortest[i] = -1
	}

	pq := NewPriorityQueue()
	pq.Push(Path{src, 0})

	for len(pq.items) > 1 {
		path := pq.Pop()
		if v, e := shortest[path[0]]; e && v != -1 {
			continue
		}
		shortest[path[0]] = path[1]

		for _, neighbor := range adj[path[0]] {
			if v, e := shortest[neighbor[1]]; e && v != -1 {
				continue
			}
			pq.Push(Path{ neighbor[1], neighbor[2] + path[1] })
		}
	}

	return shortest
}
