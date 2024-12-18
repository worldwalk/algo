package maze

import "fmt"

// backtracking find all paths
func backtracking(maze [][]int, cur Point, end Point, visited [][]bool, path []Point, paths *[][]Point) {
	if cur == end {
		*paths = append(*paths, append([]Point{}, path...))
		return
	}
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右
	for _, d := range directions {
		next := Point{cur.x + d.x, cur.y + d.y}
		if next.x < 0 || next.x >= len(maze) || next.y < 0 || next.y >= len(maze[0]) {
			continue
		}
		if maze[next.x][next.y] == 1 || visited[next.x][next.y] {
			continue
		}
		visited[next.x][next.y] = true
		path = append(path, next)
		backtracking(maze, next, end, visited, path, paths)
		visited[next.x][next.y] = false
		path = path[:len(path)-1]
	}
}

func MazeBacktrackingTest() {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
	}
	start := Point{0, 0}
	end := Point{4, 4}
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}
	visited[start.x][start.y] = true
	path := []Point{start}
	var paths [][]Point
	backtracking(maze, start, end, visited, path, &paths)
	for _, p := range paths {
		for _, q := range p {
			fmt.Printf("(%d, %d) ", q.x, q.y)
		}
		fmt.Println()
	}
}
