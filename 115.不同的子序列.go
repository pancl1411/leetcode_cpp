/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	// special case
	if len(s) < len(t) {
		return 0
	}
	// dp[i][j] 表示t的前i个元素子串，在s的前j个元素子序列个数
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	// loop to parse dp
	for i := 0; i < len(dp); i++ {
		for j := i; j < len(dp[i]); j++ {
			// cpu分支预测，高概率写前面
			if i != 0 {
				if t[i-1] != s[j-1] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
				}
			} else {
				// i == 0, 空串是任何串的子序列
				dp[i][j] = 1
			}
		}
	}
	//
	return dp[len(t)][len(s)]
}

// @lc code=end

