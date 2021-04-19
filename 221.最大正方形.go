/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxSide int = 0
	//
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if maxSide == 0 && dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	//
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = 1 + min(min(dp[i-1][j-1], dp[i][j-1]), min(dp[i-1][j-1], dp[i-1][j]))
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}
	//
	return maxSide * maxSide
}

func min(f, s int) int {
	if f < s {
		return f
	}
	return s
}

func max(f, s int) int {
	if f > s {
		return f
	}
	return s
}

// @lc code=end

