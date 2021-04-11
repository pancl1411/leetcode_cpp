/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	//https://leetcode-cn.com/circle/article/qiAgHn/
	// 优化三维数组为2个二维数组
	if len(prices) <= 1 {
		return 0
	}
	//
	k = min(k, len(prices)/2)
	buy := make([][]int, len(prices))
	sale := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		buy[i] = make([]int, k+1)
		sale[i] = make([]int, k+1)
	}
	// init
	for j := 1; j <= k; j++ {
		buy[0][j] = -prices[0]
	}
	//
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			sale[i][j] = max(sale[i-1][j], buy[i-1][j]+prices[i])
			buy[i][j] = max(buy[i-1][j], sale[i-1][j-1]-prices[i])
		}
	}
	//
	return sale[len(prices)-1][k]
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

// @lc code=end

