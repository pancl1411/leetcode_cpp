/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 转换为2个有序链表合并问题
// min办法反而更慢？
func mergeKLists(lists []*ListNode) *ListNode {
	/*
		// result init to nil for special case len(lists) == 0
		var result *ListNode = nil
		for i := 0; i < len(lists); i++ {
			result = mergeTwoList(result, lists[i])
		}
		return result
	*/
	// special case
	if len(lists) == 0 {
		return nil
	}
	//
	dump := &ListNode{Val: 0, Next: nil}
	tail := dump
	for tail != nil {
		tail.Next = getMinOne(lists)
		tail = tail.Next
	}
	return dump.Next
}

func mergeTwoList(first, second *ListNode) *ListNode {
	//
	if first == nil {
		return second
	} else if second == nil {
		return first
	}
	//
	if first.Val < second.Val {
		first.Next = mergeTwoList(first.Next, second)
		return first
	} else {
		second.Next = mergeTwoList(first, second.Next)
		return second
	}
}

func getMinOne(lists []*ListNode) *ListNode {
	//
	var idx, minVal int
	var isInit bool = false
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil && (!isInit || lists[i].Val < minVal) {
			minVal = lists[i].Val
			idx = i
			isInit = true
		}
	}
	//
	if lists[idx] == nil {
		return nil
	}
	//
	node := lists[idx]
	lists[idx] = node.Next
	return node
}

// @lc code=end

