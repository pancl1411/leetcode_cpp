/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	left := make([]int, len(ratings))
	// left
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	//
	result, right := 0, 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		result += max(right, left[i])
	}
	//
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

