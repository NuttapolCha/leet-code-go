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

	fmt.Println(maxDepth(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	fmt.Printf("lh = %d\n", lh)

	rh := height(root.Right)
	fmt.Printf("rh = %d\n", rh)

	if lh > rh {
		return lh + 1
	}
	return rh + 1
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
