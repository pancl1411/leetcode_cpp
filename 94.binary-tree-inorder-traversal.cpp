/*
 * @lc app=leetcode id=94 lang=cpp
 *
 * [94] Binary Tree Inorder Traversal
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
#include <stack>
class Solution {
public:
    //（1）中序遍历递归实现
    //
    void inorderTraversalImpl(TreeNode* root, std::vector<int>& answer)
    {
        // special case
        if(root == nullptr) return;
        // inorder traversal
        inorderTraversalImpl(root->left, answer);
        answer.push_back(root->val);
        inorderTraversalImpl(root->right, answer);
    }
    
    std::vector<int> inorderTraversal(TreeNode* root) {
        std::vector<int> answer;
        inorderTraversalImpl(root, answer);
        return answer;
    }
    
   /*
   // (2)interative method
   std::vector<int> inorderTraversal(TreeNode* root) {
       std::vector<int> answer;
       std::stack<TreeNode *> node_stack;
       TreeNode* current = root;
       //
       while (current != nullptr || !node_stack.empty())
       {
           //
           while (current != nullptr)
           {
               node_stack.push(current);
               current = current->left;
           }
           //
           current = node_stack.top();
           node_stack.pop();
           answer.push_back(current->val);
           //
           current = current->right;
       }
       //
       return answer;
   }
   */

};
// @lc code=end

