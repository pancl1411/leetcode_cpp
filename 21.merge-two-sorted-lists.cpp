/*
 * @lc app=leetcode id=21 lang=cpp
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
	// (1) Recursive Solutions
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        // special case
		if (l1 == nullptr)
		{
			return l2;
		}
		else if (l2 == nullptr)
		{
			return l1;
		}
		//
		if (l1->val < l2->val)
		{
			l1->next = mergeTwoLists(l1->next, l2);
			return l1;
		}
		else
		{
			l2->next = mergeTwoLists(l1, l2->next);
			return l2;
		}
		//
		return nullptr;
    }
	
	// (2) Iteratively
	/*
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        // special case
        if (l1 == nullptr)
        {
            return l2;
        }
        if (l2 == nullptr)
        {
            return l1;
        }
        // preserve head
        ListNode head = ListNode(0);
        // Iteratively
        ListNode* cur = &head;
        while (l1 && l2)
        {
            if (l1->val < l2->val)
            {
                cur->next = l1;
                l1 = l1->next;
            }
            else
            {
                cur->next = l2;
                l2 = l2->next;
            }
            cur = cur->next;
        }
        cur->next = (l1 == nullptr) ? l2 : l1;
        //
        return head.next;
    }*/
};
// @lc code=end

