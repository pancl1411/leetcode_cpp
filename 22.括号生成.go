/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
var answers []string

func generateParenthesis(n int) []string {
	answers = []string{}
	backtrace(n, 0, 0, "")
	return answers
}

func backtrace(n, left, right int, s string) {
	//
	if left >= n && right >= n {
		answers = append(answers, s)
	}
	//
	if left < n {
		backtrace(n, left+1, right, s+"(")
	}
	// 注意这里的right < left保证括号有效性
	if right < left && right < n {
		backtrace(n, left, right+1, s+")")
	}
}

// @lc code=end

