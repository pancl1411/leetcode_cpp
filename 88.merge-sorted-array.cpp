/*
 * @lc app=leetcode id=88 lang=cpp
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
class Solution {
public:
    // 关键思想：倒序处理
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        while (n > 0)
        {
            if (m <= 0)
            {
                nums1[n-1] = nums2[n-1];
                n--;
            }
            else if (nums1[m-1] < nums2[n-1])
            {
                nums1[m + n - 1] = nums2[n-1];
                n--;
            }
            else
            {
                nums1[m + n - 1] = nums1[m-1];
                m--;
            }
        }
    }
};
// @lc code=end

