/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersImpl(l1, l2, 0)
}

func addTwoNumbersImpl(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		l1.Val = sum % 10
		l1.Next = addTwoNumbersImpl(l1.Next, l2.Next, carry)
		return l1
	} else if l1 != nil {
		sum := l1.Val + carry
		carry = sum / 10
		l1.Val = sum % 10
		l1.Next = addTwoNumbersImpl(l1.Next, l2, carry)
		return l1
	} else if l2 != nil {
		sum := l2.Val + carry
		carry = sum / 10
		l2.Val = sum % 10
		l2.Next = addTwoNumbersImpl(l1, l2.Next, carry)
		return l2
	} else {
		// l1/l2均不存在
		if carry != 0 {
			return &ListNode{Val: carry, Next: nil}
		}
	}

	return nil
}

// @lc code=end

