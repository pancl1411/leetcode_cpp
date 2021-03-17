/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
// 思路参考：https://leetcode-cn.com/problems/maximal-rectangle/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-1-8/
func maximalRectangle(matrix [][]byte) int {
	// special case
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// var statement
	var maxArea int = 0
	heights := make([]int, len(matrix[0]))
	left := make([]int, len(matrix[0]))
	right := make([]int, len(matrix[0]))
	var zeroIndex int = -1
	// O(mn) time cost
	for i := 0; i < len(matrix); i++ {
		// j++ loop
		zeroIndex = -1
		for j := 0; j < len(matrix[i]); j++ {
			// init values
			if i == 0 {
				left[j] = -1
				right[j] = len(matrix[0])
			}
			// update height and left
			if matrix[i][j] == '1' {
				heights[j] += 1
				left[j] = maxInt(left[j], zeroIndex)
			} else {
				heights[j] = 0
				left[j] = -1
				zeroIndex = j
			}
		}
		// j-- loop
		zeroIndex = len(matrix[i])
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			// update right
			if matrix[i][j] == '1' {
				right[j] = minInt(right[j], zeroIndex)
			} else {
				right[j] = len(matrix[i])
				zeroIndex = j
			}
			// update maxArea
			maxArea = maxInt(maxArea, (right[j]-left[j]-1)*heights[j])
		}
	}
	//
	return maxArea
}

func minInt(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func maxInt(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

// @lc code=end

