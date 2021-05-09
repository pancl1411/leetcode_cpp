/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	// special case
	if len(s) == 0 {
		return true
	}
	//
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || !isMatch(stack[len(stack)-1], s[i]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	//
	return len(stack) == 0
}

func isMatch(left, right byte) bool {
	if left == '(' && right == ')' {
		return true
	} else if left == '{' && right == '}' {
		return true
	} else if left == '[' && right == ']' {
		return true
	}
	return false
}

// @lc code=end

