/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	spilt byte
	null  byte
	info  []string
}

func Constructor() Codec {
	return Codec{
		spilt: ',',
		null:  '#',
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := string(this.serializeImpl(root, []byte{}))
	//fmt.Println("code:", result)
	return result
}

func (this *Codec) serializeImpl(root *TreeNode, pre []byte) []byte {
	//
	if len(pre) != 0 {
		pre = append(pre, this.spilt)
	}
	//
	if root == nil {
		pre = append(pre, this.null)
	} else {
		pre = append(pre, []byte(strconv.Itoa(root.Val))...)
		pre = this.serializeImpl(root.Left, pre)
		pre = this.serializeImpl(root.Right, pre)
	}
	//
	return pre
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.info = strings.Split(data, string(this.spilt))
	// fmt.Println("decode:", this.info)
	return this.deserializeImpl()
}

func (this *Codec) deserializeImpl() *TreeNode {
	//
	if this.info[0] == string(this.null) {
		this.info = this.info[1:]
		return nil
	}
	//
	val, _ := strconv.Atoi(this.info[0])
	this.info = this.info[1:]
	//
	root := &TreeNode{
		Val:   val,
		Left:  this.deserializeImpl(),
		Right: this.deserializeImpl(),
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

