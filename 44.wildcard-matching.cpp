/*
 * @lc app=leetcode id=44 lang=cpp
 *
 * [44] Wildcard Matching
 */

// @lc code=start
#include <vector>
#include <string>
class Solution {
public:
    bool isMatch(std::string s, std::string p) {
        dp = std::vector<std::vector<bool>>(s.size()+1, std::vector<bool>(p.size()+1, false));
        dp[0][0] = true;
        // i 表示s中前i个元素，从1计数， i==0时s为空
        // j 表示p中前j个元素，从1计数， j==0时p为空
        // dp[i][j] 表示s前i个元素和p前j个元素是否匹配
        for (size_t i = 0; i <= s.size(); i++)
        {
            for (size_t j = 1; j <= p.size(); j++)
            {
                // s为空，特殊处理
                if (i == 0)
                {
                    if (p[j-1] == '*' )
                    {
                        if (j == 1)
                        {
                            dp[0][1] = true;
                        }
                        else
                        {
                            dp[0][j] = dp[0][j-1]; 
                        }
                    }
                }
                else
                {
                    // CASE 1: 字符相等
                    // CASE 2: '?'可以匹配任意一个字符
                    if (s[i-1] == p[j-1] || p[j-1] == '?')
                    {
                        dp[i][j] = dp[i-1][j-1];
                    }
                    // CASE 3: '*'可以匹配:
                    //  1)任意一个字符dp[i-1][j-1]
                    //  2)或者为空dp[i][j-1]
                    //  3)或者多个字符dp[i-1][j]
                    else if (p[j-1] == '*')
                    {
                        dp[i][j] = (dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]);
                    }
                }
            }
        }
        return dp[s.size()][p.size()];
    }
private:
    std::vector<std::vector<bool>> dp;
};
// @lc code=end

