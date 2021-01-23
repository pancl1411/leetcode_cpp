/*
 * @lc app=leetcode id=60 lang=cpp
 *
 * [60] Permutation Sequence
 */

// @lc code=start
#include <string>
#include <vector>
#include <unordered_set>

class Solution {
public:
    std::string getPermutation(int n, int k) {
        //
        std::string answer;
        // N的阶乘
        std::vector<int> factorial = std::vector<int>(n, 0);
        int total = 1;
        for (int i = 1; i <= n; i++)
        {
            total *= i;
            factorial[i-1] = total;
        }
        // 初始化可选数字
        initCandidate(n);
        // 根据计数循环查找
        int step = 0;
        int off = 0;
        int move = 0;
        int digit = 0;
        while (n > 1)
        {
            // 注意存在off==0的特殊情况，后续需要单独处理
            step = k / factorial[n-2];
            off = k % factorial[n-2];
            // push一个数字
            move = (off == 0) ? step : step + 1;
            digit = getKthCandidate(move);
            _candidate.erase(digit);
            answer.push_back('0' + digit);
            // 循环更新
            k =  (off == 0) ? factorial[n-2] : off;
            n--;
        }
        // n == 1
        answer.push_back('0' + getKthCandidate(1));
        //
        return answer;
    }
    /* 全部求取，耗时很长
    // std::string getPermutation(int n, int k) {
    std::string getPermutation(int n, int k) {
        std::string seq;
        initCandidate(n);
        kthPermutationBP(n, k, seq);
        return _s_answer;
    }*/
private:
    int _i_max_cand;
    int _i_count;
    std::string _s_answer;
    std::unordered_set<int> _candidate;
    //
    void initCandidate(int n)
    {
        _i_max_cand = n;
        for (int i = 1; i <= n; i++)
        {
            _candidate.insert(i);
        }
    }
    //
    int getKthCandidate(int k)
    {
        int has = 0;
        for (int i = 1; i <= _i_max_cand; i++)
        {
            if (_candidate.count(i) > 0)
            {
                has++;
                if (has == k)
                {
                    return i;
                }
            }
            
        }

        return 0;
    }
    //
    void kthPermutationBP(int n, int k, std::string& seq)
    {
        //
        if (_candidate.empty())
        {
            _i_count++;
            if (_i_count == k)
            {
                _s_answer = seq;
            }
            return;
        }
        //
        if (!_s_answer.empty())
        {
            return;
        }
        //
        for (int i = 1; i <= n; i++)
        {
            if (_candidate.count(i) > 0)
            {
                _candidate.erase(i);
                seq.push_back('0'+ i);
                kthPermutationBP(n, k, seq);
                _candidate.insert(i);
                seq.pop_back();
            }
        }
    }
};
// @lc code=end

