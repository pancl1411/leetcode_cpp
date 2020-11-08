/*
 * @lc app=leetcode id=67 lang=cpp
 *
 * [67] Add Binary
 */
#include <string>
// @lc code=start
class Solution {
public:
    std::string addBinary(std::string a, std::string b) {
        // get longer/shorter string
        std::string* longer = a.size() > b.size() ? &a : &b;
        std::string* shorter = a.size() > b.size() ? &b : &a;
        //
        char carry = '0';
        for (int i = longer->size() -1; i >= 0; i--)
        {
            int j = i + shorter->size() - longer->size();

            int bit_add = (carry - '0') + ((*longer)[i] - '0');
            if (j >= 0)
            {
                bit_add += ((*shorter)[j] - '0');
            }
            if (bit_add < 2)
            {
                carry = '0';
                (*longer)[i] = bit_add + '0';
                if (j < 0)
                {
                    // 加至中间结束
                    return *longer;
                }   
            }
            else
            {
                carry = '1';
                (*longer)[i] = bit_add - 2 + '0';
            }
        }
        //
        if (carry == '0')
        {
            return *longer;
        }
        // 相加后最高位还要进位的情形      
        std::string new_string = std::string(1, '1');
        new_string.append(*longer);
        return new_string;
    }
};
// @lc code=end

