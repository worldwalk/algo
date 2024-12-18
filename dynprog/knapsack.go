package dyoprog

import "gopkg/util"

// 0-1 背包问题
func knapsack01(w []int, val []int, capacity int) int {
	dp := make([][]int, len(val)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = util.Max(dp[i-1][j], dp[i-1][j-w[i-1]]+val[i-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
