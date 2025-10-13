package maze

import "fmt"

// backtracking find all paths 回溯算法实现
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
		/*
			这里visited的设置放在for内，也是正确的，因为是回溯算法，每次都是针对一个树枝（两个节点）！
			和backtracking2不同！backtracking2是针对一个节点（在访问节点的位置写的append）！
		*/
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

// backtracking2 find all paths 此实现是DFS算法！
func backtracking2(maze [][]int, cur Point, end Point, visited [][]bool, path []Point, paths *[][]Point) {
	if cur.x < 0 || cur.x >= len(maze) || cur.y < 0 || cur.y >= len(maze[0]) ||
		maze[cur.x][cur.y] == 1 || visited[cur.x][cur.y] {
		return
	}
	path = append(path, cur)
	if cur == end {
		*paths = append(*paths, append([]Point{}, path...))
		return
	}
	visited[cur.x][cur.y] = true
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右
	for _, d := range directions {
		next := Point{cur.x + d.x, cur.y + d.y}
		backtracking2(maze, next, end, visited, path, paths)
	}
	visited[cur.x][cur.y] = false
}
func MazeBacktracking2Test() {
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
	path := []Point{}
	var paths [][]Point
	backtracking2(maze, start, end, visited, path, &paths)
	for _, p := range paths {
		for _, q := range p {
			fmt.Printf("(%d, %d) ", q.x, q.y)
		}
		fmt.Println()
	}
}
