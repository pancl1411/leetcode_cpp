/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	//
	if len(temperatures) == 0 {
		return nil
	}
	// 递减栈，存的数组idx
	stack := make([]int, 0)
	answer := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			answer[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	// 小技巧，利用golang的默认0值
	//for len(stack) > 0 {
	//	answer[stack[len(stack)-1]] = 0
	//	stack = stack[:len(stack)-1]
	//}
	//
	return answer

}

// @lc code=end

