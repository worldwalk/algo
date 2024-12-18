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
	if cur.x > 0 && dfs(maze, Point{cur.x - 1, cur.y}, end, visited, path) {
		return true
	}
	if cur.x < len(maze)-1 && dfs(maze, Point{cur.x + 1, cur.y}, end, visited, path) {
		return true
	}
	if cur.y > 0 && dfs(maze, Point{cur.x, cur.y - 1}, end, visited, path) {
		return true
	}
	if cur.y < len(maze[0])-1 && dfs(maze, Point{cur.x, cur.y + 1}, end, visited, path) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	visited[cur] = false // 仅找一条路径，这行貌似没用
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
