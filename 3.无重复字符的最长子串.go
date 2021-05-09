/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

// tricks : 滑动窗口，右侧边界递增，左侧边界视map情况增长
func lengthOfLongestSubstring(s string) int {
	has := make(map[byte]int)
	var start, maxLen int
	// 滑动窗口
	for end := 0; end < len(s); end++ {
		if idx, ok := has[s[end]]; ok {
			start = max(start, idx+1)
		}
		has[s[end]] = end
		maxLen = max(maxLen, end-start+1)
	}
	//
	return maxLen
}

func max(f, s int) int {
	if f > s {
		return f
	}
	return s
}

// @lc code=end

