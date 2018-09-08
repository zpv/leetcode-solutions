#include <stdlib.h>
#include <algorithm>

using namespace std;
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
static int x = []{
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    return 0;
}();

class Solution {
  public:
    struct TreeNode {
        int val;
        TreeNode *left;
        TreeNode *right;
        TreeNode(int x) : val(x), left(NULL), right(NULL) {}
    };

    int maxDepth(TreeNode* root) {
      return maxDepthHelper(root, 0);
    }

    int maxDepthHelper(TreeNode* root, int depth) {
      if (root != NULL) {
        return max(maxDepthHelper(root->left, depth + 1), maxDepthHelper(root->right, depth + 1));
      }
      return depth;
    }
};

int main() {
  return 0;
}