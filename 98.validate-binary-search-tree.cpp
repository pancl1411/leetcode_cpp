/*
 * @lc app=leetcode id=98 lang=cpp
 *
 * [98] Validate Binary Search Tree
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
public:
    int isBSTImpl(TreeNode* root, int& out_min, int& out_max)
    {
        // special case
        if (root == nullptr)
        {
            return 0;
        }
        // left tree
        int left_max, left_min;
        int left = isBSTImpl(root->left, left_min, left_max);
        if (left < 0) return -1;
        if (left > 0 && root->val <= left_max) return -1;
        // right tree
        int right_max, right_min;
        int right = isBSTImpl(root->right, right_min, right_max);
        if (right < 0) return -1;
        if (right > 0 && root->val >= right_min) return -1;
        // update tree max/min value, then return
        out_min = (left == 0) ? root->val : left_min;
        out_max = (right == 0) ? root->val : right_max; 
        return 1;
    }

    bool isValidBST(TreeNode* root) {
        int max, min;
        return isBSTImpl(root, min, max) < 0 ? false : true;
    }
};
// @lc code=end

