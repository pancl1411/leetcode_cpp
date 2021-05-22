/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
// https://leetcode-cn.com/problems/spiral-matrix/solution/cxiang-xi-ti-jie-by-youlookdeliciousc-3/
func spiralOrder(matrix [][]int) []int {
	//
	if len(matrix) == 0 {
		return nil
	}
	// 上下左右的边界
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	//
	answer := make([]int, 0)
	for {
		// 上边界
		for i := left; i <= right; i++ {
			answer = append(answer, matrix[up][i])
		}
		if up++; up > down {
			break
		}
		// 右边界
		for i := up; i <= down; i++ {
			answer = append(answer, matrix[i][right])
		}
		if right--; right < left {
			break
		}
		// 下边界
		for i := right; i >= left; i-- {
			answer = append(answer, matrix[down][i])
		}
		if down--; down < up {
			break
		}
		// 左边界
		for i := down; i >= up; i-- {
			answer = append(answer, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}
	//
	return answer
}

// @lc code=end

