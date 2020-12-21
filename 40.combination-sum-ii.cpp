/*
 * @lc app=leetcode id=40 lang=cpp
 *
 * [40] Combination Sum II
 */

// @lc code=start
#include <vector>
#include <algorithm>
class Solution {
public:
    // 问题点：candidates内有重复元素，答案如何去重
    std::vector<std::vector<int>> combinationSum2(std::vector<int>& candidates, int target) {
        std::vector<int> selected;
        // 排序后答案去重更方便
        std::sort(candidates.begin(), candidates.end());
        backtrace(candidates, 0, target, selected);
        return answer;
    }
private:
    std::vector<std::vector<int>> answer;
    void backtrace(std::vector<int>& candidates, int start, int remain, std::vector<int> selected)
    {
        if (remain == 0)
        {
            answer.push_back(selected);
            return;
        }
        for (int index = start; index < candidates.size(); index++)
        {
            // 答案去重（排序后）
            if (index > start && candidates[index] == candidates[index - 1])
                continue;
            // backtrace
            if (candidates[index] <= remain)
            {
                selected.push_back(candidates[index]);
                backtrace(candidates, index+1, remain-candidates[index], selected);
                selected.pop_back();
            }
        }
    }
};
// @lc code=end

