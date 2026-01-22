package dyoprog

import "fmt"

// LongestIncreasingSubsequence
// https://writings.sh/post/algorithm-longest-increasing-subsequence

// lengthOfLIS 计算最长递增子序列（Longest Increasing Subsequence, LIS）的长度
//
// 题目描述：
// 给定一个整数数组 nums，找到其中最长严格递增子序列的长度。
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
//
// 示例：
//
//	输入: [10,9,2,5,3,7,101,18]
//	输出: 4
//	解释: 最长递增子序列是 [2,3,7,101]，因此长度为 4
//
// 算法分析：
// 使用动态规划方法解决：
// 1. 状态定义：dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
// 2. 状态转移：对于每个位置 i，遍历所有 j < i：
//   - 如果 nums[i] > nums[j]，说明可以将 nums[i] 接在以 nums[j] 结尾的子序列后面
//   - 更新 dp[i] = max(dp[i], dp[j] + 1)
//
// 3. 初始状态：每个位置的最短子序列长度至少为 1（只包含自身）
// 4. 最终答案：遍历所有 dp[i]，取最大值
//
// 时间复杂度：O(n²)，其中 n 是数组长度
//   - 外层循环遍历 n 个元素
//   - 内层循环对每个元素最多遍历 i 次（i 从 1 到 n-1）
//   - 总时间复杂度为 O(1+2+...+(n-1)) = O(n²)
//
// 空间复杂度：O(n)，用于存储 dp 数组
//
// 优化思路：
// 可以使用二分查找优化到 O(n log n) 的时间复杂度，通过维护一个 tails 数组，
// 其中 tails[i] 表示长度为 i+1 的所有递增子序列中，最小的末尾元素值。
func lengthOfLIS(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	dp := make([]int, len(arr))
	dp[0] = 1
	maxL := dp[0]
	for i := 1; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxL = max(maxL, dp[i])
	}

	return maxL
}

// printLISPath 打印所有最长递增子序列的路径
//
// 功能说明：
// 不仅计算最长递增子序列的长度，还找出并打印所有可能的最长递增子序列。
// 例如：对于数组 [10,9,2,5,3,7,101,18]，最长长度为 4，可能的路径有：
//   - [2,3,7,101]
//   - [2,3,7,18]
//   - [2,5,7,101]
//   - [2,5,7,18]
//
// 算法流程：
// 1. 调用 pathOfLIS 计算最长长度并构建路径矩阵（记录每个位置的前驱位置）
// 2. 找到所有以最长长度结尾的位置（dp[i] == maxL）
// 3. 从这些位置开始，通过 DFS 回溯所有可能的路径并打印
func printLISPath(arr []int) {
	dim := len(arr)
	dp := make([]int, dim)       // dp[i] 表示以 arr[i] 结尾的最长递增子序列长度
	matrix := make([][]int, dim) // matrix[i] 存储所有可以转移到位置 i 的前驱位置索引
	for i := 0; i < dim; i++ {
		matrix[i] = make([]int, dim)
	}
	maxL := pathOfLIS(arr, dp, matrix) // 计算最长长度并构建路径矩阵
	fmt.Println("maxL", maxL)

	path := make([]int, dim) // 用于存储当前回溯路径
	// 找到所有以最长长度结尾的位置，从这些位置开始回溯
	for i := 0; i < dim; i++ {
		if maxL == dp[i] {
			dfs_matrix(matrix, arr, i, path, 0) // 从位置 i 开始深度优先搜索回溯路径
		}
	}

}

