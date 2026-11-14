type UnionFind struct {
	parent, rank map[int]int
	components int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make(map[int]int, n),
		rank: make(map[int]int, n),
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

func (uf *UnionFind) IsSameComponent(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
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

func (uf *UnionFind) GetNumComponents() int {
	return uf.components
}
