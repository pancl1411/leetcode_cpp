/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// special case
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	// dp[i][j] 表示至s[i+1] p[j+1]是否是匹配的
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	//
	for i := 0; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if i == 0 {
				dp[0][j] = (p[j-1] == '*') && dp[0][j-2]
			} else if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 注意这里的处理
				// CASE 3: pattern为'*', 可分为子情况：
				//  1、 *的前一字符，和s的当前字符不等，a*只能当空处理，a*==null
				//  2、 *的前一字符，和s的当前字符相等，a*可以：
				//          1)当空a*==null；  dp[i][j-2]
				//          2）当单个前一字符 a*==a ;  dp[i-1][j]
				//          3) 当多个前一字符a* == aa; dp[i-1][j-1]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = (dp[i][j-2] || dp[i-1][j] || dp[i-1][j-1])
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				dp[i][j] = (s[i-1] == p[j-1] && dp[i-1][j-1])
			}
		}
	}
	//
	return dp[len(dp)-1][len(dp[0])-1]
}

// @lc code=end

