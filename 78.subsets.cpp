/*
 * @lc app=leetcode id=78 lang=cpp
 *
 * [78] Subsets
 */

// @lc code=start
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> subsets(std::vector<int>& nums) {
        std::vector<int> set;
        //_answer.push_back(set); // push empty set
        subsetsBPImpl(nums, 0, set);
        return _answer;        
    }
private:
    std::vector<std::vector<int>> _answer;
    void subsetsBPImpl(std::vector<int>& nums, int start, std::vector<int>& set)
    {
        //
        _answer.push_back(set);
        //
        if (start >= nums.size())
        {
            return;
        }
        //
        for (int i = start; i < nums.size(); i++)
        {
            set.push_back(nums[i]);
            subsetsBPImpl(nums, i+1, set);
            set.pop_back();
        }

    }
};
// @lc code=end

