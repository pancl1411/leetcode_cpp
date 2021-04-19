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
	dp[1] = nums[0] //注意起点不同
	//
	for i := 2; i < len(dp); i++ {
		// 偷窃第 k 间房屋，那么就不能偷窃第 k-1 间房屋
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
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

