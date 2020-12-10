/*
 * @lc app=leetcode id=103 lang=cpp
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
#include <list>
class Solution
{
public:
    // BFS with flag to read from front/back
    std::vector<std::vector<int>> zigzagLevelOrder(TreeNode *root)
    {
        std::vector<std::vector<int>> answer;
        if (root == nullptr)
            return answer;
        // init
        std::list<TreeNode *> level_nodes;
        level_nodes.push_back(root);
        bool is_left_to_right = true;
        // BFS with flag
        while (!level_nodes.empty())
        {
            std::vector<int> level_values;
            TreeNode *node = nullptr;
            int level_size = level_nodes.size();

            for (int i = 0; i < level_size; i++)
            {
                // left to right, 
                if (is_left_to_right)
                {
                    node = level_nodes.front();
                    level_nodes.pop_front();
                    if (node->left)
                        level_nodes.push_back(node->left);
                    if (node->right)
                        level_nodes.push_back(node->right);
                }
                // right to left
                else
                {
                    node = level_nodes.back();
                    level_nodes.pop_back();
                    if (node->right)
                        level_nodes.push_front(node->right);
                    if (node->left)
                        level_nodes.push_front(node->left);
                }
                //
                level_values.push_back(node->val);
            }
            // record this level values, then reverse flag
            answer.push_back(level_values);
            is_left_to_right = !is_left_to_right;
        }
        return answer;
    }
};
// @lc code=end
