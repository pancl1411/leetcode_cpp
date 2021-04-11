/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	// special case
	if len(s) <= 1 {
		return 0
	}
	// dp[i][j] 表示s[i,j]是回文子串
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	// 求解dp
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}
		}
	}
	// dp_num[i]表示s[0，i]串需要切割次数
	dp_num := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if dp[0][i] {
			dp_num[i] = 0
		} else {
			dp_num[i] = i
			for j := 1; j <= i; j++ {
				if dp[j][i] {
					dp_num[i] = min(dp_num[i], dp_num[j-1]+1)
				}
			}
		}
	}
	//
	return dp_num[len(s)-1]
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

// @lc code=end

