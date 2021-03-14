/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// special case
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	// dp[i][j]是起点到[i+1,j+1]位置的路径数
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	// 初始条件
	dp[0][0] = 1
	//
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			// 特殊情况
			if obstacleGrid[i][j] == 1 || (i == 0 && j == 0) {
				continue
				// 边界一
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
				//边界二
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

// @lc code=end

