/*
 * @lc app=leetcode id=93 lang=cpp
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
#include <vector>
#include <string>
class Solution
{
public:
    std::vector<std::string> restoreIpAddresses(std::string s)
    {
        if (s.size() >= 4 && s.size() <= 12)
        {
            backTrack(s, 0);
        }
        return _answer;
    }

private:
    std::vector<std::string> _answer;
    //IP[0].IP[1].IP[2].IP[3]
    std::vector<std::string> _IP;
    void backTrack(std::string &s, int start)
    {
        //
        if (start >= s.size())
        {
            if (_IP.size() == 4)
            {
                _answer.push_back(_IP[0] + "." + _IP[1] + "." + _IP[2] + "." + _IP[3]);
            }
            return;
        }
        //
        std::string add;

        for (int j = 0; add.size() <= 3 && start+j < s.size(); j++)
        {
            add.push_back(s[start + j]);
            if (isVaildAdd(add))
            {
                _IP.push_back(add);
                backTrack(s, start + j + 1);
                _IP.pop_back();
            }
        }
    }

    bool isVaildAdd(std::string &add)
    {
        int value = std::stoi(add);
        if (value < 0 || value > 255)
        {
            return false;
        }
        if (add.size() > 1 && add[0] == '0')
        {
            return false;
        }
        
        return true;
    }
};
// @lc code=end
