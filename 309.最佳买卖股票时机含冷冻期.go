/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i] is :
	//		dp[i][0] : 最终卖出状态（冷冻期，即昨天卖的）
	//		dp[i][1] ：最终卖出状态（非冷冻期）
	//		dp[i][2] ：已买入
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][2] = -prices[0]
	//
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][2] + prices[i]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
	}
	// 最终卖出的状态
	return max(dp[len(dp)-1][0], dp[len(dp)-1][1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

