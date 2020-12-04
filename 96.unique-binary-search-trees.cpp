/*
 * @lc app=leetcode id=96 lang=cpp
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
#include <vector>
class Solution {
public:
    // (1) Divide and Conquer
    /*
    int DivideConquer(int min, int max)
    {
        //
        if (min > max)
        {
            return 0;
        }
        else if (min == max)
        {
            return 1;
        }
        //
        int answer = 0;
        for (int index = min; index <= max; index++)
        {
            int left = DivideConquer(min, index -1);
            int right = DivideConquer(index + 1, max);
            if (left == 0)
            {
                answer += right;
            }
            else if (right == 0)
            {
                answer += left;
            }
            else
            {
                answer += left * right;
            }
        }
        return answer;
    }

    int numTrees(int n) {
        return DivideConquer(1, n);    
    }
    */
   //(2) DP Solution
    int numTrees(int n) {
        std::vector<int> dp(n+1, 0);
        dp[0] = dp[1] = 1;
        for (int i = 2; i <= n; i++)
        {
            for (int j = 1; j <= i; j++)
            {
                dp[i] += dp[j-1] * dp[i-j];
            }
        }
        return dp[n];
    }   
};
// @lc code=end

