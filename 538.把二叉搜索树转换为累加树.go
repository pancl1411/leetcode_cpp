/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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

var sum int // 需要用一个全局的变量记录，传递值返回会存在左右子数不好处理情况

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	return dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	//
	if root == nil {
		return nil
	}
	// 中序遍历（先right分支）
	dfs(root.Right)
	sum += root.Val
	root.Val = sum
	dfs(root.Left)
	return root
}

// @lc code=end

