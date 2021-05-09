/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

var digit_map map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var answers []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	// answer需要在函数内初始化,下个case的会还在
	answers = []string{}
	backtrace(digits, 0, "")
	return answers
}

func backtrace(digits string, start int, combine string) {
	//
	if start >= len(digits) {
		answers = append(answers, combine)
		return
	}
	//
	chars, _ := digit_map[digits[start]]
	for i := 0; i < len(chars); i++ {
		backtrace(digits, start+1, combine+string(chars[i]))
	}
}

// @lc code=end

