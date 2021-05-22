/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	//
	dump := &ListNode{0, head}
	// left[前一个]位置的节点
	leftNode, rightNode := dump, dump
	for left > 1 && leftNode != nil {
		left--
		leftNode = leftNode.Next
	}
	// right位置的节点
	for right > 0 && rightNode != nil {
		right--
		rightNode = rightNode.Next
	}
	// 切割+局部反转+拼接回来
	end, next := leftNode.Next, rightNode.Next
	rightNode.Next = nil
	leftNode.Next = reverse(leftNode.Next)
	end.Next = next
	//
	return dump.Next
}

//
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}

// @lc code=end

