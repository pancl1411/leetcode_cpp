/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start

// 没有做出来，放弃！！！
var answer [][]string

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	//
	wordMap := make(map[string]bool, len(wordList))
	for word, _ := range wordList {
		wordMap[word] = true
	}
	//
	ladder := make([]string, 0)
	backTrack(beginWord, endWord, wordMap, ladder)
	//
	return answer
}

func backTrack(curWord string, endWord string, wordMap map[string]bool, ladder []string) {
	//
	if curWord == endWord {
		append(answer, ladder)
		return
	}
	//
	for word, used := range wordMap {
		if isOneCharDiff([]byte(curWord), []byte(word)) {
			append(ladder, word)
			//wordMap[word] = false
			backTrack(word, endword, wordMap, ladder)
			ladder = ladder[0 : len(ladder)-1]
			//wordMap[word] = true
		}
	}
}

func isOneCharDiff(first, second []byte) bool {
	if len(first) != len(second) {
		return false
	}

	count := 0
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}

	if count != 1 {
		return false
	}
	return true
}

// @lc code=end

