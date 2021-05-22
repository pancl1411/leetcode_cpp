/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
var answer int

func diameterOfBinaryTree(root *TreeNode) int {
	answer = 0
	postorder(root)
	return answer
}

func postorder(root *TreeNode) (maxlen int) {
	if root == nil {
		return 0
	}
	//
	left := postorder(root.Left)
	right := postorder(root.Right)
	sum := left + 1 + right
	// sum是节点数，算成半径需要减1
	if sum-1 > answer {
		answer = sum - 1
	}
	return max(left, right) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

