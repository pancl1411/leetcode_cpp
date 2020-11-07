/*
 * @lc app=leetcode id=53 lang=cpp
 *
 * [53] Maximum Subarray
 */
#include <vector>
#include <algorithm>
// @lc code=start
class Solution {
public:
    // (1) O(n) solution
    
    int maxSubArray(std::vector<int>& nums) {
        int answer = nums[0], max_sum = nums[0];
        // O(n) loop for answer
        for (size_t i = 1; i < nums.size(); i++)
        {
            max_sum = std::max(max_sum + nums[i], nums[i]);
            answer = std::max(answer, max_sum);
        }
        return answer;
    }

    /*
    // record every answer if required
    int maxSubArray(std::vector<int>& nums) {
        // dp[i] means max subarray end with nums[i]
        int dp[nums.size()];
        dp[0] = nums[0];
        int answer = nums[0];
        // O(n) loop for answer
        for (size_t i = 1; i < nums.size(); i++)
        {
            dp[i] = std::max(dp[i - 1] + nums[i], nums[i]);
            answer = std::max(answer, dp[i]);
        }
        return answer;
    }
    */
};
// @lc code=end

