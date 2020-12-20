/*
 * @lc app=leetcode id=37 lang=cpp
 *
 * [37] Sudoku Solver
 */

// @lc code=start
#include <vector>
class Solution {
public:
    void solveSudoku(std::vector<std::vector<char>>& board)
    {
        backtraceImpl(board);
    }
private:
    bool backtraceImpl(std::vector<std::vector<char>>& board) 
    {
        for (int col = 0; col < board.size(); col++)
        {
            for (int row = 0; row < board.size(); row++)
            {
                // （1）不是空格，跳过
                if (board[col][row] != '.')
                    continue;
                // （2）对空格进行填数，递归后回溯
                for (char digit = '1'; digit <= '9'; digit++)
                {
                    if (isVaildDigit(board, col, row, digit))
                    {
                        board[col][row] = digit;
                        // (3) 说明递归由（5）完成，得到答案，避免下方回溯，返回
                        if (backtraceImpl(board))
                            return true;
                        board[col][row] = '.';
                    }
                }
                // （4）未进入(3), 说明本轮填数失败，回溯后重新开始
                return false;
            }
        }
        // (5) 说明全部从(1)出跳转出来
        return true;
    }

    bool isVaildDigit(std::vector<std::vector<char>>& board, int col, int row, char digit)
    {
        // the same col/row has this digit 
        for (int i = 0; i < board.size(); i++)
        {
            if (board[col][i] == digit || board[i][row] == digit)
            {
                return false;
            }
        }
        // the 3x3 sub-boxes has this digit
        // 注意不是上下左右相邻（边界点也在3x3子正方形内）
        for (int i = 3 * (col / 3 ); i < 3 * (col / 3) + 3; i++)
        {
            for (int j = 3 * (row / 3); j < 3 * (row / 3) + 3; j++)
            {
                if (board[i][j] == digit)
                {
                    return false;
                }
            }
        }
        // 3x3 sub-boxes false understand
        /*for (int i = col -1; i <= col + 1 && i < board.size(); i++)
        {
            for (int j = row -1; j <= row + 1 && j < board.size(); j++)
            {
                if (i >=0 && j >=0 && board[i][j] == digit)
                {
                    return false;
                }
            }
        }
        */
        //
        return true;
    }  
};
// @lc code=end

