package dyoprog

import "gopkg/util"

// 0-1背包问题：
// 问题描述：给定一个容量为capacity的背包和n个物品，每个物品有重量w[i]和价值val[i]，
// 每个物品只能选择放或不放，求在不超过背包容量的前提下，能获得的最大价值。
//
// 动态规划解法：
// 1. 状态定义：dp[i][j]表示考虑前i个物品，背包容量为j时能获得的最大价值
// 2. 状态转移：
//   - 如果第i个物品的重量超过当前容量j：dp[i][j] = dp[i-1][j]
//   - 否则，可以选择放或不放：dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i-1]] + val[i-1])
//
// 3. 边界条件：
//   - dp[0][j] = 0（考虑0个物品，价值为0）
//   - dp[i][0] = 0（容量为0，无法放任何物品）
func knapsack01(w []int, val []int, capacity int) int {
	// 创建二维DP数组：(物品数量+1) x (容量+1)
	// 多一维是为了处理i=0（0个物品）的情况
	dp := make([][]int, len(val)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}

	// 填充DP数组
	for i := 1; i < len(dp); i++ { // i从1开始，对应第i-1个实际物品
		for j := 1; j < len(dp[0]); j++ { // j从1开始，遍历所有可能的容量
			if j < w[i-1] { // 当前物品重量超过容量，不能放
				dp[i][j] = dp[i-1][j]
			} else { // 可以选择放或不放，取最大值
				dp[i][j] = util.Max(dp[i-1][j], dp[i-1][j-w[i-1]]+val[i-1])
			}
		}
	}

	// 返回最终结果：考虑所有物品，容量为capacity时的最大价值
	return dp[len(dp)-1][len(dp[0])-1]
}
