package main

func driver(input [][]byte) int {
	return numIslands(input)
}

func numIslands(grid [][]byte) (islandCount int) {
	stack := Stack{}
	for i, row := range grid {
		for j, val := range row {
			if val == '1' {
				islandCount++
				// explore all land linked which turns
				// explore(i, j, &grid)
				grid[i][j] = '0'
				stack.Push(getNeighbours(Point{i, j}, grid))
				for stack.Size() > 0 {
					var last Point
					last = stack.Pop()
					grid[last.i][last.j] = '0'
					stack.Push(getNeighbours(last, grid))
				}
			}
		}
	}
	return
}

type Point struct{ i, j int }

type Stack struct {
	stack []Point
}

func (s *Stack) Pop() Point {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return last
}

func (s *Stack) Push(ps []Point) {
	s.stack = append(s.stack, ps...)
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func getNeighbours(p Point, grid [][]byte) (land []Point) {
	neighbours := []Point{
		{p.i - 1, p.j},
		{p.i + 1, p.j},
		{p.i, p.j - 1},
		{p.i, p.j + 1},
	}

	for _, neighbour := range neighbours {
		if isInGrid(neighbour, grid) && grid[neighbour.i][neighbour.j] == '1' {
			land = append(land, neighbour)
		}
	}
	return
}

func isInGrid(p Point, grid [][]byte) bool {
	return p.i < len(grid) &&
		p.i >= 0 &&
		p.j < len(grid[0]) &&
		p.j >= 0
}
