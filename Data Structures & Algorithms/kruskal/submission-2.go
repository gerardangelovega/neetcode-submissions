type UnionFind struct {
	parent, rank map[int]int
	components int
}
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make(map[int]int),
		rank: make(map[int]int),
		components: n,
	}
	for i := range n { uf.parent[i], uf.rank[i] = i, 0 }
	return uf
}
func (uf *UnionFind) Find(x int) int {
	var p int
	for p = uf.parent[x]; p != uf.parent[p]; p = uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
	}
	return p
}
func (uf *UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py { return false }

	if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else {
		uf.parent[px] = py
		uf.rank[py]++
	}
	uf.components--
	return true
}

type PriorityQueue struct {
	items [][3]int
	less func([3]int, [3]int) bool
}
func NewPriorityQueue(n int) *PriorityQueue {
	return &PriorityQueue{
		items: make([][3]int, 1, n+1),
		less: func(a, b [3]int) bool { return a[2] < b[2] },
	}
}
func (pq* PriorityQueue) Push(value [3]int) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq* PriorityQueue) Pop() [3]int {
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

func minimumSpanningTree(n int, edges [][]int) int {
	uf := NewUnionFind(n)
	pq := NewPriorityQueue(n)
	for _, edge := range edges { pq.Push([3]int{ edge[0], edge[1], edge[2] }) }

	mst := [][3]int{}
	for len(mst) != n-1 && len(pq.items) > 1 {
		current := pq.Pop()

		if !uf.Union(current[0], current[1]) { continue }
		mst = append(mst, current)
	}

	res := 0
	for _, edge := range mst { res = res + edge[2] }
	if uf.components == 1 { return res } else { return -1 }
}
