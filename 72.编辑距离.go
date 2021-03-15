/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// special case
	if (len(word1) == 0) && (len(word2) == 0) {
		return 0
	} else if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}
	// dp[i][j] 表示word1前i个元素和word2前j个元素的最少操作数
	// 		注意i、j从0开始计数，比word多一个元素
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	// dp with O(mn)
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 {
				// CASE 1: 边界：word1未选元素
				dp[i][j] = j
			} else if j == 0 {
				// CASE 2: 边界：word2未选元素
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				// CASE 3 : word1和word2的第i/j个元素相等，不增加操作
				dp[i][j] = dp[i-1][j-1]
			} else {
				// CASE 4 : 增/删/改选一个，取最小值
				dp[i][j] = 1 + MinFunc(dp[i-1][j-1], MinFunc(dp[i][j-1], dp[i-1][j]))
			}
		}
	}
	//
	return dp[len(dp)-1][len(dp[0])-1]
}

func MinFunc(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

// @lc code=end

