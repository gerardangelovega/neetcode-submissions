import (
	"slices"
	"cmp"
)

type UnionFind struct {
	parent, rank []int
	components, n int
}
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank: make([]int, n),
		components: n,
		n: n,
	}
	for i := range n { uf.parent[i] = i }
	return uf
}
func (uf *UnionFind) Find(x int) (p int) {
	for p = uf.parent[x]; p != uf.parent[p]; p = uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
	}
	return
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
	for i := range uf.n { uf.parent[i] = i }
	uf.components = uf.n
}


func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	sorted := make([][4]int, 0, len(edges))
	for i, e := range edges { sorted = append(sorted, [4]int{ e[0], e[1], e[2], i }) }
	slices.SortFunc(sorted, func(a, b [4]int) int { return cmp.Compare( a[2], b[2] )})

	uf := NewUnionFind(n)
	weight := 0
	for _, e := range sorted {
		if !uf.Union(e[0], e[1]) { continue }
		weight = weight + e[2]
	}

	res := [][]int{ []int{}, []int{} }
	for _, current := range sorted {
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
		if crw + current[2] == weight {
			res[1] = append(res[1], current[3] ) 
			continue
		}
	
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