/*
 * @lc app=leetcode id=46 lang=cpp
 *
 * [46] Permutations
 */

// @lc code=start
#include <vector>
#include <unordered_set>
class Solution {
public:
    std::vector<std::vector<int>> permute(std::vector<int>& nums) {
        std::vector<int> selected;
        backtrace(nums, selected);
        return answer;
    }
private:
    std::vector<std::vector<int>> answer;
    std::unordered_set<int> is_have;
    void backtrace(std::vector<int>& nums, std::vector<int>& selected)
    {
        if (selected.size() == nums.size())
        {
            answer.push_back(selected);
            return;
        }
        
        for (int index = 0; index < nums.size(); index++)
        {
            if (!is_have.count(nums[index]))
            {
                selected.push_back(nums[index]);
                is_have.insert(nums[index]);
                backtrace(nums, selected);
                selected.pop_back();
                is_have.erase(nums[index]);
            }
        }
        
    }
};
// @lc code=end

