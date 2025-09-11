package maze

import (
	"fmt"
)

type Point struct {
	x int // x -> row   better
	y int // y -> col
}

func findPath(maze [][]int, start Point, end Point) []Point {
	visited := make(map[Point]bool)
	path := make([]Point, 0)
	if dfs(maze, start, end, visited, &path) {
		return path
	}
	return nil
}

func dfs(maze [][]int, cur Point, end Point, visited map[Point]bool, path *[]Point) bool {
	if cur == end {
		*path = append(*path, cur)
		return true
	}
	if maze[cur.x][cur.y] == 1 || visited[cur] {
		return false
	}
	visited[cur] = true
	*path = append(*path, cur)

	// 定义四个方向：上、右、下、左
	directions := []Point{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}

	for _, dir := range directions {
		next := Point{cur.x + dir.x, cur.y + dir.y}
		// 检查边界条件
		if next.x >= 0 && next.x < len(maze) && next.y >= 0 && next.y < len(maze[0]) {
			if dfs(maze, next, end, visited, path) {
				return true
			}
		}
	}

	*path = (*path)[:len(*path)-1]
	visited[cur] = false
	return false
}

func MazeDFSTest() {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
	}
	start := Point{0, 0}
	end := Point{4, 4}
	path := findPath(maze, start, end)
	if path != nil {
		for _, p := range path {
			fmt.Printf("(%d, %d) ", p.x, p.y)
		}
	} else {
		fmt.Println("No path found")
	}
}
