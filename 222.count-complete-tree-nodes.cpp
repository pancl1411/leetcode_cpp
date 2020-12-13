/*
 * @lc app=leetcode id=222 lang=cpp
 *
 * [222] Count Complete Tree Nodes
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
#include <math.h>
class Solution {
public:
    int countNodes(TreeNode* root) {
        if (root == nullptr)
        {
            return 0;
        }
        // left depth
        TreeNode* left_node = root;
        int left_depth = 0;
        while (left_node != nullptr)
        {
            left_depth++;
            left_node = left_node->left;
        }
        // right depth
        TreeNode* right_node = root;
        int right_depth = 0;
        while (right_node != nullptr)
        {
            right_depth++;
            right_node = right_node->right;
        }
        // end condition
        if (left_depth == right_depth)
        {
            return pow(2, left_depth) - 1;
        }
        // recursively 
        return 1 + countNodes(root->left) + countNodes(root->right);
    }
};
// @lc code=end

