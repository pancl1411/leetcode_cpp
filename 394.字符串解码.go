/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	stack := make([]byte, 0)
	//
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			// 字符出栈
			idx := len(stack) - 1
			for ; stack[idx] != '['; idx-- {
			}
			chars := string(stack[idx+1:]) // 注意此处的chars转成string，否则slice存在覆盖问题
			stack = stack[:idx]
			// 数字出栈
			// 注意：可能出现'12'的多个数字
			idx = len(stack) - 1
			for ; idx >= 0; idx-- {
				if !(stack[idx] >= '0' && stack[idx] <= '9') {
					break
				}
			}
			// 数字转换
			digits := string(stack[idx+1:])
			stack = stack[:idx+1]
			// 倍数相乘
			num, _ := strconv.Atoi(digits)
			for i := 1; i <= num; i++ {
				stack = append(stack, chars...)
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

// @lc code=end

