/*
 * @lc app=leetcode id=99 lang=cpp
 *
 * [99] Recover Binary Search Tree
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
class Solution {
private:
    TreeNode* first;
    TreeNode* second;
    TreeNode* previous;
    void findTwoSwapNode(TreeNode* root)
    {
        // special case
        if (root == nullptr) return;
        // left child tree
        findTwoSwapNode(root->left);
        // in order traversal and operation
        if (previous != nullptr && previous->val >= root->val)
        {   
            // first只记录一次,[6,3]和[5,2]只记录6, 5不记录
            if (first == nullptr) first = previous;
            // second可能记录2次:
            //  1. [3,2]这种相邻的swap, 只记录了一次，并且与first同时记录
            //  2. [6,3]和[5,2]不是相邻的swap, 先记录了3, 后记录了2
            second = root;
        }
        previous = root;
        // right child tree
        findTwoSwapNode(root->right);
    }
public:
    void recoverTree(TreeNode* root) {
        // 题解：BST中序遍历的输出是递增顺序输出（in order traversal）, 题目替换了2个元素
        //  (1)相邻元素swap, 如：1, 3, 2, 4, 5, 6, 则[3, 2]不满足递增,first和second交换即可。
        //  (2)非相邻元素swap,如：1, 6, 3, 4, 5, 2, 则[6,3]和[5,2]不满足递增, 第一组的first/第二组的second交换即可。
        
        //
        first = nullptr;
        second = nullptr;
        previous = nullptr;
        findTwoSwapNode(root);
        //
        int swap_val = first->val;
        first->val = second->val;
        second->val = swap_val;
    }
};
// @lc code=end

