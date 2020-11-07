/*
 * @lc app=leetcode id=4 lang=cpp
 *
 * [4] Median of Two Sorted Arrays
 */
#include <vector>
#include <algorithm>
// @lc code=start
class Solution {
public:
    double findMedianSortedArrays(std::vector<int>& nums1, std::vector<int>& nums2) {
        //
        std::vector<int>* shorter= nums1.size() > nums2.size() ? &nums2 : &nums1; 
        std::vector<int>* longer = nums1.size() > nums2.size() ? &nums1 : &nums2;
        int short_len = shorter->size();
        int long_len = longer->size();
        int mid = (short_len + long_len + 1) / 2;
        int is_prime = (short_len + long_len) % 2; //奇数直接返回，偶数需要取2中间值均值
        // special case
        if (longer->empty())
        {
            return 0.0;
        }
        else if (shorter->empty())
        {
            return is_prime == 1 ? (*longer)[mid-1] : ((*longer)[mid-1] + (*longer)[mid]) / 2.0;
        }
        //
        int left = 0, right = short_len;
        while (left <= right)
        {
            int short_index = (left + right) / 2;
            int long_index = mid - short_index;
            if (short_index < short_len && (*shorter)[short_index] < (*longer)[long_index - 1])
            {
                left = short_index + 1;
            }
            else if (short_index > 0 && (*longer)[long_index] < (*shorter)[short_index - 1])
            {
                right = short_index - 1;
            }
            else
            {
                //
                int max_left;
                if (short_index == 0)
                {
                    max_left = (*longer)[long_index - 1];
                }
                else if (long_index == 0)
                {
                    max_left = (*shorter)[short_index - 1];
                }
                else
                {
                    max_left = std::max((*shorter)[short_index - 1], (*longer)[long_index - 1]);
                }
                //
                if (is_prime == 1)
                {
                    return max_left;
                }
                //
                int min_right;
                if (short_index == short_len)
                {
                    min_right = (*longer)[long_index];
                }
                else if (long_index == long_len)
                {
                    min_right = (*shorter)[short_index];
                }
                else
                {
                    min_right = std::min((*shorter)[short_index], (*longer)[long_index]);
                }
                return (max_left + min_right) / 2.0;
            }
        }
        //
        return 0.0;
    }
};
// @lc code=end

