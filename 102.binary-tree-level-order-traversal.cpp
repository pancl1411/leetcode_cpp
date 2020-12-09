/*
 * @lc app=leetcode id=102 lang=cpp
 *
 * [102] Binary Tree Level Order Traversal
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
private:
    // DFS 实现
    void levelOrderDFS(TreeNode* root, int depth, std::vector<std::vector<int>>& answer)
    {
        //
        if (root == nullptr) return;
        //
        levelOrderDFS(root->left, depth + 1, answer);
        //
        while (answer.size() < depth + 1)
        {
            answer.push_back(std::vector<int>());
        }
        answer[depth].push_back(root->val);
        //
        levelOrderDFS(root->right, depth + 1, answer);
    }
public:
    std::vector<std::vector<int>> levelOrder(TreeNode* root) {
        std::vector<std::vector<int>> answer;
        levelOrderDFS(root, 0, answer);
        return answer;
    }
};
// @lc code=end

