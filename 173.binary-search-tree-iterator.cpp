/*
 * @lc app=leetcode id=173 lang=cpp
 *
 * [173] Binary Search Tree Iterator
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
#include <stack>
class BSTIterator {
public:
    BSTIterator(TreeNode* root) {
        push_all_left(root);
    }
    
    int next() {
        TreeNode* node = _node_stack.top();
        _node_stack.pop();
        // 此处O(1), 需要按n次平均理解
        push_all_left(node->right);
        return node->val;
    }
    
    bool hasNext() {
        return !_node_stack.empty();
    }
private:
    std::stack<TreeNode*> _node_stack;
    void push_all_left(TreeNode* root)
    {
        while (root != nullptr)
        {
            _node_stack.push(root);
            root = root->left;
        }
    }
};

/**
 * Your BSTIterator object will be instantiated and called as such:
 * BSTIterator* obj = new BSTIterator(root);
 * int param_1 = obj->next();
 * bool param_2 = obj->hasNext();
 */
// @lc code=end

