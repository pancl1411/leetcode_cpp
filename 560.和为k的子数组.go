/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	// 前缀和思路
	preSum := make(map[int]int, 0)
	preSum[0] = 1
	//
	var result, sum int = 0, 0
	for _, v := range nums {
		sum += v
		result += preSum[sum-k]
		preSum[sum]++
	}
	return result
}

// @lc code=end

