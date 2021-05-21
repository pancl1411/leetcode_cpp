/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
var result int

func pathSum(root *TreeNode, targetSum int) int {
	result = 0
	has := make(map[int]int, 0)
	has[0] = 1 // 注意：这个初始化很重要
	backtrace(root, targetSum, has, 0)
	return result
}

// 前缀和思路
func backtrace(root *TreeNode, targetSum int, has map[int]int, totalSum int) {
	//
	if root == nil {
		return
	}
	//
	totalSum += root.Val
	result += has[totalSum-targetSum]
	//
	has[totalSum]++
	backtrace(root.Left, targetSum, has, totalSum)
	backtrace(root.Right, targetSum, has, totalSum)
	has[totalSum]--
}

// @lc code=end

