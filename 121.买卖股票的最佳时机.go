/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// special case
	if len(prices) <= 1 {
		return 0
	}
	//
	var buy, profit int
	buy = prices[0]
	//
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			profit = max(profit, prices[i]-buy)
		}
	}
	//
	return profit
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

// @lc code=end

