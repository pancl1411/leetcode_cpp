/*
 * @lc app=leetcode id=47 lang=cpp
 *
 * [47] Permutations 
 * II
 */

// @lc code=start
#include <vector>
#include <algorithm>
class Solution {
public:
    std::vector<std::vector<int>> permuteUnique(std::vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        is_have = std::vector<bool>(nums.size(), false);
        std::vector<int> selected;
        backtraceImpl(nums, selected);
        return answer;
    }
private:
    std::vector<std::vector<int>> answer;
    std::vector<bool> is_have;
    void backtraceImpl(std::vector<int>& nums, std::vector<int>& selected)
    {
        if (selected.size() == nums.size())
        {
            answer.push_back(selected);
            return;
        }
        
        for (int index = 0; index < nums.size(); index++)
        {
            // 用过的
            if (is_have[index]) continue;
            // 关键：排序后去除重复答案（相同元素，只有排在前面的使用过，才会继续用。
            //          保证一个答案不会由不同位置相同元素构造多次）
            if (index > 0 && nums[index-1] == nums[index] && !is_have[index-1])
                continue;
            //
            selected.push_back(nums[index]);
            is_have[index] = true;
            backtraceImpl(nums, selected);
            selected.pop_back();
            is_have[index] = false;
        }
        
    }
};
// @lc code=end

