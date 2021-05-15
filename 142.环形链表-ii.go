/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 注意这里初始化, 是都走了一步
	fast, slow := head.Next.Next, head.Next
	//
	for fast != nil && fast.Next != nil {
		if fast == slow {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	//
	return nil
}

// @lc code=end

