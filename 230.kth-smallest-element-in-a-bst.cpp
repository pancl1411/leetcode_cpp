/*
 * @lc app=leetcode id=230 lang=cpp
 *
 * [230] Kth Smallest Element in a BST
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
public:
    int kthSmallest(TreeNode* root, int k) {
        if (root == nullptr)
        {
            // TODO: 此处可能不合适
            return 0;
        }
        if (node_cout >= k)
        {
            return kth_val;
        }
        // left
        kthSmallest(root->left, k);
        // root  (root->val need count as 1)
        node_cout++;
        if (node_cout == k)
        {
            kth_val = root->val;
        }
        // right
        kthSmallest(root->right, k);
        return kth_val;
    }
private:
    int kth_val;
    int node_cout;
};
// @lc code=end

