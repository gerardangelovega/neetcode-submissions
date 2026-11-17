// [4]int{ u, v, w, i }
// u - src,
// v - dst,
// w - weight,
// i - index

// UnionFind by vertex
type UnionFind struct {
	parent, rank map[int]int
	components, n int
}
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make(map[int]int, n),
		rank: make(map[int]int, n),
		components: n,
		n: n,
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
	} else if uf.rank[py] > uf.rank[px] {
		uf.parent[px] = py
	} else {
		uf.parent[px] = py
		uf.rank[py]++
	}
	uf.components--
	return true
}
func (uf *UnionFind) Clear() {
	for i := range uf.n { uf.parent[i], uf.rank[i] = i, 0 }
	uf.components = uf.n
}

// PriorityQueue by weight
type PriorityQueue struct {
	items [][4]int
	less func([4]int, [4]int) bool
}
func NewPriorityQueue(n int) *PriorityQueue {
	return &PriorityQueue{
		items: make([][4]int, 1, n+1),
		less: func(a, b [4]int) bool { return a[2] < b[2] },
	}
}
func (pq *PriorityQueue) Push(value [4]int) {
	pq.items = append(pq.items, value)
	pq.Up(len(pq.items)-1)
}
func (pq *PriorityQueue) Pop() [4]int {
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
func (pq *PriorityQueue) Clear() {
	pq.items = pq.items[1:]
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	pq := NewPriorityQueue(n)
	for i, e := range edges { pq.Push([4]int{ e[0], e[1], e[2], i }) }

	sorted := make([][4]int, 0, n)
	for len(pq.items) > 1 { sorted = append(sorted, pq.Pop()) }

	uf := NewUnionFind(n)
	mst := [][4]int{}
	for _, e := range sorted {
		if !uf.Union(e[0], e[1]) { continue }
		mst = append(mst, e)
	}
	weight := 0
	for _, e := range mst { weight = weight + e[2] }

	res := [][]int{ []int{}, []int{} }
	for _, current := range sorted {
		// critical check
		crw := 0
		uf.Clear()
		for _, ed := range sorted {
			if ed[3] == current[3] || !uf.Union(ed[0], ed[1]) { continue }
			crw = crw + ed[2]
		}
		if uf.components != 1 || crw > weight { 
			res[0] = append(res[0], current[3] ) 
			continue
		}
	
		// pseudo-critical check
		pcrw := 0
		uf.Clear()
		uf.Union(current[0], current[1])
		pcrw = current[2]
		for _, ed := range sorted {
			if ed[3] == current[3] || !uf.Union(ed[0], ed[1]) { continue }
			pcrw = pcrw + ed[2]
		}
		if pcrw == weight { res[1] = append(res[1], current[3]) }
	}

	return res
}