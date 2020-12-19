/*
 * @lc app=leetcode id=10 lang=cpp
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
#include <string>
#include <vector>
//
//  '.' Matches any single character.​​​​
//      匹配任意一个字符
//  '*' Matches zero or more of the preceding element.
//      前一元素出现0-N次（'*'的前面必须有有效的字符，可是'.'）
//          (注意可以消除前元素，如a*可表示为空)
//
class Solution {
public:
    bool isMatch(std::string s, std::string p) {
        //
        if (s.empty() && p.empty())
        {
            return true;
        }
        // dp[i][j]
        //  i表示s中的第i个元素，i==0表示s为空, i==1时表示s[0]
        //  j表示p中的第i个元素，i==0表示p为空
        std::vector<std::vector<bool>> dp (s.size()+1, std::vector<bool>(p.size()+1, false));
        dp[0][0] = true;
        // i>=1  
        for (int i = 0; i <= s.size(); i++)
        {
            for (int j = 1; j <= p.size(); j++)
            {
                // CASE 0: i==0, (s为空)
                if (i == 0)
                {
                    // a*表示为空, a*==null
                    if (p[j-1] == '*') {
                        dp[0][j] = dp[0][j-2];
                    }
                }
                // CASE 1: 字符相等
                // CASE 2: pattern为'.', s可为任意字符
                else if (s[i-1] == p[j-1] || p[j-1] == '.')
                {
                    dp[i][j] = dp[i-1][j-1];
                }
                // CASE 3: pattern为'*', 可分为子情况：
                //  1、 *的前一字符，和s的当前字符不等，a*只能当空处理，a*==null
                //  2、 *的前一字符，和s的当前字符相等，a*可以：
                //          1)当空a*==null；  dp[i][j-2]
                //          2）当单个前一字符 a*==a ;  dp[i-1][j]
                //          3) 当多个前一字符a* == aa; dp[i-1][j-1]
                else if (p[j-1] == '*')
                {
                    if (s[i-1] != p[j-2] && p[j-2] != '.')
                    {
                        dp[i][j] = dp[i][j-2];
                    }
                    else
                    {
                        dp[i][j] = (dp[i][j-2]  || dp[i-1][j] || dp[i-1][j-1]);
                    }
                }
                // CASE 4: 其他情况，默认false
            }
        }

        return dp[s.size()][p.size()];
    }
};
// @lc code=end

