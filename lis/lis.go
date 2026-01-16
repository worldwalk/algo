package lis

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

// print all lis path
func printLISPath(arr []int) {
	dim := len(arr)
	dp := make([]int, dim)
	matrix := make([][]int, dim)
	for i := 0; i < dim; i++ {
		matrix[i] = make([]int, dim)
	}
	maxL := pathOfLIS(arr, dp, matrix)
	fmt.Println("maxL", maxL)

	path := make([]int, dim)
	for i := 0; i < dim; i++ {
		if maxL == dp[i] {
			dfs(matrix, arr, i, path, 0)
		}
	}

}

func dfs(matrix [][]int, arr []int, s int, path []int, idx int) {
	if s >= 0 {
		path[idx] = arr[s]
		for i := 0; matrix[s][i] != -2; i++ {
			dfs(matrix, arr, matrix[s][i], path, idx+1)
		}
	} else {
		fmt.Println("path: ")
		for i := idx - 1; i >= 0; i-- {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
	}
}

func pathOfLIS(arr []int, dp []int, path [][]int) int {
	if len(arr) == 0 {
		return 0
	}
	dp[0] = 1
	maxL := dp[0]
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxL = max(maxL, dp[i])
		k := 0
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] == dp[j]+1 {
				path[i][k] = j
				k++
			}
		}
		if k == 0 {
			path[i][k] = -1
			k++
		}
		path[i][k] = -2
	}
	return maxL
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
