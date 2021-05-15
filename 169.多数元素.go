/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	// 投票算法
	candidate, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
			if count < 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}
	//
	return candidate
}

// @lc code=end

