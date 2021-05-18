/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	// special case
	if len(s) < len(p) {
		return nil
	}
	answer := make([]int, 0)
	//
	target, current := make(map[byte]int, 0), make(map[byte]int, 0)
	for i, _ := range p {
		target[p[i]]++
	}
	// [start, end]
	start, end := 0, 0
	// loop
	for end < len(s) {
		// 右边界拓展
		if _, ok := target[s[end]]; ok {
			current[s[end]]++
		}
		end++
		// 左边界收缩
		if end > len(p) {
			if _, ok := current[s[start]]; ok {
				current[s[start]]--
			}
			start++
		}
		//
		if judgeEqual(target, current) {
			answer = append(answer, start)
		}
	}
	//
	return answer
}

func judgeEqual(f, s map[byte]int) bool {
	for k1, v1 := range f {
		if _, ok := s[k1]; !ok {
			return false
		}
		if v1 != s[k1] {
			return false
		}
	}
	return true
}

// @lc code=end

