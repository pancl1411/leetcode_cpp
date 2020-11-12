/*
 * @lc app=leetcode id=101 lang=cpp
 *
 * [101] Symmetric Tree
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
    bool is_already_false;
    bool isSymmetricTwo(TreeNode* left, TreeNode* right)
    {
        // return in advance
        if (is_already_false)
        {
            return false;
        }
        // special case
        if (left == nullptr && right == nullptr)
        {
            return true;
        }
        else if (left == nullptr || right == nullptr)
        {
            is_already_false = true;
            return false;
        }
        // compare val between left and right
        if (left->val != right->val)
        {
            is_already_false = true;
            return false;
        }
        // recursive
        bool outter = isSymmetricTwo(left->left, right->right);
        bool inner = isSymmetricTwo(left->right, right->left);
        if (!(outter && inner))
        {
            is_already_false = true;
            return false;
        }
        return true;
    }
public:
    bool isSymmetric(TreeNode* root) {
        if (root == nullptr)
        {
            return true;
        }
        return isSymmetricTwo(root->left, root->right);
    }
};
// @lc code=end

