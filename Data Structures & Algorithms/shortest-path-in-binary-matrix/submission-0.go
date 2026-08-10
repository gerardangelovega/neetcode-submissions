type Queue struct {
	items []Coords
}
func NewQueue(capacity int) Queue {
	return Queue{
		items: make([]Coords, 0, capacity),
	}
}
func (q *Queue) Enqueue(val Coords) {
	q.items = append(q.items, val)
}
func (q *Queue) Dequeue() Coords {
	res := q.items[0]
	q.items = q.items[1:]
	return res
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
func (cd *Coords) Neighbors() []Coords {
	return []Coords{
		Coords{ r: cd.r + 1, c: cd.c },
		Coords{ r: cd.r - 1, c: cd.c },
		Coords{ r: cd.r, c: cd.c + 1 },
		Coords{ r: cd.r, c: cd.c - 1 },
		Coords{ r: cd.r + 1, c: cd.c + 1 },
		Coords{ r: cd.r + 1, c: cd.c - 1 },
		Coords{ r: cd.r - 1, c: cd.c + 1 },
		Coords{ r: cd.r - 1, c: cd.c - 1 },
	}
}

func shortestPathBinaryMatrix(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	if grid[0][0] == 1 {
		return -1
	}
	if grid[rows-1][cols-1] == 1 {
		return -1
	}

	visited := make(map[Coords]struct{})
	queue := NewQueue(6)

	visited[NewCoords(0, 0)] = struct{}{}
	queue.Enqueue(NewCoords(0, 0))

	length := 1

	for queue.Length() != 0 {
		for range queue.Length() {
			coords := queue.Dequeue()

			if coords.r == (rows - 1) && coords.c == (cols - 1) {
				return length
			}

			for _, n := range coords.Neighbors() {
				if min(n.r, n.c) < 0 {
					continue
				}
				if n.r == rows || n.c == cols {
					continue
				}
				if grid[n.r][n.c] == 1 {
					continue
				}
				if _, exists := visited[n]; exists {
					continue
				}
				visited[n] = struct{}{}
				queue.Enqueue(n)
			}
		}
		length = length + 1
	}

	return -1
}
