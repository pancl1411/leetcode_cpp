/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//
	if headA == nil || headB == nil {
		return nil
	}
	// 相遇在交点，或者走完A+B同时到达nil
	moveA, moveB := headA, headB
	for moveA != moveB {
		// A
		if moveA == nil {
			moveA = headB
		} else {
			moveA = moveA.Next
		}
		// B
		if moveB == nil {
			moveB = headA
		} else {
			moveB = moveB.Next
		}
	}
	//
	return moveA
}

// @lc code=end

