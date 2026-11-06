type Path struct {
	coords [2]int
	depth int
}

type PriorityQueue struct {
	items []Path
	less func(Path, Path) bool
}
func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		items: make([]Path, 1, capacity + 1),
		less: func(a, b Path) bool { return a.depth < b.depth },
	}
}
func (pq *PriorityQueue) Push(value Path) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() Path {
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

func swimInWater(grid [][]int) int {
	n, m, nxm := len(grid), len(grid[0]), len(grid) * len(grid[0])
	adj := make(map[[2]int][]Path, nxm)
	shortest := make(map[[2]int]Path, nxm)

	for i := range n {
		for j := range m {
			adj[[2]int{ i, j }] = make([]Path, 0)
		}
	}
	for i, row := range grid {
		for j, _ := range row {
			for _, neighbor := range neighbors(i, j) {
				if min(neighbor[0], neighbor[1]) < 0 { continue }
				if neighbor[0] == n || neighbor[1] == m { continue }
				adjacent := Path{
					[2]int{ neighbor[0], neighbor[1] },
					grid[neighbor[0]][neighbor[1]],
				}
				adj[[2]int{ i, j }] = append(adj[[2]int{ i, j }], adjacent)
			}
		}
	}

	pq := NewPriorityQueue(nxm)
	pq.Push(Path{ [2]int{ 0, 0 }, grid[0][0] })

	for len(pq.items) > 1 {
		curr := pq.Pop()

		if _, e := shortest[curr.coords]; e {
			continue
		}
		shortest[curr.coords] = curr
		if curr.coords[0] == n-1 && curr.coords[1] == m-1 {
			break
		}

		for _, adjacent := range adj[curr.coords] {
			if _, e := shortest[adjacent.coords]; e {
				continue
			}
			pq.Push(Path{adjacent.coords, max(curr.depth, adjacent.depth)})
		}
	}

	for k, v := range shortest{
		fmt.Println(k, v)
	}

	return shortest[[2]int{ n-1, m-1 }].depth
}

func neighbors(i, j int) [][2]int {
	return [][2]int{
		[2]int{ i + 1, j },
		[2]int{ i, j + 1 },
		[2]int{ i - 1, j },
		[2]int{ i, j - 1 },
	}
}