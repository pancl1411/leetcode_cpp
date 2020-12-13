/*
 * @lc app=leetcode id=113 lang=cpp
 *
 * [113] Path Sum II
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
    std::vector<std::vector<int>> pathSum(TreeNode* root, int sum) {
        std::vector<std::vector<int>> paths;
        std::vector<int> one_path;
        findPaths(root, sum, paths, one_path);
        return paths;
    }
private:
    void findPaths(TreeNode* root, int sum, std::vector<std::vector<int>>& paths, std::vector<int>& one_path)
    {
        //
        if (root == nullptr) return;
        //
        one_path.push_back(root->val);
        // find one result and record it
        if (sum == root->val && root->left == nullptr && root->right == nullptr)
        {
            paths.push_back(one_path);
            // if you want return there, must do one_path.pop_back()
            //one_path.pop_back();
            //return;
        }
        // recursively with left/right tree
        findPaths(root->left, sum - root->val, paths, one_path);
        findPaths(root->right, sum - root->val, paths, one_path);
        // passing by reference and pop_back, to save memory
        one_path.pop_back();
    }
};
// @lc code=end

