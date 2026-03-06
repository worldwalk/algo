package tree

import "fmt"

// bfs 二叉树的广度优先搜索（BFS）遍历，也称为层序遍历
//
// 算法说明：
// BFS 使用队列（FIFO）来实现，按照从上到下、从左到右的顺序访问节点。
// 1. 将根节点入队
// 2. 当队列不为空时：
//   - 取出队首节点并访问
//   - 将该节点的左子节点（如果存在）入队
//   - 将该节点的右子节点（如果存在）入队
//
// 3. 重复步骤2直到队列为空
//
// 时间复杂度：O(n)，其中 n 是树中节点的数量
// 空间复杂度：O(w)，其中 w 是树的最大宽度（最宽层的节点数）
//
// 示例：
//
//	对于树：
//	      1
//	     / \
//	    2   3
//	   / \   \
//	  4   5   6
//	BFS 遍历结果：1 -> 2 -> 3 -> 4 -> 5 -> 6
func bfs(root *Node) {
	if root == nil {
		return
	}

	// 使用 slice 作为队列
	queue := []*Node{root}

	for len(queue) > 0 {
		// 取出队首节点
		node := queue[0]
		queue = queue[1:]

		// 访问当前节点
		fmt.Printf("%d ", node.Val)

		// 将左右子节点入队（如果存在）
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

// bfsWithResult 二叉树的 BFS 遍历，返回遍历结果数组
//
// 返回值：按 BFS 顺序的节点值数组
func bfsWithResult(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// bfsLevelOrder 按层返回 BFS 遍历结果（每层一个数组）
//
// 返回值：二维数组，每个子数组代表一层的节点值
//
// 示例：
//
//	对于树：
//	      1
//	     / \
//	    2   3
//	   / \   \
//	  4   5   6
//	返回：[[1], [2, 3], [4, 5, 6]]
func bfsLevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数
		level := []int{}

		// 处理当前层的所有节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			// 将下一层的节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
