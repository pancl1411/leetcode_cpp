/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//
	if root == nil {
		return nil
	}
	// 找到一个节点，其子节点可以不用找了
	// （如果在子节点中,那这个节点就是最低的公共祖先）
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//
	if left != nil && right != nil {
		// 左右各一个，是最低的公共祖先
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

// @lc code=end

