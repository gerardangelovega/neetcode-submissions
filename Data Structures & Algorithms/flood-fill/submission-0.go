type Coords struct {
	r int
	c int
}
func NewCoords(r, c int) Coords {
	return Coords{ r: r, c: c }
}
func (cd* Coords) Neighbors() []Coords {
	return []Coords {
		Coords{ r: cd.r + 1, c: cd.c },
		Coords{ r: cd.r - 1, c: cd.c },
		Coords{ r: cd.r, c: cd.c + 1 },
		Coords{ r: cd.r, c: cd.c - 1 },
	}
}

type Queue struct {
	items []Coords
}
func NewQueue(capacity int) Queue {
	return Queue{ items: make([]Coords, 0, capacity) }
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

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	fill(image, sr, sc, color)
	return image
}

func fill(image [][]int, sr, sc, color int) {
	rows, cols := len(image), len(image[0])
	visited := make(map[Coords]struct{})
	queue := NewQueue(4)

	visited[NewCoords(sr, sc)] = struct{}{}
	queue.Enqueue(NewCoords(sr, sc))

	base := image[sr][sc]

	for queue.Length() != 0 {
		for range queue.Length() {
			coords := queue.Dequeue()
			image[coords.r][coords.c] = color

			for _, neighbor := range coords.Neighbors() {
				if min(neighbor.r, neighbor.c) < 0 {
					continue
				}
				if neighbor.r == rows || neighbor.c == cols {
					continue
				}
				if image[neighbor.r][neighbor.c] != base {
					continue
				}
				if _, exists := visited[neighbor]; exists {
					continue
				}
				queue.Enqueue(neighbor)
				visited[neighbor] = struct{}{}
			}
		}
	}
}