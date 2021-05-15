/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	// 不包含i的左侧乘积
	left := make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	// 左侧乘积
	right := 1
	answer := left[:] //指向同一个底层数组add
	//
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = left[i] * right
		right *= nums[i]
	}
	return answer
}

// @lc code=end

