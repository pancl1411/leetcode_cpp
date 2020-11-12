/*
 * @lc app=leetcode id=100 lang=cpp
 *
 * [100] Same Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
private:
    bool is_already_flase;
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        // return in advance
        if (is_already_flase)
        {
            return false;
        }
        // special case
        if (p == nullptr && q == nullptr)
        {
            return true;
        }
        else if (p == nullptr || q == nullptr)
        {
            is_already_flase = true;
            return false;
        }
        // compare node val
        if (p->val != q->val)
        {
            is_already_flase = true;
            return false;
        }
        // recursion with left and right node
        bool left = isSameTree(p->left, q->left);
        bool right = isSameTree(p->right, q->right);
        return left && right;
    }
};
// @lc code=end

