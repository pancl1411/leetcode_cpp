/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
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
    // 注意直接复用了输入的node, 若可以新建node，则可以简化代码
    ListNode* add(ListNode* l1, ListNode* l2, int carry)
    {
        // special case with null node
        if (l1 == nullptr && l2 == nullptr)
        {
            if (carry == 0)
            {
                return nullptr;
            }
            else
            {
                ListNode* new_node = new ListNode(1, nullptr);
                return new_node;
            }
        }
        else if (l1 == nullptr && l2 != nullptr)
        {
            if (l2->val + carry < 10)
            {
                l2->val += carry;
            }
            else
            {
                l2->val = 0;
                l2->next = add(nullptr, l2->next, 1);
            }
            return l2;
        }
        else if (l2 == nullptr && l1 != nullptr)
        {
            if (l1->val + carry < 10)
            {
                l1->val += carry;
            }
            else
            {
                l1->val = 0;
                l1->next = add(nullptr, l1->next, 1);
            }
            return l1;
        }
        // cursively
        if (l1->val + l2->val + carry < 10)
        {
            l1->val = l1->val + l2->val + carry;
            l1->next = add(l1->next, l2->next, 0);
        }
        else
        {
            l1->val = l1->val + l2->val + carry - 10;
            l1->next = add(l1->next, l2->next, 1);
        }
        return l1;
    }

    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        //
        if (l1 == nullptr)
        {
            return l2;
        }
        else if (l2 == nullptr)
        {
            return l1;
        }
        
        //
        return add(l1, l2, 0);
    }
};
// @lc code=end

