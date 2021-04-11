/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// special case
	height := len(triangle)
	if height == 0 {
		return 0
	}
	// dp[i] 表示当前层终点为triangle[*][i]的最小路径和
	dp := make([]int, len(triangle[height-1]))
	dp[0] = triangle[0][0]
	//
	for i := 1; i < height; i++ {
		//内存循环倒序，避免dp[j-1]被新值覆盖了
		for j := i; j >= 0; j-- {
			if j == i {
				dp[j] = triangle[i][j] + dp[j-1]
			} else if j == 0 {
				dp[j] = triangle[i][j] + dp[j]
			} else {
				dp[j] = triangle[i][j] + min(dp[j-1], dp[j])
			}
		}
	}
	// 输出最小值
	result := dp[0]
	for i := 1; i < len(dp); i++ {
		result = min(result, dp[i])
	}
	return result
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

// @lc code=end

