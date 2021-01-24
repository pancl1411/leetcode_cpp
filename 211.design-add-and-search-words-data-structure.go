/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start
type WordDictionary struct {
	head *trieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	dict := WordDictionary{}
	dict.head = NewtrieNode()
	return dict
}

func (this *WordDictionary) AddWord(word string) {
	node := this.head
	for _, letter := range word {
		if node.child[letter-'a'] == nil {
			node.child[letter-'a'] = NewtrieNode()
		}
		node = node.child[letter-'a']
	}
	// important, sign last node is end of a word
	node.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.head.search(word)
}

// trieNode to record word's letter tree
type trieNode struct {
	// child[i] stands next char('a'+i)
	child []*trieNode
	// true if this node is the end of one/multi word
	end bool
}

// NewtrieNode to get a pointer to trieNode
func NewtrieNode() *trieNode {
	node := trieNode{}
	// only need 26 lower-case English letters
	node.child = make([]*trieNode, 26)
	node.end = false
	return &node
}

// search method with DFS
func (this *trieNode) search(word string) bool {
	// DFS finish, true if precise end-node
	if len(word) == 0 {
		return this.end
	}
	// CASE '.' : DFS all child
	if word[0] == '.' {
		for _, node := range this.child {
			if node != nil && node.search(word[1:]) {
				return true
			}
		}
		// CASE others : DFS if exist
	} else {
		return (this.child[word[0]-'a'] != nil) && this.child[word[0]-'a'].search(word[1:])
	}
	//
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

