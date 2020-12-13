/*
 * @lc app=leetcode id=145 lang=cpp
 *
 * [145] Binary Tree Postorder Traversal
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
    std::vector<int> postorderTraversal(TreeNode* root) {
        if (root == nullptr)
            return postorder;
        //
        postorderTraversal(root->left);
        postorderTraversal(root->right);
        postorder.push_back(root->val);
        //
        return postorder;
    }
private:
    std::vector<int> postorder;
};
// @lc code=end

