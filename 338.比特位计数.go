/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		//正整数 xx，将其二进制表示右移一位，
		//等价于将其二进制表示的最低位去掉
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}

// @lc code=end

