/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
// 数组分3个段：前面递增段，中间乱序段，后面递增段
// 注意分段的边界条件
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var unorderMin, unorderMax int = math.MaxInt32, math.MinInt32
	var left, right int = 0, len(nums) - 1
	// 获取乱序中最小值 O(n)
	var upFlag bool = true
	for i := 1; i < len(nums); i++ {
		if upFlag && nums[i] <= nums[i-1] {
			upFlag = false
		}
		if !upFlag {
			unorderMin = min(unorderMin, nums[i])
		}
	}
	// 获取乱序中最大值 O(n)
	var downFlag bool = true
	for i := len(nums) - 2; i >= 0; i-- {
		if downFlag && nums[i] >= nums[i+1] {
			downFlag = false
		}
		if !downFlag {
			unorderMax = max(unorderMax, nums[i])
		}
	}
	// 特殊情况: 已经是升序
	if unorderMin > unorderMax {
		return 0
	}
	// 分别确定左/右边界
	for left < len(nums) && nums[left] <= unorderMin {
		left++
	}
	for right >= 0 && nums[right] >= unorderMax {
		right--
	}
	if left >= right {
		return 0
	}
	return right - left + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

