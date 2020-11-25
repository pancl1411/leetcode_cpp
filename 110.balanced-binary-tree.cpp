/*
 * @lc app=leetcode id=110 lang=cpp
 *
 * [110] Balanced Binary Tree
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
    /*
    * Return :
    *   -1  : not height-balanced tree
    *   >=0 : height of this tree
    */
    int balancedImpl(TreeNode* root)
    {
        // special case
        if (root == nullptr) return 0;
        // DFS
            // left
        int left_depth = balancedImpl(root->left);
        if(left_depth == -1) return -1;
            // right
        int right_depth = balancedImpl(root->right);
        if(right_depth == -1) return -1;
        // differ in height
        if (std::abs(left_depth-right_depth) > 1) return -1;
        // is balance and return height
        return std::max(left_depth, right_depth) + 1;
    }
public:
    bool isBalanced(TreeNode* root) {
        // DFS  
        return balancedImpl(root) == -1 ? false : true;
    }
};
// @lc code=end

