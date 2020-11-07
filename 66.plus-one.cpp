/*
 * @lc app=leetcode id=66 lang=cpp
 *
 * [66] Plus One
 */
#include <vector>
// @lc code=start
class Solution {
public:
    std::vector<int> plusOne(std::vector<int>& digits) {
        // 加一，有进位则进位，无进位后return
        for (int i = digits.size() - 1; i >= 0; i--)
        {
            if (digits[i] < 9)
            {
                digits[i]++;
                return digits;
            }
            // 进位
            digits[i] = 0;
        }
        // 全9情形，需要在数组头部加一位
        std::vector<int> max = std::vector<int> (digits.size() + 1, 0);
        max[0] = 1;
        return max;
    }
};
// @lc code=end

