/*
 * @lc app=leetcode id=129 lang=cpp
 *
 * [129] Sum Root to Leaf Numbers
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
    int sumNumbers(TreeNode* root) {
        _total_sum = 0;
        sumToLeaf(root, 0);
        return _total_sum;
    }
private:
    int _total_sum;
    void sumToLeaf(TreeNode* root, int sum)
    {
        // special case
        if (root == nullptr) return;
        // up to leaf node, add it
        sum = sum*10 + root->val;
        if (root->left == nullptr && root->right == nullptr)
            _total_sum += sum;
        // recursively
        sumToLeaf(root->left, sum);
        sumToLeaf(root->right, sum);
    }
};
// @lc code=end

