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

	converted := increasingBST(tree)
	bs, _ := json.MarshalIndent(converted, "", "\t")
	fmt.Println(string(bs))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	sorted := appendInorder(root, []int{})

	ret := new(TreeNode)
	next := ret

	for i := 0; i < len(sorted); i++ {
		if i != len(sorted)-1 { // i.e. not the last index
			next.Right = new(TreeNode)
		}
		next.Val = sorted[i]
		next = next.Right
	}

	return ret
}

func appendInorder(tree *TreeNode, list []int) []int {
	if tree == nil {
		return list
	}

	list = appendInorder(tree.Left, list)
	list = append(list, tree.Val)
	list = appendInorder(tree.Right, list)

	return list
}
