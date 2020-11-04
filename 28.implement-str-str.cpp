/*
 * @lc app=leetcode id=28 lang=cpp
 *
 * [28] Implement strStr()
 */
#include <string>
#include <vector>
// @lc code=start
class Solution {
private:
    int _i_pattern_len;
    int* _p_kmp_dp;
public:
    void kmp_init(std::string pattern)
    {
        //
        _i_pattern_len = pattern.size();
        _p_kmp_dp = new int[_i_pattern_len * 256]();
        // base status
        _p_kmp_dp[0 * 256 + pattern[0]] = 1;
        //printf("_p_kmp_dp[%d] = %d\n", pattern[0], _p_kmp_dp[0 * 256 + pattern[0]]);
        int X = 0;
        //
        for (int status = 1; status < _i_pattern_len; status++)
        {
            //
            for (int ascii = 0; ascii < 256; ascii++)
            {
                _p_kmp_dp[status * 256 + ascii] = _p_kmp_dp[X * 256 + ascii];
            }
            //
            _p_kmp_dp[status * 256 + pattern[status]] = status + 1;
            //printf("_p_kmp_dp[%d] = %d\n", status * 256 + pattern[status], _p_kmp_dp[status * 256 + pattern[status]]);
            X = _p_kmp_dp[X * 256 + pattern[status]];
        }
    }

    int kmp_search(std::string haystack)
    {
        int status = 0;
        for (int index = 0; index < haystack.size(); index++)
        {
            status = _p_kmp_dp[status * 256 + haystack[index]];
            //printf("[_i_pattern_len:%d][haystack[index]:%d][status:%d]\n", _i_pattern_len, haystack[index], status);
            if (status == _i_pattern_len)
            {
                return index - _i_pattern_len + 1;
            }
        }
        return -1;
    }

    int strStr(std::string haystack, std::string needle) {
        //
        if (needle.empty())
        {
            return 0;
        }
        //
        kmp_init(needle);
        return kmp_search(haystack);
    }
};
// @lc code=end

