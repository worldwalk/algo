package maze

import (
	"testing"
)

func TestFindPath(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]int
		start    Point
		end      Point
		expected bool // 是否期望找到路径
	}{
		{
			name: "simple path exists",
			maze: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			start:    Point{0, 0},
			end:      Point{2, 2},
			expected: true,
		},
		{
			name: "no path due to wall",
			maze: [][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			start:    Point{0, 0},
			end:      Point{2, 2},
			expected: false,
		},
		{
			name: "start is end",
			maze: [][]int{
				{0, 1, 0},
				{1, 0, 1},
				{0, 1, 0},
			},
			start:    Point{1, 1},
			end:      Point{1, 1},
			expected: true,
		},
		{
			name: "complex maze with path",
			maze: [][]int{
				{0, 1, 0, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 1, 0},
			},
			start:    Point{0, 0},
			end:      Point{4, 4},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := findPath(tt.maze, tt.start, tt.end)

			if tt.expected && path == nil {
				t.Errorf("Expected to find path but got nil")
			}
			if !tt.expected && path != nil {
				t.Errorf("Expected no path but found: %v", path)
			}

			// 验证路径的有效性
			if path != nil {
				// 检查路径起点和终点
				if path[0] != tt.start {
					t.Errorf("Path starts at %v, expected %v", path[0], tt.start)
				}
				if path[len(path)-1] != tt.end {
					t.Errorf("Path ends at %v, expected %v", path[len(path)-1], tt.end)
				}

				// 检查路径连续性
				for i := 1; i < len(path); i++ {
					prev, curr := path[i-1], path[i]
					dx := abs(prev.x - curr.x)
					dy := abs(prev.y - curr.y)
					if (dx == 1 && dy == 0) || (dx == 0 && dy == 1) {
						continue // 有效移动
					}
					t.Errorf("Invalid move from %v to %v", prev, curr)
				}
			}
		})
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestMazeDFSTest(t *testing.T) {
	// 这个测试主要是验证函数不会panic
	MazeDFSTest()
}
