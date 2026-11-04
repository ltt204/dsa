// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree

package maximumdepthofbinarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var result int

	var dfs func(root *TreeNode, depth int) int
	dfs = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth
		}

		result = max(dfs(root.Left, depth+1), dfs(root.Right, depth+1))

		return result
	}

	return dfs(root, 0)
}
