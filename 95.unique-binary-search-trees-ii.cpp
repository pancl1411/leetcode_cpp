/*
 * @lc app=leetcode id=95 lang=cpp
 *
 * [95] Unique Binary Search Trees II
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
    std::vector<TreeNode*> generateTreesImpl(int min, int max) {
        std::vector<TreeNode*> answer, left, right;
        if (min > max)
        {
            answer.push_back(nullptr);  //要点:确保边界有null节点进行循环
            return answer;
        }
        else if (min == max)
        {
            TreeNode* new_node = new TreeNode(min);
            answer.push_back(new_node);
            return answer;
        }
        
        for (int i = min; i <= max; i++)
        {
            //
            left = generateTreesImpl(min, i-1);
            right = generateTreesImpl(i+1, max);
            //
            for (auto left_iter : left)
            {
                for ( auto right_iter : right)
                {
                    // root节点在循环体内，避免被覆盖
                    TreeNode* root = new TreeNode(i); 
                    root->left = left_iter;
                    root->right = right_iter;
                    answer.push_back(root);
                }
            }
        }
        //
        return answer;
    }
public:
    std::vector<TreeNode*> generateTrees(int n) {        
        // special case
        if (n == 0)
        {
            return std::vector<TreeNode*>();
        }
        // Divide and Conquer
        return generateTreesImpl(1, n);
    }
};
// @lc code=end

