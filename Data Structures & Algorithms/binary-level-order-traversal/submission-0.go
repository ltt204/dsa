// 102. Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal

package binarylevelordertraversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var result [][]int

	if root == nil {
		return result
	}

	queue = append(queue, root)
	for true {
		var level []int
		var children []*TreeNode
		for len(queue) > 0 {
			pop := queue[0]
			queue = queue[1:]

			level = append(level, pop.Val)

			if pop.Left != nil {
				children = append(children, pop.Left)
			}

			if pop.Right != nil {
				children = append(children, pop.Right)
			}
		}

		if len(level) == 0 {
			break
		}

		queue = append([]*TreeNode{}, children...)
		result = append(result, level)
	}

	return result
}
