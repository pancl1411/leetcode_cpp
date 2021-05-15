/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// positive, negtive
	var p, n, result int = max(0, nums[0]), min(0, nums[0]), nums[0]
	// 注意边界条件i，注意0的重置
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			p = 0
			n = 0
		} else if nums[i] < 0 {
			p, n = max(0, nums[i]*n), min(nums[i], nums[i]*p)
		} else if nums[i] > 0 {
			p, n = max(nums[i], nums[i]*p), min(0, nums[i]*n)
		}
		result = max(result, p)
	}
	//
	return result
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

// @lc code=end

