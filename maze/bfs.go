package maze

import "fmt"

// can find the shortest path
// https://blog.csdn.net/xxxzzzqqq_/article/details/130379211
func bfs(maze [][]int, start, end Point) []Point {
	/*
		在这个示例中，我们在BFS算法中添加了一个parents数组，用于记录每个节点的父节点。
		在搜索结束后，我们可以使用parents数组来构造从终点到起点的路径。
		我们从终点开始，依次沿着父节点往回走，直到到达起点。最后，我们将路径反转，以便从起点到终点的顺序打印出来。
		在这个示例中，我们使用了一个简单的迷宫，其中0表示可以通过的路，1表示墙。你可以根据实际情况修改迷宫的大小和形状，以及起点和终点的位置。
	*/
	queue := []Point{start}
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}
	visited[start.x][start.y] = true
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右
	parents := make([][]Point, len(maze))
	for i := range parents {
		parents[i] = make([]Point, len(maze[i]))
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == end {
			break
		}
		for _, d := range directions {
			next := Point{cur.x + d.x, cur.y + d.y}
			if next.x < 0 || next.x >= len(maze) || next.y < 0 || next.y >= len(maze[0]) {
				continue
			}
			if maze[next.x][next.y] == 1 || visited[next.x][next.y] {
				continue
			}
			visited[next.x][next.y] = true
			parents[next.x][next.y] = cur
			queue = append(queue, next)
		}
	}
	path := []Point{end}
	for path[len(path)-1] != start {
		cur := path[len(path)-1]
		parent := parents[cur.x][cur.y]
		path = append(path, parent)
	}
	reverse(path)
	return path
}

func reverse(path []Point) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
func MazeBFSTest() {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
	}
	start := Point{0, 0}
	end := Point{4, 4}
	path := bfs(maze, start, end)
	if path != nil {
		for _, p := range path {
			fmt.Printf("(%d, %d) ", p.x, p.y)
		}
	} else {
		fmt.Println("No path found")
	}
}
