/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
// 思路参考：https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
func nextPermutation(nums []int) {
	// special case
	if len(nums) <= 1 {
		return
	}
	//
	small, next, big := len(nums)-2, len(nums)-1, len(nums)-1
	// 从end开始找
	for small >= 0 && nums[small] >= nums[next] {
		small--
		next--
	}
	// 存在下一个更大的排列
	if small >= 0 {
		for nums[big] <= nums[small] {
			big--
		}
		// swap
		nums[big], nums[small] = nums[small], nums[big]
	}
	// 后半部分改升序
	end := len(nums) - 1
	for next < end {
		nums[next], nums[end] = nums[end], nums[next]
		next++
		end--
	}
}

// @lc code=end

