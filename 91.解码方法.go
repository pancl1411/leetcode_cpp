/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	// dp[i]表示s的前i个字符解码总数
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	// 此题细节较多，需要注意
	for i := 2; i < len(dp); i++ {
		if s[i-1] == '0' {
			if s[i-2] == '1' || s[i-2] == '2' {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	//
	return dp[len(dp)-1]
}

// @lc code=end

