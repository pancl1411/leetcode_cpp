/*
 * @lc app=leetcode id=90 lang=cpp
 *
 * [90] Subsets II
 */

// @lc code=start
#include <vector>
#include <algorithm>
class Solution {
public:
    std::vector<std::vector<int>> subsetsWithDup(std::vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        std::vector<int> set;
        subsetsWithDupImpl(nums, 0, set);
        return _answer;
    }
private:
    std::vector<std::vector<int>> _answer;
    void subsetsWithDupImpl(std::vector<int>& nums, int start, std::vector<int>& set)
    {
        //
        _answer.push_back(set);
        //
        if (start >= nums.size())
        {
            return;
        }
        // 注意此处i的增长方式
        for (int i = start; i < nums.size(); i=getNextDiffIndex(nums, i))
        {
            set.push_back(nums[i]);
            // 此处i+1, 可以重复数字
            subsetsWithDupImpl(nums, i+1, set);
            set.pop_back();
        }
    }

    int getNextDiffIndex(std::vector<int>& nums, int now)
    {
        int next = now + 1;
        while (next < nums.size() && nums[next] == nums[now])
        {
            next++;
        }
        return next;
    }
};
// @lc code=end

