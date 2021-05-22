/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
var answer int

func findTargetSumWays(nums []int, target int) int {
	answer = 0
	dfs(nums, target, 0, 0)
	return answer
}

func dfs(nums []int, target, start, sum int) {
	//
	if start >= len(nums) {
		if target == sum {
			answer++
		}
		return
	}
	//
	dfs(nums, target, start+1, sum+nums[start])
	dfs(nums, target, start+1, sum-nums[start])
}

// @lc code=end

