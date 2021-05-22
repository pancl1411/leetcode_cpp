/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	//
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}
	// 注意这里的abs操作，因为可能被前面的数改了正负号
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i])
		if idx <= length {
			nums[idx-1] = -abs(nums[idx-1])
		}
	}
	//
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	//
	return length + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

