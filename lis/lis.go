package lis

import "fmt"

// LongestIncreasingSubsequence
// https://writings.sh/post/algorithm-longest-increasing-subsequence

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
