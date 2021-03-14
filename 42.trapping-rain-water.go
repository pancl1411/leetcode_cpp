/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	// Solution I, with two pointer, time O(n), space O(1)
	return trapWithTwoPointer(height)
}

func trapWithTwoPointer(height []int) int {
	var left, right, left_max, right_max, water int
	for left, right = 0, len(height)-1; left <= right; {
		if left_max <= right_max {
			if left_max <= height[left] {
				left_max = height[left]
			} else {
				water += left_max - height[left]
			}
			left++
		} else {
			if right_max <= height[right] {
				right_max = height[right]
			} else {
				water += right_max - height[right]
			}
			right--
		}
	}
	//
	return water
}

// @lc code=end

