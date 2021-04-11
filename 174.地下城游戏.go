/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	// special case
	if len(dungeon) == 0 {
		return 1
	}
	//
	dp := make([][]int, len(dungeon))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(dungeon[0]))
	}
	//
	var chose int
	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp[i]) - 1; j >= 0; j-- {
			// 前三项是递推公式的特殊情况
			if i == len(dp)-1 && j == len(dp[i])-1 {
				chose = 1
			} else if i == len(dp)-1 {
				chose = dp[i][j+1]
			} else if j == len(dp[i])-1 {
				chose = dp[i+1][j]
			} else {
				// 递推公式
				chose = min(dp[i][j+1], dp[i+1][j])
			}
			dp[i][j] = max(chose-dungeon[i][j], 1)
		}
	}
	//
	return dp[0][0]
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

