/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	answer := make([][]string, 0)
	palindrome := make([]string, 0)
	backTrack(s, palindrome, &answer)
	return answer
}

func backTrack(s string, palindrome []string, answer *[][]string) {
	//fmt.Println(len(s), "\t", s, palindrome)
	if 0 >= len(s) {
		// 注意此处若不进行拷贝，可能值被后续迭代改变
		newOne := make([]string, len(palindrome))
		copy(newOne, palindrome)
		*answer = append(*answer, newOne)
		//fmt.Println(answer)
		return
	}

	for i := 1; i <= len(s); i++ {
		sub := s[:i]
		if isPalindrome(sub) {
			palindrome = append(palindrome, sub)
			backTrack(s[i:], palindrome, answer)
			palindrome = palindrome[:len(palindrome)-1]
		}

	}
}

func isPalindrome(str string) bool {
	if len(str) == 1 {
		return true
	}

	// i start with 0, len(str)-i-1
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

// @lc code=end

