/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// special case
	if head == nil || head.Next == nil {
		return true
	}
	// 查找回文的终点: n/2 向下取整
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 后半段链表反转
	first, second := head, reverseList(slow.Next)
	// 逐个节点判断值是否相等
	// 	second可能少一个节点,所以second循环结束即可
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode = nil, head, nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end

