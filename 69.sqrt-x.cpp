/*
 * @lc app=leetcode id=69 lang=cpp
 *
 * [69] Sqrt(x)
 */

// @lc code=start
class Solution {
public:
    //  Binary Search Solution
    int mySqrt(int x) {
        // special case
        if (x < 2)
        {
            return x;
        }
        // Binary Search
        int left = 1;
        int right = x / 2;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            int devide = x / mid;
            int devide_plus = x / (mid + 1);
            if (mid == devide || (mid < devide && mid+1 > devide_plus))
            {
                return mid;
            }
            else if (mid < devide)
            {
                left = mid + 1;
            }
            else if (devide < mid)
            {
                right = mid - 1;
            }
        }
        return left;
    }
};
// @lc code=end

