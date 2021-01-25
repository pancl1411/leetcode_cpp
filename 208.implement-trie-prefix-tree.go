/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	// use array not slice will speed up
	child [26]*Trie
	// word end flag, true if this node is end of a word ("end-word-node")
	end bool
}

func NewTrie() *Trie {
	return &Trie{}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, val := range word {
		if nil == node.child[val-'a'] {
			node.child[val-'a'] = NewTrie()
		}
		node = node.child[val-'a']
	}
	// important, set word end flag, named "end-word-node"
	node.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, val := range word {
		if nil == node.child[val-'a'] {
			return false
		}
		node = node.child[val-'a']
	}
	// false if not the "end-word-node"
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, val := range prefix {
		if nil == node.child[val-'a'] {
			return false
		}
		node = node.child[val-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

