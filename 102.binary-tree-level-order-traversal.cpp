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
#include <queue>
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

    // BFS 实现
    void levelOrderBFS(TreeNode* root, std::vector<std::vector<int>>& answer)
    {
        // init
        std::queue<TreeNode*> bfs_queue;
        TreeNode* node = nullptr;
        bfs_queue.push(root);
        // BFS
        while (!bfs_queue.empty())
        {
            std::vector<int> level_val;
            int level_size = bfs_queue.size();  // 规避queue的size变化，导致循环次数错误
            for (int i = 0; i < level_size; i++)
            {
                //
                node = bfs_queue.front();
                bfs_queue.pop();
                //
                level_val.push_back(node->val);
                //
                if(node->left != nullptr) bfs_queue.push(node->left);
                if(node->right != nullptr) bfs_queue.push(node->right);
            }
            answer.push_back(level_val);
        }
    }
public:
    std::vector<std::vector<int>> levelOrder(TreeNode* root) {
        std::vector<std::vector<int>> answer;
        if (root == nullptr) return answer;
        
        //levelOrderDFS(root, 0, answer);
        levelOrderBFS(root, answer);
        return answer;
    }
};
// @lc code=end

