/*
 * @lc app=leetcode id=124 lang=cpp
 *
 * [124] Binary Tree Maximum Path Sum
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
#include <limits.h>
class Solution {
public:
    int maxPathSum(TreeNode* root) {
        _max_sum = INT_MIN;
        maxSumSubtree(root);
        return _max_sum;
    }
private:
    int _max_sum;
    int maxSumSubtree(TreeNode* root)
    {
        if (root == nullptr)
            return 0;
        // calculate left/right max sum, *set to zero* if sum is negative;
        int left_max = std::max(0, maxSumSubtree(root->left));
        int right_max = std::max(0, maxSumSubtree(root->right));
        // update current max node-to-node-path-sum
        _max_sum = std::max(_max_sum, left_max + right_max + root->val);
        //
        return std::max(left_max + root->val, right_max + root->val);
    }
};
// @lc code=end

