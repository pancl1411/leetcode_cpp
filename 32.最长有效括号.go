/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// special case
	if len(s) < 2 {
		return 0
	}
	// dp[i] 是以s[i]结尾的最长有效子串长度
	dp := make([]int, len(s))
	var maxLen int = 0
	//
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		// '('结尾的字符串不可能是有效的
		if s[i] == '(' {
			continue
		}
		// s[i] == ')'时, 看前面的字符
		if s[i-1] == '(' {
			if i > 1 {
				dp[i] = 2 + dp[i-2]
			} else {
				dp[i] = 2
			}
		} else if (dp[i-1] > 0) && (i > dp[i-1]) && s[i-dp[i-1]-1] == '(' {
			if (i - dp[i-1] - 2) >= 0 {
				dp[i] = 2 + dp[i-1] + dp[i-dp[i-1]-2]
			} else {
				dp[i] = 2 + dp[i-1]
			}
		}
		// 更新最长子串
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	//
	return maxLen
}

// @lc code=end

