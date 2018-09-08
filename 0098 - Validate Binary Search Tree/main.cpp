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

    bool isValidBST(TreeNode* root) {
      return isValidBSTHelper(root, 0, 0, false, false);
    }
    
    bool isValidBSTHelper(TreeNode* root, int min, int max, bool hasMin, bool hasMax) {
      if (root == NULL) {
        return true;
      }

      if ((hasMax && root->val >= max) || (hasMin && root->val <= min)) {
        return false;
      }

      return isValidBSTHelper(root->left, min, root->val, hasMin, true) && isValidBSTHelper(root->right, root->val, max, true, hasMax);
    }
};

int main() {
  return 0;
}