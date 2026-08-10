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

func countPaths(grid [][]int) int {
	return dfs(grid, Coords{ r: 0, c: 0 }, make(map[Coords]struct{}))
}

func dfs(grid [][]int, coords Coords, visited map[Coords]struct{}) int {
	rows, columns := len(grid), len(grid[0])
	if min(coords.r, coords.c) < 0 {
		return 0
	}
	if coords.r == rows || coords.c == columns {
		return 0
	}
	if grid[coords.r][coords.c] == 1 {
		return 0
	}
	if _, exists := visited[coords]; exists {
		return 0
	}

	if coords.r == (rows - 1) && coords.c == (columns - 1) {
		return 1
	}

	visited[coords] = struct{}{}

	count := 0
	for _, neighbor := range coords.Neighbors() {
		count = count + dfs(grid, neighbor, visited)
	}

	delete(visited, coords)
	return count
}
