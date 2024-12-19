package lcs

import "gopkg/util"

// longest common subsequence

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
