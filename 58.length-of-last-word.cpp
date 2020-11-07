/*
 * @lc app=leetcode id=58 lang=cpp
 *
 * [58] Length of Last Word
 */
#include <string>
// @lc code=start
class Solution {
public:
    int lengthOfLastWord(std::string s) {
        // special case
        if (s.empty())
        {
            return 0;
        }
        //
        int len = 0;
        for (int i = s.size() - 1; i >= 0; i--)
        {
            if (s[i] != ' ')
            {
                len++;
            }
            // avoid ' ' in the end
            else if (len != 0)
            {
                return len;
            }
        }
        return len;
    }
};
// @lc code=end

