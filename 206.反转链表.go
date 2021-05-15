/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// special case
	if head == nil || head.Next == nil {
		return head
	}
	//
	pre, cur, next := head, head.Next, head.Next.Next
	pre.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	//
	return pre
}

// @lc code=end

