/*
 * @lc app=leetcode id=111 lang=cpp
 *
 * [111] Minimum Depth of Binary Tree
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
#include <algorithm>
class Solution {
public:
    int minDepth(TreeNode* root) {
        // special case : empty tree
        if (root == nullptr) return 0;
        // DFS
        int left = minDepth(root->left);
        int right = minDepth(root->right);
        // TWO CASE:
        //  1) left/right tree has/both empty, empty one will return 0, "left + right + 1"
        //  2) left/right tree both not empty, both return positive depth, "std::min() + 1" 
        return (left == 0 || right == 0) ? left + right + 1 : std::min(left, right)+1;
    }
};
// @lc code=end

