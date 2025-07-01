package main

import (
	"algo/binary_tree"
	"fmt"
)

func main() {
	root := &binary_tree.TreeNode{
		Left: &binary_tree.TreeNode{
			Left: &binary_tree.TreeNode{
				Left: &binary_tree.TreeNode{},
			},
		},
		Right: &binary_tree.TreeNode{},
	}

	//maxDept := binary_tree.GetMaxDepth(root)
	//fmt.Println(maxDept)

	//nums := binary_tree.NumsOfKLevelTreeNode(root, 2)
	//fmt.Println(nums)

	isBalance, depth := binary_tree.IsBalanceTree(root)
	fmt.Println(isBalance, depth)
}
