type Coords struct {
	r int
	c int
}
func (cd *Coords) Neighbors() []Coords {
	return []Coords{
		Coords{ r: cd.r + 1, c: cd.c },
		Coords{ r: cd.r - 1, c: cd.c },
		Coords{ r: cd.r, c: cd.c + 1 },
		Coords{ r: cd.r, c: cd.c - 1 },
	}
}
func NewCoords(r, c int) Coords {
	return Coords{ r: r, c: c }
}

type Queue struct {
	items []Coords
}
func NewQueue() Queue {
	return Queue{ items: []Coords{} }
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

func shortestPath(grid [][]int) int {
	return bfs(grid)
}

func bfs(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	if grid[0][0] == 1 {
		return -1
	}
	if grid[rows-1][cols-1] == 1 {
		return -1
	}

	queue := NewQueue()
	visited := make(map[Coords]struct{})

	queue.Enqueue(NewCoords(0, 0))
	visited[NewCoords(0, 0)] = struct{}{}

	length := 0

	for queue.Length() != 0 {
		for range queue.Length() {
			coords := queue.Dequeue()
			if coords.r == (rows - 1) && coords.c == (cols - 1) {
				return length
			}

			for _, neighbor := range coords.Neighbors() {
				if min(neighbor.r, neighbor.c) < 0 {
					continue
				}
				if neighbor.r == rows || neighbor.c == cols {
					continue
				}
				if _, exists := visited[neighbor]; exists {
					continue
				}
				if grid[neighbor.r][neighbor.c] == 1 {
					continue
				}
				queue.Enqueue(neighbor)
				visited[neighbor] = struct{}{}
			}
		}
		length = length + 1
	}

	return -1
}