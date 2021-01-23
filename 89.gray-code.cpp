/*
 * @lc app=leetcode id=89 lang=cpp
 *
 * [89] Gray Code
 */

// @lc code=start
#include <vector>
class Solution {
public:
    std::vector<int> grayCode(int n) {
        std::vector<int> answer = std::vector<int>(1, 0);
        for (int i = 0; i < n; i++)
        {
            int size = answer.size();
            // 必须反向迭代，确保i跳转时code连续
            for (int j = size -1; j >= 0; j--)
            {
                // 改变高位0->1 （第i位）
                answer.push_back(answer[j] | 1<<i);
            }
        }
        return answer;
    }
/*
    std::vector<int> grayCode(int n) {
        grayCodeImpl(n, 0, 0);
        return _answer;
    }
*/
private:
    std::vector<int> _answer;
    // start 从0计数
    void grayCodeImpl(int n, int start, int code)
    {
        _answer.push_back(code);
        if (start >= n)
        {
            return;
        }
        //
        int reverse = 0;
        for (int i = start; i < n; i++)
        {
            reverse = reverseOneBit(code, i);
            grayCodeImpl(n, i+1, reverse);

        }
    }

    int reverseOneBit(int code, int bit_offset)
    {
        int result = 0;
        int shift = 1 << bit_offset;
        if (shift & code)
        {
            result = code - shift;
        }
        else
        {
            result = code + shift;
        }
        return result;
    }
};
// @lc code=end

