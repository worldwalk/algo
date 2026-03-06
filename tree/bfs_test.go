package tree

import (
	"fmt"
	"testing"
)

// TestBFS 测试 BFS 遍历
func TestBFS(t *testing.T) {
	// 构建测试树：
	//         1
	//        / \
	//       2   3
	//      / \   \
	//     4   5   6
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 6,
			},
		},
	}

	fmt.Println("BFS 遍历（打印方式）：")
	bfs(root)

	fmt.Println("BFS 遍历（返回数组）：")
	result := bfsWithResult(root)
	fmt.Println(result)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !equal(result, expected) {
		t.Errorf("bfsWithResult() = %v, want %v", result, expected)
	}

	fmt.Println("BFS 按层遍历：")
	levelOrder := bfsLevelOrder(root)
	for i, level := range levelOrder {
		fmt.Printf("第 %d 层: %v\n", i+1, level)
	}
	expectedLevelOrder := [][]int{{1}, {2, 3}, {4, 5, 6}}
	if !equal2D(levelOrder, expectedLevelOrder) {
		t.Errorf("bfsLevelOrder() = %v, want %v", levelOrder, expectedLevelOrder)
	}
}

// equal 比较两个整数切片是否相等
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// equal2D 比较两个二维整数切片是否相等
func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equal(a[i], b[i]) {
			return false
		}
	}
	return true
}
