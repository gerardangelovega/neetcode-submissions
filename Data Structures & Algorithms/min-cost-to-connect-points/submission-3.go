type Point struct {
	x, y int
}
func distance(a, b *Point) int {
	return abs(a.x - b.x) + abs(a.y - b.y)
}
func abs(x int) int {
	if x < 0 { x = -1 * x }
	return x
}

type Edge struct {
	a, b *Point
	distance int
}
func NewEdge(a, b *Point) *Edge {
	return &Edge{
		a: a,
		b: b,
		distance: distance(a, b),
	}
}

type UnionFind struct {
	parent map[*Point]*Point
	rank map[*Point]int
	components int
}
func NewUnionFind(points []*Point) *UnionFind {
	uf := &UnionFind{
		parent: make(map[*Point]*Point, len(points)),
		rank: make(map[*Point]int, len(points)),
		components: len(points),
	}
	for _, point := range points {
		uf.parent[point] = point
		uf.rank[point] = 0
	}
	return uf
}
func (uf *UnionFind) Find(x *Point) *Point {
	var p *Point
	for p = uf.parent[x]; p != uf.parent[p]; p = uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
	}
	return p
}
func (uf *UnionFind) Union(x, y *Point) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py { return false }

	if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else if uf.rank[py] > uf.rank[px] {
		uf.parent[px] = py
	} else {
		uf.parent[px] = py
		uf.rank[py]++
	}
	uf.components--
	return true
}

type PriorityQueue struct {
	items []*Edge
	less func(*Edge, *Edge) bool
}
func NewPriorityQueue(n int) *PriorityQueue {
	return &PriorityQueue {
		items: make([]*Edge, 1, n+1),
		less: func(a, b *Edge) bool { return a.distance < b.distance },
	}
}
func (pq* PriorityQueue) Push(value *Edge) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq* PriorityQueue) Pop() *Edge {
	value := pq.items[1]
	pq.items[1] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.Down(1)
	return value
}
func (pq* PriorityQueue) Up(i int) {
	for ; i > 1 && pq.less(pq.items[i], pq.items[i >> 1]); i = i >> 1 {
		pq.items[i], pq.items[i >> 1] = pq.items[i >> 1], pq.items[i]
	}
}
func (pq* PriorityQueue) Down(i int) {
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
func (pq* PriorityQueue) Clear() {
	pq.items = pq.items[:1]
}

func minCostConnectPoints(in [][]int) int {
	points := make([]*Point, 0, len(in))
	for _, point := range in { points = append(points, &Point{ point[0], point[1] }) }

	uf := NewUnionFind(points)
	pq := NewPriorityQueue(len(points))

	for _, src := range points {
		for _, dst := range points {
			if src.x == dst.x && src.y == dst.y { continue }
			pq.Push(NewEdge(src, dst))
		}
	}

	mst := []*Edge{}
	for len(mst) != len(in)-1 && len(pq.items) > 1 {
		current := pq.Pop()

		if !uf.Union(current.a, current.b) { continue }
		mst = append(mst, current)
	}

	res := 0
	for _, edge := range mst { res = res + edge.distance }
	return res
}