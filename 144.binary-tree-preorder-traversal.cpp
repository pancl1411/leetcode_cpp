/*
 * @lc app=leetcode id=144 lang=cpp
 *
 * [144] Binary Tree Preorder Traversal
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
#include <vector>
class Solution {
public:
    std::vector<int> preorderTraversal(TreeNode* root) {
        if (root == nullptr) 
            return preorder;
        //
        preorder.push_back(root->val);
        preorderTraversal(root->left);
        preorderTraversal(root->right);
        //
        return preorder;
    }
private:
    std::vector<int> preorder;
};
// @lc code=end

