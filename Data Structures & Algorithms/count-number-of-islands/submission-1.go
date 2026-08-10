type Queue struct {
	items []Coords
}
func NewQueue() Queue {
	return Queue{ items: make([]Coords, 0) }
}
func (q *Queue) Enqueue(val Coords) {
	q.items = append(q.items, val)
}
func (q *Queue) Dequeue() Coords {
	coords := q.items[0]
	q.items = q.items[1:]
	return coords
}
func (q *Queue) Length() int {
	return len(q.items)
}

type Coords struct {
	r int
	c int
}
func NewCoords(r, c int) Coords {
	return Coords{ r: r, c: c }
}
func (cd * Coords) Neighbors() []Coords {
	return []Coords{
		Coords{ r: cd.r + 1, c: cd.c },
		Coords{ r: cd.r - 1, c: cd.c },
		Coords{ r: cd.r, c: cd.c + 1 },
		Coords{ r: cd.r, c: cd.c - 1 },
	}
}

func numIslands(grid [][]byte) int {
	visited := make(map[Coords]struct{})
	queue := NewQueue()

	for r, row := range grid {
		for c, col := range row {
			if col == byte('1') { queue.Enqueue(NewCoords(r, c)) }
		}
	}

	islands := 0

	for queue.Length() != 0 {
		coords := queue.Dequeue()

		if _, exists := visited[coords]; exists {
			continue
		}

		dfs(grid, coords, visited)

		islands = islands + 1
	}

	return islands
}

func dfs(grid [][]byte, c Coords, visited map[Coords]struct{}) {
	rows, cols := len(grid), len(grid[0])

	if min(c.r, c.c) < 0 {
		return
	}
	if c.r == rows || c.c == cols {
		return
	}
	if grid[c.r][c.c] == byte('0') {
		return
	}
	if _, exists := visited[c]; exists {
		return
	}

	visited[c] = struct{}{}

	for _, n := range c.Neighbors() {
		dfs(grid, n, visited)
	}
}
