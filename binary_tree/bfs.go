package main

import (
	"container/list"
	"slices"
)

type QueueNode struct {
	treeNode *TreeNode
	depth    int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make(map[int][]int, 0)
	queue := list.New()
	queue.PushBack(QueueNode{treeNode: root, depth: 0})
	for queue.Len() > 0 {
		frontNode := queue.Front()
		queue.Remove(frontNode)

		u := frontNode.Value.(QueueNode)
		result[u.depth] = append(result[u.depth], u.treeNode.Val)

		if u.treeNode.Left != nil {
			queue.PushBack(QueueNode{treeNode: u.treeNode.Left, depth: u.depth + 1})
		}
		if u.treeNode.Right != nil {
			queue.PushBack(QueueNode{treeNode: u.treeNode.Right, depth: u.depth + 1})
		}
	}

	maxDepth := 0
	for k, _ := range result {
		maxDepth = slices.Max([]int{maxDepth, k})
	}

	res := make([][]int, maxDepth+1)
	for k, v := range result {
		res[k] = append(res[k], v...)
	}
	return res
}
