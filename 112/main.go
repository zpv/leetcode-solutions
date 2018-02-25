package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumHelper(root, sum, 0)
}

func hasPathSumHelper(root *TreeNode, sum int, rsf int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == rsf+root.Val
	}

	return hasPathSumHelper(root.Left, sum, rsf+root.Val) || hasPathSumHelper(root.Right, sum, rsf+root.Val)
}

func main() {

}
