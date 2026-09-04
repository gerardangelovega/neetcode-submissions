type UnionFind struct {
    parent map[int]int
    rank   map[int]int
    count  int
}

func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{
        parent: make(map[int]int),
        rank: make(map[int]int),
        count: n,
    }

    for i := range n {
        uf.parent[i] = i
        uf.rank[i] = 0
    }

    return uf
}

func (uf *UnionFind) Find(x int) int {
    curr := uf.parent[x]
    for curr != uf.parent[curr] {
        uf.parent[curr] = uf.parent[uf.parent[curr]]
        curr = uf.parent[curr]
    }
    return curr
}

func (uf *UnionFind) IsSameComponent(x, y int) bool {
    p1, p2 := uf.Find(x), uf.Find(y)
    return p1 == p2
}

func (uf *UnionFind) Union(x, y int) bool {
    p1, p2 := uf.Find(x), uf.Find(y)
    if p1 == p2 { 
        return false 
    }

    if uf.rank[p1] > uf.rank[p2] {
        uf.parent[p2] = p1
    } else if uf.rank[p1] < uf.rank[p2] {
        uf.parent[p1] = p2
    } else {
        uf.parent[p1] = p2
        uf.rank[p2]++
    }
    uf.count--
    return true
}

func (uf *UnionFind) GetNumComponents() int {
    return uf.count
}
