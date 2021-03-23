/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	//
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	// dp[l][i][j]
	//	表示以s1[i],s2[j]开头，长度为l+1的子串是否是扰乱字符串
	dp := make([][][]bool, len(s1))
	for l := 0; l < len(dp); l++ {
		dp[l] = make([][]bool, len(s1))
		for i := 0; i < len(dp[l]); i++ {
			dp[l][i] = make([]bool, len(s1))
		}
	}
	// O(n^4) time cost
	for l := 0; l < len(dp); l++ {
		for i := 0; i < len(dp[l])-l; i++ {
			for j := 0; j < len(dp[l][i])-l; j++ {
				if l == 0 {
					// CASE: 一个字符，相等则是
					//  注意：必须放在下面for循环外，否则l==0时无法进入
					dp[l][i][j] = (s1[i] == s2[j])
					continue
				}
				for mid := 0; mid < l; mid++ {
					if dp[mid][i][j] && dp[l-mid-1][i+mid+1][j+mid+1] {
						// CASE: 中间切割，前/后半段分别是
						dp[l][i][j] = true
						break
					} else if dp[mid][i][j+l-mid] && dp[l-mid-1][i+mid+1][j] {
						// CASE：中间切割，前/后半段扰乱后是
						// 	 注意：前后半段的起点计算
						dp[l][i][j] = true
						break
					}
				}
			}
		}
	}
	//
	return dp[len(dp)-1][0][0]
}

// @lc code=end

