/*
 * @lc app=leetcode id=116 lang=cpp
 *
 * [116] Populating Next Right Pointers in Each Node
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
        Node* left_node = root;
        Node* current_node = nullptr;
        while (left_node && left_node->left)
        {
            current_node = left_node;
            while (current_node)
            {
                current_node->left->next = current_node->right;
                if(current_node->next)
                    current_node->right->next = current_node->next->left;
                current_node = current_node->next;
            }
            left_node = left_node->left;
        }
        return root;
    }
/*
private:
    std::vector<std::stack<Node*>> nextNode;
    void connectImpl(Node* root, int level) {
        if (root == nullptr)
        {
            return;
        }
        if (nextNode.size() <= level)
        {
            nextNode.push_back(std::stack<Node*>());
        }
        //
        connectImpl(root->right, level + 1);
        connectImpl(root->left, level + 1);
        if (!nextNode[level].empty())
        {
            root->next = nextNode[level].top();
            nextNode[level].pop();
        }
        nextNode[level].push(root);
    }
*/
};
// @lc code=end

