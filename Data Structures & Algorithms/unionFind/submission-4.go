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
    // loop continues until the parent of current is itself (aka. root)
    for curr != uf.parent[curr] {
        // sets the parent of the current element to its grandparent (compression)
        uf.parent[curr] = uf.parent[uf.parent[curr]]
        // sets the next current element to the parent of the current element
        curr = uf.parent[curr]
    }
    // return the root
    return curr
}

func (uf *UnionFind) IsSameComponent(x, y int) bool {
    // find the root components of x and y
    p1, p2 := uf.Find(x), uf.Find(y)
    // return true if the root components of x and y are the same,
    // return false if otherwise
    return p1 == p2
}

func (uf *UnionFind) Union(x, y int) bool {
    // find the root components of x and y
    p1, p2 := uf.Find(x), uf.Find(y)
    // return false if the root components of x and y are the same
    if p1 == p2 { 
        return false 
    }

    if uf.rank[p1] > uf.rank[p2] {
        // if the rank (or height) of the root component of x is greater than y's,
        // then set the parent of the root component of y to be the root component of x
        uf.parent[p2] = p1
    } else if uf.rank[p1] < uf.rank[p2] {
        // if the rank (or height) of the root component of x is less than y's,
        // then set the parent of the root component of x to be the root component of y
        uf.parent[p1] = p2
    } else {
        // if the rank (or height) of the root component of x and y are equal,
        // arbitrarily set the parent of the root component of y to the root component of x
        // and increment the rank (or height) of the root component of y
        uf.parent[p1] = p2
        uf.rank[p2]++
    }
    uf.count--
    return true
}

func (uf *UnionFind) GetNumComponents() int {
    return uf.count
}
