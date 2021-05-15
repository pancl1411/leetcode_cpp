/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
// 思路：https://leetcode-cn.com/problems/perfect-squares/solution/hua-jie-suan-fa-279-wan-quan-ping-fang-shu-by-guan/
func numSquares(n int) int {
	//
	dp := make([]int, n+1)
	dp[0] = 0
	//
	for i := 1; i < len(dp); i++ {
		// init, 最差情况全1相加
		dp[i] = i
		// 从2开始, 做减法
		for j := 2; j*j <= i; j++ {
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[len(dp)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

