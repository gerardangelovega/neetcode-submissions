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
func (uf* UnionFind) Find(x int) int {
	var p int
	for p = uf.parent[x]; p != uf.parent[p]; p = uf.parent[p] { uf.parent[p] = uf.parent[uf.parent[p]]}
	return p
}
func (uf* UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py { return false }

	if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else if uf.rank[px] < uf.rank[py] { 
		uf.parent[px] = py
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
	uf.components--
	return true
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	uf := NewUnionFind(n)
	for i := range n {
		if left := leftChild[i]; left != -1 && !uf.Union(i, left) { return false }
		if right := rightChild[i]; right != -1 && !uf.Union(i, right) { return false }
	}
	return uf.components == 1
}
