/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	var answer string
	// special case
	if len(s) == 0 {
		return answer
	}
	// dp[i][j]表示s[i]到s[j]的子串是否回文串
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	// dp for solution
	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			// CASE 1 : 单个字母
			if i == j {
				dp[i][j] = true
				// CASE 2 : 两个字母，相等则是回文串
			} else if j == i+1 {
				dp[i][j] = (s[i] == s[j])
				// CASE 3 ： 两个以上字母，两头相等且内部是回文串
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}
			// 更新最长回文子串结果
			if dp[i][j] && ((j - i + 1) > len(answer)) {
				answer = s[i : j+1]
			}
		}
	}
	//
	return answer
}

// @lc code=end

