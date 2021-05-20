/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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

func rob(root *TreeNode) int {
	answer = 0
	robDFS(root)
	return answer
}

func robDFS(root *TreeNode) (withRoot, noRoot int) {
	//
	if root == nil {
		return 0, 0
	}
	//
	leftWith, leftNo := robDFS(root.Left)
	rightWith, rightNo := robDFS(root.Right)
	//
	withRoot = root.Val + leftNo + rightNo
	noRoot = max(leftWith, leftNo) + max(rightWith, rightNo)
	//
	if withRoot > answer {
		answer = withRoot
	}
	if noRoot > answer {
		answer = noRoot
	}

	return
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

