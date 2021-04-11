/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	// special case
	if len(s3) != (len(s1) + len(s2)) {
		return false
	}
	// dp init
	// dp[i][j]表示s1的前i个，s2的前j个字符串，是否构成s3前i+j的交错字符串
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	//
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			// CPU分支预测，高概率的if写前面
			if i != 0 && j != 0 {
				// 状态转移方程
				dp[i][j] = (dp[i][j-1] && (s2[j-1] == s3[i+j-1])) || dp[i-1][j] && (s1[i-1] == s3[i+j-1])
			} else if i == 0 && j == 0 {
				// 以下均是特殊条件
				dp[0][0] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && (s2[j-1] == s3[j-1])
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && (s1[i-1] == s3[i-1])
			}
		}
	}
	//
	return dp[len(s1)][len(s2)]
}

// @lc code=end

