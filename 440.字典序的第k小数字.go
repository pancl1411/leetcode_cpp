/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] 字典序的第K小数字
 */

// @lc code=start
// 十叉树思路：https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/solution/ju-ran-shuang-100chu-hu-yi-liao-can-kao-liao-bie-r/
func findKthNumber(n int, k int) int {
	node, count := 1, 1
	for count != k {
		cur := getPreNodeCount(node, n)
		if cur+count > k {
			node *= 10
			count++
		} else {
			node++
			count += cur
		}
	}
	return node
}

// 以当前节点下，满足<=n的节点个数（含自己）
func getPreNodeCount(preNode, maxVal int) int {
	nextNode := preNode + 1
	var nodeCount int = 0
	for preNode <= maxVal {
		nodeCount += min(nextNode, maxVal+1) - preNode
		if preNode > maxVal/10 {
			break
		}
		preNode *= 10
		nextNode *= 10
	}
	return nodeCount
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

