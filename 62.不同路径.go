/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// special case
	if m <= 0 || n <= 0 {
		return 0
	}
	// dp[i][j] 是到 [i+1,j+1] 的路径数量
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	//
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 边界
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//
	return dp[m-1][n-1]
}

// @lc code=end

