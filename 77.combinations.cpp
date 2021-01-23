/*
 * @lc app=leetcode id=77 lang=cpp
 *
 * [77] Combinations
 */

// @lc code=start
#include <vector>
#include <unordered_set>

class Solution {
public:
    std::vector<std::vector<int>> combine(int n, int k) {
        std::vector<int> comb;
        _i_max = n;
        combineBPImpl(1, k);
        return _answer;
    }
private:
    int _i_max;
    std::vector<int> _v_comb;
    std::vector<std::vector<int>> _answer;

    void combineBPImpl(int start, int k)
    {
        //
        if (k <= 0)
        {
            _answer.push_back(_v_comb);
            return;
        }
        //
        for (int i = start; i <= _i_max; i++)
        {
            _v_comb.push_back(i);
            combineBPImpl(i+1, k-1);
            _v_comb.pop_back();          
        }
    }
};
// @lc code=end

