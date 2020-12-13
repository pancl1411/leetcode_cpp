/*
 * @lc app=leetcode id=117 lang=cpp
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/
#include <vector>
#include <stack>
class Solution {
public:
    Node* connect(Node* root) {
        std::vector<std::stack<Node*>> next_node;
        connectImpl(root, 0, next_node);
        return root;
    }
private:
    void connectImpl(Node* root, int level, std::vector<std::stack<Node*>>& next_node)
    {
        //
        if (root == nullptr) return;
        //
        if (next_node.size() <= level)
            next_node.push_back(std::stack<Node*>());
        //
        connectImpl(root->right, level+1, next_node);
        connectImpl(root->left, level+1, next_node);
        //
        if (!next_node[level].empty())
        {
            root->next = next_node[level].top();
            next_node[level].pop();
        }
        next_node[level].push(root);
    }
};
// @lc code=end

