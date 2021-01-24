/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	// cache in map like: ' word1 word2 ...'
	wordCache := make(map[string][]string, 0)
	result, _ := backWithBuffer(s, wordDict, wordCache)
	// 删除头部的空格
	answer := make([]string, 0)
	for _, one := range result {
		answer = append(answer, one[1:])
	}
	//
	return answer
}

func backWithBuffer(s string, wordDict []string, wordCache map[string][]string) (answer []string, isSolved bool) {
	// 命中cache map
	if value, exist := wordCache[s]; exist {
		if value == nil {
			return nil, false
		} else {
			return value, true
		}
	}
	// DFS查找结束
	if len(s) == 0 {
		return nil, true
	}
	// DFS查找: trick 从dict里反向取值
	for _, word := range wordDict {
		if len(s) >= len(word) && word == s[:len(word)] {
			buffer, isSolved := backWithBuffer(s[len(word):], wordDict, wordCache)
			if isSolved {
				cache := make([]string, 0)
				if len(buffer) == 0 {
					cache = append(cache, " "+word)
				} else {
					for _, str := range buffer {
						cache = append(cache, " "+word+str)
					}
				}
				// cache结果，此处用append而不是赋值，可能因为不同word新增答案
				wordCache[s] = append(wordCache[s], cache...)
			}
		}
	}
	// cache无解字符串，避免重复
	answer, ok := wordCache[s]
	if !ok {
		wordCache[s] = nil
	}
	return answer, ok
}

// @lc code=end

