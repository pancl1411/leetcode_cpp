/*
 * @lc app=leetcode id=107 lang=cpp
 *
 * [107] Binary Tree Level Order Traversal II
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
    void BFS(std::vector<TreeNode*>& roots, std::vector<std::vector<int>>& answer)
    {
        if (roots.empty())
        {
            return;
        }
        //
        std::vector<TreeNode*> next_level;
        std::vector<int> this_level_val;
        for (size_t i = 0; i < roots.size(); i++)
        {
            this_level_val.push_back(roots[i]->val);
            if (roots[i]->left != nullptr)
            {
                next_level.push_back(roots[i]->left);
            }
            if (roots[i]->right != nullptr)
            {
                next_level.push_back(roots[i]->right);
            }
        }
        //
        if (!next_level.empty())
        {
            BFS(next_level, answer);
        }
        //
        answer.push_back(this_level_val);
    }
public:
    std::vector<std::vector<int>> levelOrderBottom(TreeNode* root) {
        std::vector<std::vector<int>> answer;
        // special case
        if (root == nullptr)
        {
            return answer;
        }
        // Breadth-first-search
        std::vector<TreeNode*> next_level;
        next_level.push_back(root);
        BFS(next_level, answer);
        return answer;
    }
};
// @lc code=end

