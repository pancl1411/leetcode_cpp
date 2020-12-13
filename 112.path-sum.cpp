/*
 * @lc app=leetcode id=112 lang=cpp
 *
 * [112] Path Sum
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
    bool hasPathSum(TreeNode* root, int sum) {
        // special case : empty nodde
        if (root == nullptr)
        {
            return false;
        }
        // find the path success (loop to leaf node meet condition)
        //      NOTE : A leaf is a node with no children.
        if (sum == root->val && root->left == nullptr && root->right == nullptr)
        {
            return true;
        }
        // recursive with child tree
        bool left = hasPathSum(root->left, sum - root->val);
        bool right = hasPathSum(root->right, sum - root->val);
        return left || right;
    }
};
// @lc code=end

