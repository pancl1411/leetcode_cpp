/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
var result int

func maxCoins(nums []int) int {
	//
	result = 0
	selected := make([]bool, len(nums))
	backtrace(nums, selected, 0, 0)
	return result
}

// 回溯法会超时，需要优化
func backtrace(nums []int, selected []bool, count, coin int) {
	//
	if count >= len(nums) {
		if coin > result {
			result = coin
		}
		return
	}
	//
	for i := 0; i < len(nums); i++ {
		if !selected[i] {
			//
			selected[i] = true
			// calcute
			left, right := 1, 1
			for j := i - 1; j >= 0; j-- {
				if !selected[j] {
					left = nums[j]
					break
				}
			}
			for j := i + 1; j < len(nums); j++ {
				if !selected[j] {
					right = nums[j]
					break
				}
			}
			add := left * nums[i] * right
			//
			backtrace(nums, selected, count+1, coin+add)
			// backtrace
			selected[i] = false
		}
	}
}

// @lc code=end

