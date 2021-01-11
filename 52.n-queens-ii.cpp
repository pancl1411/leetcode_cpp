/*
 * @lc app=leetcode id=52 lang=cpp
 *
 * [52] N-Queens II
 */

// @lc code=start
#include <vector>
#include <string>
class Solution {
public:
    int totalNQueens(int n) {
        // speclal case
        if (n <= 0)
        {
            return 0;
        }
        //
        std::vector<std::string> board = std::vector<std::string>(n, std::string(n, '.'));
        totalNQueensImpl(board, 0);
        return this->_i_answer;
    }
private:
    int _i_answer;
    void totalNQueensImpl(std::vector<std::string>& board, int col)
    {
        //
        if (col == board.size())
        {
            this->_i_answer++;
            return;
        }
        //
        for (int row = 0; row < board.size(); row++)
        {
            if (isValidPlace(board, col, row))
            {
                board[col][row] = 'Q';
                totalNQueensImpl(board, col + 1);
                board[col][row] = '.';
            }
            
        }
        
    }

    bool isValidPlace(std::vector<std::string>& board, int col, int row)
    {
        for (int i = 0; i < col; i++)
        {
            if (board[i][row] == 'Q')
            {
                return false;
            }
        }

        for (int i = col-1, j = row-1; i >= 0 && j >= 0; i--, j--)
        {
            if (board[i][j] == 'Q')
            {
                return false;
            }
        }
        
        for (int i = col-1, j = row+1; i >= 0 && j < board.size(); i--, j++)
        {
            if (board[i][j] == 'Q')
            {
                return false;
            }
        }

        return true;
    }
};
// @lc code=end

