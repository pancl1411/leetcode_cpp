/*
 * @lc app=leetcode id=27 lang=cpp
 *
 * [27] Remove Element
 */
#include <vector>
// @lc code=start
class Solution {
public:
/*
Accepted
113/113 cases passed (0 ms)
Your runtime beats 100 % of cpp submissions
Your memory usage beats 100 % of cpp submissions (9.2 MB)
*/
    // iteratively
    int removeElement(std::vector<int>& nums, int val) {
        // special case
        if (nums.empty())
        {
            return 0;
        }
        // loop until swap finish
        int head = 0;
        int tail = nums.size() - 1;
        while (head <= tail && tail >= 0)
        {   
            // swap value between head and tail
            if (nums[head] == val)
            {
                nums[head] = nums[tail];
                nums[tail] = val;
                tail--;
            }
            // move head only not val
            if (nums[head] != val)
            {
                head++;
            }
        }
        //
        return head;
    }
};
// @lc code=end

