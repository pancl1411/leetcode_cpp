/*
 * @lc app=leetcode id=83 lang=cpp
 *
 * [83] Rehead Duplicates from Sorted List
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
    // Straight-Forward Approach
    ListNode* deleteDuplicates(ListNode* head) {
        // special case
        if (head == nullptr)
        {
            return nullptr;
        }
        //
        ListNode dump_node = ListNode(0, head);
        while (head->next != nullptr)
        {
            if (head->val == head->next->val)
            {
                head->next = head->next->next;
            }
            else
            {
                head = head->next;
            }
        }
        return dump_node.next;
    }
};
// @lc code=end

