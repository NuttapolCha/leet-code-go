package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	// tree := &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val:   9,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 20,
	// 		Left: &TreeNode{
	// 			Val:   15,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 			Left: &TreeNode{
	// 				Val:   12,
	// 				Left:  nil,
	// 				Right: nil,
	// 			},
	// 			Right: &TreeNode{
	// 				Val:   13,
	// 				Left:  nil,
	// 				Right: nil,
	// 			},
	// 		},
	// 	},
	// }

	fmt.Println(isBalanced(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(tree *TreeNode) int {
	if tree == nil {
		return -1
	}
	return 1 + max(height(tree.Left), height(tree.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
