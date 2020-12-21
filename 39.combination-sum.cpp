/*
 * @lc app=leetcode id=39 lang=cpp
 *
 * [39] Combination Sum
 */

// @lc code=start
#include <vector>
class Solution {
public:
    // 问题关键在去重
    std::vector<std::vector<int>> combinationSum(std::vector<int>& candidates, int target) {
        std::vector<int> selected;
        backtrace(candidates, 0, target, selected);
        return answer;
    }
private:
    std::vector<std::vector<int>> answer;
    void backtrace(std::vector<int>& candidates, int candidates_index, int remain, std::vector<int>& selected)
    {
        //
        if (remain == 0)
        {
            answer.push_back(selected);
        }
        // 注意此处i的起点不是从0开始，因为答案需要unique
        for (int i = candidates_index; i < candidates.size(); i++)
        {
            if (candidates[i] <= remain)
            {
                selected.push_back(candidates[i]);
                // 注意此处i未加一，因为可以复用多次
                backtrace(candidates, i, remain-candidates[i], selected);
                selected.pop_back();
            }
        }
    }
};
// @lc code=end

