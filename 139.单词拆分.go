/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// map init
	dict := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	//
	dp := make([]bool, len(s))
	//
	for i := 0; i < len(s); i++ {
		//
		if _, ok := dict[string(s[0:i+1])]; ok {
			dp[i] = true
			continue
		}
		// 注意dp[j-1], i+1的边界
		for j := 1; j <= i; j++ {
			_, ok := dict[string(s[j:i+1])]
			if dp[j-1] && ok {
				dp[i] = true
				continue
			}
		}
	}
	//
	return dp[len(s)-1]
}

// @lc code=end

