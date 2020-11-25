/*
 * @lc app=leetcode id=108 lang=cpp
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
    TreeNode* sortedArrayToBSTImpl(vector<int>& nums, int left, int right)
    {
        // special case, null node
        if (left > right)
        {
            return nullptr;
        }
        // valid node
        // 注意此处 : left + (right - left + 1) / 2, mid优先右侧
        //          : left + (right - left) / 2, mid优先左侧
        int mid = left + (right - left + 1) / 2;
        TreeNode* root = new TreeNode(nums[mid]);
        root->left = sortedArrayToBSTImpl(nums, left, mid -1);
        root->right = sortedArrayToBSTImpl(nums, mid+1, right);
        return root;
    }
public:
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        // binary search
        return sortedArrayToBSTImpl(nums, 0, nums.size() -1);
    }
};
// @lc code=end

