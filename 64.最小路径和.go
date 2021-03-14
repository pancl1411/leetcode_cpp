/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	//
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	//
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				} else {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				}
			}
		}
	}
	//
	return dp[len(dp)-1][len(dp[0])-1]
}

// @lc code=end

