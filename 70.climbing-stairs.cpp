/*
 * @lc app=leetcode id=70 lang=cpp
 *
 * [70] Climbing Stairs
 */
#include <vector>
// @lc code=start
class Solution {
private:
    vector<int> buffer;
public:
    // Recursion with Memoization
    int climbStairsRecursion(int n)
    {
        // special case
        if (n == 1)
        {
            return 1;
        }
        else if (n == 2)
        {
            return 2;
        }
        // read buffer
        if (buffer[n-1] != 0)
        {
            return buffer[n-1];
        }
        // calcute it
        int n_1 = climbStairsRecursion(n - 1);
        int n_2 = climbStairsRecursion(n - 2);
        buffer[n-1] = n_2 + n_1;
        return buffer[n-1];
    }

    // Dynamic Programming
    int climbStairsDp(int n)
    {
        // special case
        if (n == 1)
        {
            return 1;
        }
        else if (n == 2)
        {
            return 2;
        }
        // read buffer
        buffer[0] = 1;
        buffer[1] = 2;
        for (int i = 2; i < n; i++)
        {
            buffer[i] = buffer[i-1] + buffer[i-2];
        }
        return buffer[n-1];
    }

    int climbStairs(int n) {
        buffer = vector<int>(n, 0);
        //return climbStairsRecursion(n);
        return climbStairsDp(n);
    }
};
// @lc code=end

