/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// find mid
	dump := &ListNode{0, head}
	fast, slow := dump, dump
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//
	second := slow.Next
	slow.Next = nil
	//
	first := dump.Next
	second = reverse(second)
	merge(first, second)
}

func reverse(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}

func merge(first, second *ListNode) *ListNode {
	//
	if second == nil {
		return first
	}
	//
	fnext, snext := first.Next, second.Next
	first.Next = second
	second.Next = fnext
	fnext = merge(fnext, snext)
	return first
}

// @lc code=end

