/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	//
	if head == nil {
		return nil
	}
	//
	dumpNode := &ListNode{0, head}
	prev, end := dumpNode, dumpNode
	for {
		// get k
		for i := 0; i < k; i++ {
			if end == nil || end.Next == nil {
				return dumpNode.Next
			}
			end = end.Next
		}
		start := prev.Next
		next := end.Next
		end.Next = nil
		// [start, end] -> [end, start]
		prev.Next = reverse(start)
		//
		start.Next = next
		prev = start
		end = start
	}
	//
	return dumpNode.Next
}

func reverse(head *ListNode) *ListNode {
	//
	if head == nil {
		return nil
	}
	//
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end

