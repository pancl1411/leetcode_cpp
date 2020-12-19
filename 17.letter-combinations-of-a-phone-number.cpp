/*
 * @lc app=leetcode id=17 lang=cpp
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
#include <string>
#include <vector>
#include <unordered_map>
#include <queue>
class Solution {
public:
    std::vector<std::string> letterCombinations(std::string digits) {
        std::vector<std::string> answer;
        // special case
        if (digits.empty())
        {
            return answer;
        }
        // map init
        build_phone_map();
        // queue init with first digit
        std::queue<std::string> loop_queue;
        for (auto one_letter : digit_map[digits[0]-'0'])
        {
            loop_queue.push(std::string(1, one_letter));
        }
        // BFS with queue
        for (int i = 1; i < digits.size(); i++)
        {
            int queue_size = loop_queue.size();
            for (int j = 0; j < queue_size; j++)
            {
                std::string old = loop_queue.front();
                loop_queue.pop();
                for (auto one_letter : digit_map[digits[i]-'0'])
                {
                    // 回溯法插入
                    old.push_back(one_letter);
                    loop_queue.push(old);
                    old.pop_back();
                }
            }
        }
        // convert to vector
        while (!loop_queue.empty())
        {
            std::string one_answer = loop_queue.front();
            loop_queue.pop();
            answer.push_back(one_answer);
        }
        return answer;
    }
private:
    std::unordered_map<int, std::string> digit_map;
    void build_phone_map()
    {
        digit_map[2] = "abc";
        digit_map[3] = "def";
        digit_map[4] = "ghi";
        digit_map[5] = "jkl";
        digit_map[6] = "mno";
        digit_map[7] = "pqrs";
        digit_map[8] = "tuv";
        digit_map[9] = "wxyz";
    }
};
// @lc code=end

