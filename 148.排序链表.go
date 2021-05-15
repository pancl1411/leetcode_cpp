/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return binarySortList(head, nil)
}

// [head, tail]
func binarySortList(head, tail *ListNode) *ListNode {
	// special case
	if head == nil || head == tail {
		return head
	}
	// find middle in [head, tail]
	fast, slow := head.Next, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 注意: 此处需要拆链表
	mid := slow.Next
	slow.Next = nil
	// divide and merge
	left := binarySortList(head, slow)
	right := binarySortList(mid, tail)
	return mergeTwoList(left, right)

}

func mergeTwoList(first, second *ListNode) *ListNode {
	// special case
	if first == nil {
		return second
	} else if second == nil {
		return first
	}
	// 递归合并
	if first.Val < second.Val {
		first.Next = mergeTwoList(first.Next, second)
		return first
	} else {
		second.Next = mergeTwoList(first, second.Next)
		return second
	}
}

// @lc code=end

