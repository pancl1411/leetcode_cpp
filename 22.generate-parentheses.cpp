/*
 * @lc app=leetcode id=22 lang=cpp
 *
 * [22] Generate Parentheses
 */

// @lc code=start
#include <vector>
#include <string>
class Solution {
public:
    std::vector<std::string> generateParenthesis(int n) {
        std::string parentheses = std::string("");
        backtrace(parentheses, 0, n);
        return answer;
    }
private:
    std::vector<std::string> answer;
    void backtrace(std::string& parentheses, int left_count, int n)
    {
        // get one result
        if (parentheses.size() == 2 * n)
        {
            answer.push_back(parentheses);
            return;
        }
        // CASE 1: push '('
        if (left_count < n)
        {
            parentheses.push_back('(');
            backtrace(parentheses, left_count + 1, n);
            parentheses.pop_back();
        }
        // CASE 2: push ')'
        // 注意合法的圆括号')' 必须不多余'('数量，即先有'('后有')';
        if (parentheses.size() < 2 * left_count)
        {
            parentheses.push_back(')');
            backtrace(parentheses, left_count, n);
            parentheses.pop_back();
        }
    }
};
// @lc code=end

