/*
 * @lc app=leetcode id=105 lang=cpp
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
    //  preIndex : preorder traversal index 
    //    (preorder[preIndex] will be the root value of subtree, and spit inorder[] left/right)
    //  inBeginIndex : inorder traversal begin index
    //  inEndIndex ： inorder traversal end index
    TreeNode* buildTreeImpl(int preIndex, int inBeginIndex, int inEndIndex,
                            vector<int>& preorder, vector<int>& inorder)
    {
        // 
        if (preIndex >= preorder.size() || inBeginIndex > inEndIndex)
        {
            return nullptr;
        }
        // root value
        TreeNode* root = new TreeNode(preorder[preIndex]);
        // left/right child tree
        int rootIndex = 0;
        for (rootIndex = inBeginIndex; rootIndex <= inEndIndex; rootIndex++)
        {
            if (preorder[preIndex] == inorder[rootIndex])
            {
                break;
            }
        }
        root->left = buildTreeImpl(preIndex+1, inBeginIndex, rootIndex-1, preorder, inorder);
        root->right = buildTreeImpl(preIndex+1+(rootIndex-inBeginIndex), rootIndex+1, inEndIndex, preorder, inorder);
        //
        return root;
    }
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return buildTreeImpl(0, 0, inorder.size()-1, preorder, inorder);
    }
};
// @lc code=end

