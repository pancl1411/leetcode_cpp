/*
 * @lc app=leetcode id=79 lang=cpp
 *
 * [79] Word Search
 */

// @lc code=start
#include <vector>
#include <string>
class Solution {
public:
    // 可以优化,参考：https://leetcode.com/problems/word-search/discuss/27658/Accepted-very-short-Java-solution.-No-additional-space.
    bool exist(std::vector<std::vector<char>>& board, std::string word) {
        //
        this->_col_max = board.size();
        this->_row_max = board[0].size();
        this->_b_have_word = false;
        //
        for (int i = 0; i < this->_col_max; i++)
        {
            for (int j = 0; j < this->_row_max; j++)
            {
                //std::cout << "i:" << i << "\t j:" << j << std::endl;
                if (!this->_b_have_word && board[i][j] == word[0])
                {
                    board[i][j] = '\0';
                    findCharAdjacent(board, word, i, j, 1);
                    board[i][j] = word[0];
                }
            }
        }
        //
        return this->_b_have_word;
    }
private:
    int _col_max;
    int _row_max;
    bool _b_have_word;

    void findCharAdjacent(
        std::vector<std::vector<char>>& board, 
        std::string& word, 
        int col,
        int row,
        int kthWord
        )
    {
        //std::cout << "col:" << col << "\t row:" << row << "\t kthWord:" << kthWord << std::endl;
        //
        if (this->_b_have_word || kthWord >= word.size())
        {
            this->_b_have_word = true;
            return;
        }
        //
        if (isChar(board, word[kthWord], col-1 , row))
        {
            board[col-1][row] = '\0';
            findCharAdjacent(board, word, col-1, row, kthWord+1);
            board[col-1][row] = word[kthWord];
        }
        if (isChar(board, word[kthWord], col+1 , row))
        {
            board[col+1][row] = '\0';
            findCharAdjacent(board, word, col+1, row, kthWord+1);
            board[col+1][row] = word[kthWord];
        }
        if (isChar(board, word[kthWord], col , row-1))
        {
            board[col][row-1] = '\0';
            findCharAdjacent(board, word, col, row-1, kthWord+1);
            board[col][row-1] = word[kthWord];
        }
        if (isChar(board, word[kthWord], col , row+1))
        {
            board[col][row+1] = '\0';
            findCharAdjacent(board, word, col, row+1, kthWord+1);
            board[col][row+1] = word[kthWord];
        }
    }

    bool isChar(std::vector<std::vector<char>>& board, char obj, int col, int row)
    {  
        //
        if (col < 0 || col >= this->_col_max)
        {
            return false;
        }
        //
        if (row < 0 || row >= this->_row_max)
        {
            return false;
        }
        //
        return (board[col][row] == obj) ? true : false;
    }
};
// @lc code=end

