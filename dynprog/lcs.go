package dyoprog

import "gopkg/util"

// longest common subsequence

// lengthOfLCS 计算两个字符串的最长公共子序列（Longest Common Subsequence, LCS）的长度
//
// 题目描述：
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
// 如果不存在公共子序列，返回 0。
// 子序列是指：不改变剩余字符顺序的情况下，删除某些（也可以不删除）字符后组成的新序列。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//
// 示例：
//
//	输入: text1 = "abcde", text2 = "ace"
//	输出: 3
//	解释: 最长公共子序列是 "ace"，它的长度为 3
//
//	输入: text1 = "abc", text2 = "abc"
//	输出: 3
//	解释: 最长公共子序列是 "abc"，它的长度为 3
//
//	输入: text1 = "abc", text2 = "def"
//	输出: 0
//	解释: 两个字符串没有公共子序列，返回 0
//
// 算法分析：
// 使用动态规划方法解决：
// 1. 状态定义：dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度
//   - dp[0][j] = 0 和 dp[i][0] = 0 表示空字符串与任何字符串的 LCS 长度为 0
//
// 2. 状态转移：
//   - 如果 text1[i-1] == text2[j-1]：当前字符匹配，dp[i][j] = dp[i-1][j-1] + 1
//   - 如果 text1[i-1] != text2[j-1]：当前字符不匹配，取两种情况的最大值：
//   - 忽略 text1 的当前字符：dp[i-1][j]
//   - 忽略 text2 的当前字符：dp[i][j-1]
//   - dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//
// 3. 初始状态：dp[0][j] = 0, dp[i][0] = 0（空字符串的情况）
// 4. 最终答案：dp[len(text1)][len(text2)]
//
// 时间复杂度：O(m * n)，其中 m 和 n 分别是两个字符串的长度
//   - 需要填充一个 (m+1) × (n+1) 的二维数组
//   - 每个状态的计算时间为 O(1)
//
// 空间复杂度：O(m * n)，用于存储 dp 二维数组
//
// 空间优化思路：
// 可以使用滚动数组优化到 O(min(m, n)) 的空间复杂度，因为当前状态只依赖于上一行的状态。
// 进一步优化可以使用一维数组，只保留当前行和上一行的信息。
func lengthOfLCS(a, b string) int {
	dp := make([][]int, len(a)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = util.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(a)][len(b)]
}

func findLCS(a, b string) string {
	dp := make([][]int, len(a)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = util.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	i, j := len(a), len(b)
	var lcs []byte
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			lcs = append(lcs, a[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	var rlcs []byte
	for i := len(lcs) - 1; i >= 0; i-- {
		rlcs = append(rlcs, lcs[i])
	}
	return string(rlcs)
}
