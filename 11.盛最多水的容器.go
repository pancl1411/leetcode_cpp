/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	var result, left, right, area int
	// 滑动窗口
	right = len(height) - 1
	for left < right {
		// update
		area = (right - left) * min(height[left], height[right])
		result = max(result, area)
		// move
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	//
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

