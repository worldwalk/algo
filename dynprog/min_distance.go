package dyoprog

import "gopkg/util"

// 编辑距离(Edit Distance)算法 - 莱文斯坦距离(Levenshtein Distance)
// 问题描述：计算将一个字符串转换为另一个字符串所需的最少编辑操作次数
// 允许的操作：插入(Insert)、删除(Delete)、替换(Replace)
// 应用场景：拼写检查、DNA序列比对、文本相似度分析等

// 动态规划解法 - 最优解，时间复杂度O(m*n)，空间复杂度O(m*n)
// 其中m和n分别是两个输入字符串的长度
func minDistance(word1 string, word2 string) int {
	// 创建二维DP数组，dp[i][j]表示将word1的前i个字符转换为word2的前j个字符所需的最少操作数
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	// 初始化边界条件：当其中一个字符串为空时，编辑距离等于另一个字符串的长度
	// 即需要插入全部字符或删除全部字符
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i // word2为空，需要删除word1的i个字符
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j // word1为空，需要插入word2的j个字符
	}

	// 填充DP数组，自底向上计算所有子问题的解
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// 如果当前字符相同，则不需要操作，直接继承前一个状态
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 当前字符不同，需要选择插入、删除或替换中操作数最少的一种
				// dp[i-1][j-1]+1: 替换操作
				// dp[i-1][j]+1: 删除操作（删除word1的当前字符）
				// dp[i][j-1]+1: 插入操作（在word1中插入word2的当前字符）
				dp[i][j] = util.Min(dp[i-1][j-1], util.Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	// 返回最终结果：将整个word1转换为整个word2所需的最少操作数
	return dp[len(dp)-1][len(dp[0])-1]
}

// 递归辅助函数 - 基于自顶向下的递归解法
// 参数i: word1的当前索引（从后往前）
// 参数j: word2的当前索引（从后往前）
func dp(i, j int, word1, word2 string) int {
	// 边界条件：如果word1处理完了，需要插入word2剩余的j+1个字符
	if i == -1 {
		return j + 1
	}
	// 边界条件：如果word2处理完了，需要删除word1剩余的i+1个字符
	if j == -1 {
		return i + 1
	}

	// 如果当前字符相同，直接递归处理剩余部分
	if word1[i] == word2[j] {
		return dp(i-1, j-1, word1, word2)
	}

	// 当前字符不同，选择插入、删除或替换中操作数最少的一种
	return util.Min(dp(i-1, j-1, word1, word2), util.Min(dp(i-1, j, word1, word2), dp(i, j-1, word1, word2))) + 1
}

// minDistance2 - 递归解法的入口函数
// 时间复杂度O(3^max(m,n)) - 指数级，空间复杂度O(max(m,n))
// 缺点：存在大量重复计算，效率极低，仅用于理解算法原理
func minDistance2(word1 string, word2 string) int {
	return dp(len(word1)-1, len(word2)-1, word1, word2)
}

// minDistanceRecursion - 另一种形式的递归解法
// 从字符串开头开始处理，与minDistance2的思路类似但实现方式不同
// 同样存在指数级时间复杂度问题
func minDistanceRecursion(word1 string, word2 string) int {
	// 边界条件：如果word1为空，需要插入word2的所有字符
	if len(word1) == 0 {
		return len(word2)
	}
	// 边界条件：如果word2为空，需要删除word1的所有字符
	if len(word2) == 0 {
		return len(word1)
	}

	// 如果当前字符相同，直接递归处理剩余部分
	if word1[0] == word2[0] {
		return minDistanceRecursion(word1[1:], word2[1:])
	}

	// 当前字符不同，选择插入、删除或替换中操作数最少的一种
	// minDistanceRecursion(word1[1:], word2[1:])+1: 替换操作
	// minDistanceRecursion(word1[1:], word2)+1: 删除操作
	// minDistanceRecursion(word1, word2[1:])+1: 插入操作
	return util.Min(minDistanceRecursion(word1[1:], word2[1:]), util.Min(minDistanceRecursion(word1[1:], word2), minDistanceRecursion(word1, word2[1:]))) + 1
}
