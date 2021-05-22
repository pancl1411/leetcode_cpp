/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	//
	if len(s) <= 1 {
		return len(s)
	}
	//
	var result int = 0
	for i := 0; i < len(s); i++ {
		// 自己为中心，两边拓展(j==0时只有自己) [i-j, i+j]
		for j := 0; i-j >= 0 && i+j <= len(s)-1 && s[i-j] == s[i+j]; j++ {
			result++
		}
		// (i-1, i)为中心
		for j := 0; i-j-1 >= 0 && i+j <= len(s)-1 && s[i-j-1] == s[i+j]; j++ {
			result++
		}
	}
	//
	return result
}

// @lc code=end

