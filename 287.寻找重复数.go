/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
// 思路：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/287xun-zhao-zhong-fu-shu-by-kirsche/
// (转换为双指针的题)
func findDuplicate(nums []int) int {
	//
	fast, slow := nums[nums[0]], nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	// 注意此处是0, 一步未走
	slow = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	//
	return slow
}

// @lc code=end

