/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
var result map[string]struct{}

func removeInvalidParentheses(s string) []string {
	if len(s) == 0 {
		return nil
	}
	// 计算分别删除多少的左/右括号
	var leftRemove, rightRemove int = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftRemove++
		} else if s[i] == ')' {
			if leftRemove > 0 {
				leftRemove--
			} else {
				rightRemove++
			}
		}
	}
	// 需要利用map对答案去重
	result = make(map[string]struct{}, 0)
	backtrace(s, 0, leftRemove, rightRemove, 0, 0, []byte{})
	answer := make([]string, 0)
	for k, _ := range result {
		answer = append(answer, k)
	}
	return answer
}

func backtrace(s string, idx, leftRemove, rightRemove, leftAdd, rightAdd int, selected []byte) {
	//
	if idx >= len(s) {
		if leftRemove == 0 && rightRemove == 0 {
			result[string(selected)] = struct{}{}
		}
		return
	}
	// 删除1个')'
	if s[idx] == ')' && rightRemove > 0 {
		backtrace(s, idx+1, leftRemove, rightRemove-1, leftAdd, rightAdd, selected)
	}
	// 删除1个'('
	if s[idx] == '(' && leftRemove > 0 {
		backtrace(s, idx+1, leftRemove-1, rightRemove, leftAdd, rightAdd, selected)
	}
	// 不删除
	selected = append(selected, s[idx])
	if s[idx] != '(' && s[idx] != ')' {
		backtrace(s, idx+1, leftRemove, rightRemove, leftAdd, rightAdd, selected)
	} else if s[idx] == '(' {
		backtrace(s, idx+1, leftRemove, rightRemove, leftAdd+1, rightAdd, selected)
	} else if s[idx] == ')' && rightAdd < leftAdd {
		//注意此处右括号数量不能超过左括号
		backtrace(s, idx+1, leftRemove, rightRemove, leftAdd, rightAdd+1, selected)
	}
	selected = selected[:len(selected)-1] //其实可以不删除
}

// @lc code=end

