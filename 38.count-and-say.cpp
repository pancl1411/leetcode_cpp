/*
 * @lc app=leetcode id=38 lang=cpp
 *
 * [38] Count and Say
 */
#include <string>
// @lc code=start
class Solution {
public:
    std::string say(const std::string& content)
    {
        std::string answer;
        // digit init (string never empty)
        char number = content[0] ;
        int count = 1;
        // loop until end
        char value;
        for (size_t index = 1; index < content.size(); index++)
        {
            value = content[index];
            if (value == number)
            {
                count++;
            }
            else
            {
                answer.push_back(count + '0');
                answer.push_back(number);
                number = value;
                count = 1;
            }
        }
        // append the last element
        answer.push_back(count + '0');
        answer.push_back(number);
        return answer;
    }

    std::string countAndSay(int n) {
        if (n == 1) return "1";
        return say(countAndSay(n - 1));
    }

/*
    std::string countAndSay(int n)
    {
        if (n == 1) return "1";
        std::string answer = "1";
        for (int i = 2; i <= n; i++)
        {
            answer = say(answer);
        }
        return answer;
    }
*/
};
// @lc code=end

