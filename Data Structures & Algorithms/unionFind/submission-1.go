type UnionFind struct {
    parent     map[int]int
    rank       map[int]int
    components int
}

func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{
        parent: make(map[int]int),
        rank: make(map[int]int),
        components: n,
    }

    for i := range n {
        uf.parent[i] = i
        uf.rank[i] = 0
    }

    return uf
}

func (uf *UnionFind) Find(x int) int {
    p := uf.parent[x]
    for p != uf.parent[p] {
        uf.parent[p] = uf.parent[uf.parent[p]]
        p = uf.parent[p]
    }
    return p
}

func (uf *UnionFind) IsSameComponent(x, y int) bool {
    p1, p2 := uf.Find(x), uf.Find(y)
    return p1 == p2
}

func (uf *UnionFind) Union(x, y int) bool {
    x, y = uf.Find(x), uf.Find(y)
    if x == y {
        return false
    }

    if uf.rank[x] > uf.rank[y] {
        uf.parent[y] = x
    } else if uf.rank[x] < uf.rank[y] {
        uf.parent[x] = y
    } else {
        uf.parent[x] = y
        uf.rank[y]++
    }
    uf.components--
    return true
}

func (uf *UnionFind) GetNumComponents() int {
    return uf.components
}
