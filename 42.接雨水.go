/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// special case
	if len(height) <= 1 {
		return 0
	}
	//
	l_max, l_idx, r_max, r_idx, result := height[0], 0, height[len(height)-1], len(height)-1, 0
	// 1、要么要有<=等号，idx递增放外面，因为取到最大值后r_idx--，r_idx不一定计算了
	// 2、要么<无等号，更新max时不进行idx操作
	for l_idx < r_idx {
		if l_max <= r_max {
			if height[l_idx] > l_max {
				l_max = height[l_idx]
			} else {
				result += l_max - height[l_idx]
				l_idx++
			}
		} else {
			if height[r_idx] > r_max {
				r_max = height[r_idx]
			} else {
				result += r_max - height[r_idx]
				r_idx--
			}
		}
	}
	//
	return result
}

// @lc code=end

