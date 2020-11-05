/*
 * @lc app=leetcode id=35 lang=cpp
 *
 * [35] Search Insert Position
 */
#include <vector>
// @lc code=start
class Solution {
public:
    // 二分法
    int searchInsert(std::vector<int>& nums, int target) {
        // special case
        if (nums.empty() || nums[0] >= target)
        {
            return 0;
        }
        else if (nums[nums.size() - 1] <= target)
        {
            return nums[nums.size() - 1] == target ? nums.size() - 1 : nums.size();
        }
        // 二分法
        unsigned int left, right, mid;
        left = 0;
        right = nums.size() -1;
        while (left <= right)
        {
            mid = (left + right) / 2;
            if (target == nums[mid])
            {
                return mid;
            }
            else if (target < nums[mid])
            {
                right = mid - 1;
            }
            else
            {
                left = mid + 1;
            }

        }
        return left;
    }
};
// @lc code=end

