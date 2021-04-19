/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	//
	if len(nums) == 0 {
		return 0
	}
	//
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	//
	for i := 1; i < len(nums); i++ {
		// 偷窃第 k 间房屋，那么就不能偷窃第 k-1 间房屋
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}
	//
	return dp[len(dp)-1]
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

// @lc code=end

