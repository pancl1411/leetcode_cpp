/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start

/*
//这个解法内存O(n^2), 会存在内存不足问题
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	//dp[i][j]表示[i, j]内操作一次的最大利润
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		// ** 这个解法内存O(n^2), 会存在内存不足问题
		dp[i] = make([]int, len(prices))
	}
	//
	var buy int
	for i := 0; i < len(prices); i++ {
		buy = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j] < buy {
				buy = prices[j]
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], prices[j]-buy)
			}
		}
	}
	//
	profit := dp[0][len(prices)-1]
	for i := 1; i < len(prices)-1; i++ {
		profit = max(profit, dp[0][i]+dp[i][len(prices)-1])
	}
	return profit
}
*/

// 官方思路：
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iii-by-wrnt/
func maxProfit(prices []int) int {
	//
	if len(prices) <= 1 {
		return 0
	}
	//
	buy1, sale1 := prices[0], 0
	buy2, sale2 := prices[0], 0
	//
	for i := 0; i < len(prices); i++ {
		buy1 = min(prices[i], buy1)
		sale1 = max(prices[i]-buy1, sale1)
		// 此处注意buy2的状态转移方程
		buy2 = min(prices[i]-sale1, buy2)
		sale2 = max(prices[i]-buy2, sale2)
	}
	//
	return sale2
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

// @lc code=end

