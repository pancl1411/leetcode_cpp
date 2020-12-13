/*
 * @lc app=leetcode id=106 lang=cpp
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
    TreeNode* buildTreeImpl(int postIndex, int inBeginIndex, int inEndIndex,
                            vector<int>& inorder, vector<int>& postorder)
    {
        //
        if (postIndex < 0 || inBeginIndex > inEndIndex)
            return nullptr;
        // root value
        TreeNode* root = new TreeNode(postorder[postIndex]);
        // left/right tree
        int rootIndex;
        for (rootIndex = inBeginIndex; rootIndex <= inEndIndex; rootIndex++)
        {
            if (postorder[postIndex] == inorder[rootIndex])
                break;
        }
        root->right = buildTreeImpl(postIndex-1, rootIndex+1, inEndIndex, inorder, postorder);
        root->left = buildTreeImpl(postIndex-(inEndIndex-rootIndex)-1, inBeginIndex, rootIndex-1, inorder, postorder);
        return root;
    }
public:
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        return buildTreeImpl(postorder.size()-1, 0, inorder.size()-1, inorder, postorder);
    }
};
// @lc code=end

