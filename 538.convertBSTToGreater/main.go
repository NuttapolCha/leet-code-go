package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	converted := convertBST(tree)
	bs, _ := json.MarshalIndent(converted, "", "\t")
	fmt.Printf("converted = %s\n", string(bs))
	// fmt.Printf("sum = %d\n", reverseInorder(tree, 0))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	reverseInorder(root, 0)
	return root
}

func reverseInorder(tree *TreeNode, sum int) int {
	if tree == nil {
		return sum
	}
	sum = reverseInorder(tree.Right, sum)
	sum += tree.Val
	tree.Val = sum
	sum = reverseInorder(tree.Left, sum)

	return sum
}

// cloneBST is not for this problem. Just a practicing one
func cloneBST(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return &TreeNode{
			Val:   root.Val,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Left == nil {
		return &TreeNode{
			Val:   root.Val,
			Left:  nil,
			Right: cloneBST(root.Right),
		}
	}

	if root.Right == nil {
		return &TreeNode{
			Val:   root.Val,
			Left:  cloneBST(root.Left),
			Right: nil,
		}
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  cloneBST(root.Left),
		Right: cloneBST(root.Right),
	}
}
