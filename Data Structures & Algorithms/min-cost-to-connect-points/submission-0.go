type Point struct {
	x, y int
}
func (p *Point) distance(dst *Point) int {
	return abs(dst.x - p.x) + abs(dst.y - p.y)
}
func abs(x int) int {
	if x < 0 { x = x * -1}
	return x
}

type Path struct {
	src, dst Point
	distance int
}

type PriorityQueue struct {
	items []Path
	less func(Path, Path) bool
}
func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		items: make([]Path, 1, capacity + 1),
		less: func(a, b Path) bool { return a.distance < b.distance },
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

func minCostConnectPoints(points [][]int) int {
	n := len(points)
    visited := make(map[Point]Path, n)
    adjacency := make(map[Point][]Path, n)
	for _, point := range points {
		p := Point{ point[0], point[1] }
		adjacency[p] = make([]Path, 0)
	}
	for ksrc, _ := range adjacency {
		for kdst, _ := range adjacency {
			if ksrc == kdst { continue }
			adjacency[ksrc] = append(adjacency[ksrc], Path{ ksrc, kdst, ksrc.distance(&kdst) })
		}
	}

	start := Point{ points[0][0], points[0][1] }
	pq := NewPriorityQueue(n)
	pq.Push(Path { start, start, 0 })

	for len(pq.items) > 1 && len(visited) != n {
		current := pq.Pop()

		if _, e := visited[current.dst]; e { continue }
		visited[current.dst] = current

		for _, adjacent := range adjacency[current.dst] {
			if _, e := visited[adjacent.dst]; e { continue }
			pq.Push(adjacent)
		}
	}

	res := 0
	for _, v := range visited {
		res = res + v.distance
	}
	return res
}