// dfs_matrix 深度优先搜索回溯路径，打印所有可能的最长递增子序列
//
// 参数说明：
//   - matrix: 路径矩阵，matrix[i] 存储位置 i 的所有前驱位置
//   - arr: 原始数组
//   - s: 当前回溯的位置索引
//   - path: 当前路径数组，用于存储回溯过程中的元素
//   - idx: 当前路径的索引位置
//
// 算法说明：
//
//	使用深度优先搜索从终点（最长长度的位置）向前回溯到起点。
//	当 s >= 0 时，继续回溯；当 s < 0（即遇到 -1 标记）时，说明到达起点，打印路径。
//
// 回溯过程：
//  1. 将当前位置的元素加入路径：path[idx] = arr[s]
//  2. 遍历所有前驱位置，递归回溯
//  3. 当遇到 -1 标记时，说明到达起点，打印完整路径（需要逆序打印，因为是从后往前回溯的）
//
// 示例：
//
//	对于数组 [2,5,3,7]，假设最长长度为 3，路径 [2,5,7]：
//	- 从位置 3（值为 7）开始回溯
//	- matrix[3] = [1, -2]，前驱是位置 1（值为 5）
//	- 从位置 1 继续回溯，matrix[1] = [0, -2]，前驱是位置 0（值为 2）
//	- 从位置 0 继续回溯，matrix[0] = [-1, -2]，遇到 -1，打印路径 [2,5,7]
func dfs_matrix(matrix [][]int, arr []int, s int, path []int, idx int) {
	if s >= 0 {
		// 将当前位置的元素加入路径
		path[idx] = arr[s]
		// 遍历所有前驱位置，继续回溯
		for i := 0; matrix[s][i] != -2; i++ {
			dfs_matrix(matrix, arr, matrix[s][i], path, idx+1)
		}
	} else {
		// s < 0 表示遇到 -1 标记，到达起点，打印完整路径
		// 注意：路径是从后往前构建的，所以需要逆序打印
		fmt.Println("path: ")
		for i := idx - 1; i >= 0; i-- {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
	}
}

// pathOfLIS 计算最长递增子序列的长度，并构建路径矩阵用于回溯所有可能的路径
//
// 参数说明：
//   - arr: 输入数组
//   - dp: 输出参数，dp[i] 表示以 arr[i] 结尾的最长递增子序列长度
//   - path: 输出参数，路径矩阵，path[i] 存储所有可以转移到位置 i 的前驱位置索引
//
// 返回值：
//
//	最长递增子序列的长度
//
// 路径矩阵 path 的数据结构说明：
//   - path[i] 是一个数组，存储所有可以转移到位置 i 的前驱位置 j
//   - path[i][k] = j：表示从位置 j 可以转移到位置 i（即 arr[j] < arr[i] 且 dp[i] == dp[j] + 1）
//   - path[i][k] = -1：表示位置 i 是路径的起点（没有前驱，即 dp[i] == 1）
//   - path[i][k] = -2：表示路径数组的结束标记（用于在回溯时知道何时停止）
//
// 示例：
//
//	对于数组 [10,9,2,5,3,7,101,18]：
//	- dp[3] = 2（以 arr[3]=5 结尾，最长子序列是 [2,5]，长度为 2）
//	- path[3] = [1, -2]：表示位置 3 的前驱是位置 1（arr[1]=2）
//	- dp[4] = 2（以 arr[4]=3 结尾，最长子序列是 [2,3]，长度为 2）
//	- path[4] = [1, -2]：表示位置 4 的前驱是位置 1（arr[1]=2）
//	- dp[5] = 3（以 arr[5]=7 结尾，最长子序列是 [2,5,7] 或 [2,3,7]，长度为 3）
//	- path[5] = [3, 4, -2]：表示位置 5 的前驱可以是位置 3 或位置 4
//
// 算法流程：
// 1. 计算 dp 数组（与 lengthOfLIS 相同）
// 2. 对于每个位置 i，找出所有满足条件的前驱位置 j：
//   - arr[j] < arr[i]（满足递增条件）
//   - dp[i] == dp[j] + 1（说明从 j 转移到 i 是最优的）
//
// 3. 将这些前驱位置存储在 path[i] 中
// 4. 如果没有前驱（dp[i] == 1），则标记为 -1
func pathOfLIS(arr []int, dp []int, path [][]int) int {
	if len(arr) == 0 {
		return 0
	}
	dp[0] = 1
	maxL := dp[0]
	for i := 0; i < len(arr); i++ {
		dp[i] = 1 // 初始化为 1（至少包含自身）
		// 计算 dp[i]：以 arr[i] 结尾的最长递增子序列长度
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxL = max(maxL, dp[i])

		// 构建路径矩阵：找出所有可以转移到位置 i 的前驱位置
		k := 0
		for j := 0; j < i; j++ {
			// 如果从位置 j 可以转移到位置 i（满足递增且是最优转移）
			if arr[i] > arr[j] && dp[i] == dp[j]+1 {
				path[i][k] = j // 记录前驱位置
				k++
			}
		}
		// 如果没有前驱（dp[i] == 1，即位置 i 是某个路径的起点）
		if k == 0 {
			path[i][k] = -1 // 标记为起点
			k++
		}
		path[i][k] = -2 // 结束标记
	}
	return maxL
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
