/*
 * @lc app=leetcode id=51 lang=cpp
 *
 * [51] N-Queens
 */

// @lc code=start
#include <vector>
#include <string>
class Solution {
public:
    std::vector<std::vector<std::string>> solveNQueens(int n) {
        std::vector<std::string> board = std::vector<std::string>(n, std::string(n, '.'));
        backtraceNQueensImpl(board, 0);
        return answer;
    }
private:
    std::vector<std::vector<std::string>> answer;
    void backtraceNQueensImpl(std::vector<std::string>& board, int col)
    {
        // find a result
        if (board.size() == col)
        {
            answer.push_back(board);
            return;
        }
        //
        for (int row = 0; row < board.size(); row++)
        {
            if (isVaildPlace(board, col, row))
            {
                board[col][row] = 'Q';
                backtraceNQueensImpl(board, col + 1);
                board[col][row] = '.';
            }
            
        }
    }

    bool isVaildPlace(std::vector<std::string>& board, int col, int row)
    {
        for (int i = col - 1; i >= 0; i--)
        {
            // the row has queue
            if (board[i][row] == 'Q')
                return false;
        }

        for (int i = col-1, j = row-1; i >= 0 && j >= 0; i--, j--)
        {
            // the left-up has queue
            if (board[i][j] == 'Q')
                return false;
        }

        for (int i = col-1, j = row+1; i >= 0 && j < board.size(); i--, j++)
        {
            // the right-up has queue
            if (board[i][j] == 'Q')
                return false;
        }

        return true;
    }
};
// @lc code=end

