/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哑节点思路trick
	dumy := &ListNode{0, head}
	fast, slow := head, dumy
	// move fast
	for n > 0 {
		n--
		fast = fast.Next
	}
	// move fast and slow
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// delete and return
	slow.Next = slow.Next.Next
	return dumy.Next
}

// @lc code=end

