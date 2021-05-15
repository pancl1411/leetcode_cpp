/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	root *trieNode
}

type trieNode struct {
	dict [26]*trieNode
	end  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: createNode(),
	}
}

func createNode() *trieNode {
	var empty [26]*trieNode
	return &trieNode{
		dict: empty,
		end:  false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, v := range word {
		idx := int(v - 'a')
		if node.dict[idx] == nil {
			node.dict[idx] = createNode()
		}
		node = node.dict[idx]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, v := range word {
		idx := int(v - 'a')
		if node.dict[idx] == nil {
			return false
		}
		node = node.dict[idx]
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, v := range prefix {
		idx := int(v - 'a')
		if node.dict[idx] == nil {
			return false
		}
		node = node.dict[idx]
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

