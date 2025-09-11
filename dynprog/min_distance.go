package dyoprog

import "gopkg/util"

// 编辑距离 - 典型的动态规划问题

// 动态规划解法
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = util.Min(dp[i-1][j-1], util.Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func dp(i, j int, word1, word2 string) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if word1[i] == word2[j] {
		return dp(i-1, j-1, word1, word2)
	}
	return util.Min(dp(i-1, j-1, word1, word2), util.Min(dp(i-1, j, word1, word2), dp(i, j-1, word1, word2))) + 1
}

func minDistance2(word1 string, word2 string) int {
	return dp(len(word1)-1, len(word2)-1, word1, word2)
}

func minDistanceRecursion(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	if word1[0] == word2[0] {
		return minDistanceRecursion(word1[1:], word2[1:])
	}
	return util.Min(minDistanceRecursion(word1[1:], word2[1:]), util.Min(minDistanceRecursion(word1[1:], word2), minDistanceRecursion(word1, word2[1:]))) + 1
}
