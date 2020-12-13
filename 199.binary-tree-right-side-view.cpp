/*
 * @lc app=leetcode id=199 lang=cpp
 *
 * [199] Binary Tree Right Side View
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
    std::vector<int> rightSideView(TreeNode* root) {
        buildRightView(root, 0);
        return _rightViewVal;
    }
private:
    std::vector<int> _rightViewVal;
    void buildRightView(TreeNode* root, int level)
    {
        if (root == nullptr) return;
        //
        if (_rightViewVal.size() == level)
        {
            _rightViewVal.push_back(root->val);
        }
        //
        buildRightView(root->right, level+1);
        buildRightView(root->left, level+1);
    }
};
// @lc code=end

