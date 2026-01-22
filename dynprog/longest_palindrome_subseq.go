package dyoprog

import "gopkg/util"

// 最长回文子序列

// 最长子序列类解法 https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484666&idx=1&sn=e3305be9513eaa16f7f1568c0892a468&chksm=9bd7faf2aca073e4f08332a706b7c10af877fee3993aac4dae86d05783d3d0df31844287104e&scene=21#wechat_redirect
// 最长回文子序列, 动态规划解法
func longestPalindromeSubseqDP(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	// dp[i][j]表示s[i..j]的最长回文子序列长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // 单个字符的最长回文子序列是1
	}
	// 从长度为2的子串开始计算
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				// 两端字符相等，最长回文子序列长度为中间部分+2
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 两端字符不相等，取去掉左端点或右端点的最大值
				dp[i][j] = util.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// 最长回文子序列, 递归解法（带记忆化）
func longestPalindromeSubseqRecursion(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dfs(s, 0, n-1, memo)
}

func dfs(s string, i, j int, memo [][]int) int {
	if i > j {
		return 0
	}
	if i == j {
		return 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s[i] == s[j] {
		memo[i][j] = dfs(s, i+1, j-1, memo) + 2
	} else {
		memo[i][j] = util.Max(dfs(s, i+1, j, memo), dfs(s, i, j-1, memo))
	}
	return memo[i][j]
}
