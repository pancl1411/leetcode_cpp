/*
 * @lc app=leetcode id=26 lang=cpp
 *
 * [26] Remove Duplicates from Sorted Array
 */
// @lc code=start
#include <vector>

class Solution {
public:
    int removeDuplicates(std::vector<int>& nums) {
        // special case 
        if (nums.size() <= 1)
        {
            return nums.size();
        }
        //
        int answer = 0;
        for (int index = 1; index < nums.size(); index++)
        {
            if (nums[answer] != nums[index])
            {
                answer++;
                nums[answer] = nums[index];
            }
        }
        // 从0计数, 当作长度返回时需要加一
        return answer + 1;
    }
};
// @lc code=end

