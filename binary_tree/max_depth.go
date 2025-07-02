package main

import (
	"container/list"
	"fmt"
	"github.com/samber/lo"
)

type TreeNode struct {
	ID     int
	Val    int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

// GetMaxDepth 获取node节点为根的树的最大深度
func GetMaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDept := GetMaxDepth(node.Left)
	rightDept := GetMaxDepth(node.Right)
	return lo.Max([]int{leftDept, rightDept}) + 1
}

// NumsOfKLevelTreeNode 求二叉树第K层的节点个数
func NumsOfKLevelTreeNode(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}

	result := 0
	if depth == 0 {
		result = 1
	}
	result += NumsOfKLevelTreeNode(root.Left, depth-1)
	result += NumsOfKLevelTreeNode(root.Right, depth-1)
	return result
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// IsBalanceTree 是否为平衡二叉树 空树为一个平衡二叉树 平衡二叉树的任意一个节点的左右子树高度小于等于1
func IsBalanceTree(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalance, leftDepth := IsBalanceTree(root.Left)
	isRightBalance, rightDepth := IsBalanceTree(root.Right)

	depthResult := lo.Max([]int{leftDepth, rightDepth}) + 1
	isBalance := isLeftBalance && isRightBalance && absInt(leftDepth-rightDepth) <= 1
	return isBalance, depthResult
}

// IsCBT 判断二叉树是否为完全二叉树
func IsCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root)
	shouldBeLeaf := false

	for queue.Len() > 0 {
		node := queue.Front()
		queue.Remove(node)

		nodeLeft := node.Value.(*TreeNode).Left
		nodeRight := node.Value.(*TreeNode).Right

		// 存在一个节点，其左子树为空，但是右子树非空
		if nodeLeft == nil && nodeRight != nil {
			return false
		}

		// 该节点应该为叶子节点，但是发现子树非空
		if shouldBeLeaf && (nodeLeft != nil || nodeRight != nil) {
			return false
		}

		if nodeLeft != nil {
			queue.PushBack(nodeLeft)
		}

		if nodeRight != nil {
			queue.PushBack(nodeRight)
		} else {
			shouldBeLeaf = true
		}
	}
	return true
}

// isSameTree 判定两个二叉树是否相同
func isSameTree(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	leftSame := isSameTree(node1.Left, node2.Left)
	rightSame := isSameTree(node1.Right, node2.Right)
	return leftSame && rightSame
}

// isMirrorTree 判断两个二叉树是否互为镜像
func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}
	return isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

// mirrorTreeNode 翻转二叉树
func mirrorTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	left := mirrorTree(node.Left)
	right := mirrorTree(node.Right)

	node.Left = right
	node.Right = left
	return node
}

// 二叉树的前序遍历
func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
}

// 二叉树的中序遍历
func inOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	preOrderTraverse(root.Left)
	fmt.Println(root.Val)
	preOrderTraverse(root.Right)
}

// 二叉树的后序遍历
func postOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
	fmt.Println(root.Val)
}

// 判断二叉树是否为合法的二叉查找树
func IsValidBST(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	curVal := root.Val
	leftValid := IsValidBST(root.Left, minVal, lo.Min([]int{maxVal, curVal}))
	rightValid := IsValidBST(root.Right, lo.Max([]int{minVal, curVal}), maxVal)
	return leftValid && rightValid && curVal >= minVal && curVal <= maxVal
}

// treeMaxDistance ⼆叉树内两个节点的最⻓距离 最大距离 & 最大深度
func treeMaxDistance(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftDis, leftDept := treeMaxDistance(node.Left)
	rightDis, rightDept := treeMaxDistance(node.Right)

	// 求解深度
	curDept := lo.Max([]int{leftDept, rightDept}) + 1
	// 求解距离
	leftRightMax := lo.Max([]int{leftDis, rightDis})
	curDis := lo.Max([]int{leftRightMax, leftDept + rightDept + 1})
	return curDis, curDept
}

func TreeDepth(root *TreeNode, curDept int, depthMap map[*TreeNode]int) {
	if root == nil {
		return
	}

	depthMap[root] = curDept

	TreeDepth(root.Left, curDept+1, depthMap)
	TreeDepth(root.Right, curDept+1, depthMap)
	return
}

// lca 求最近公共祖先
func lca(node1 *TreeNode, node2 *TreeNode, depthMap map[*TreeNode]int) *TreeNode {
	node1Cur := node1
	node2Cur := node2

	for depthMap[node1Cur] > depthMap[node2Cur] {
		node1Cur = node1Cur.Parent
	}
	for depthMap[node2Cur] > depthMap[node1Cur] {
		node2Cur = node2Cur.Parent
	}

	for node1Cur != node2Cur {
		node1Cur = node1Cur.Parent
		node2Cur = node2Cur.Parent
	}
	return node1Cur
}

// 前序 后序 无栈遍历构造二叉树 根左右 左右根
// curNode 当前要构造的节点
// tree 待构造的树
// 前序
// 后序
//func constructTree(curNode int, tree []int, preOrder []int, postOrder []int) int {
//	if len(preOrder) == 0 {
//		return -1
//	}
//
//	curRoot := preOrder[0]
//	// 枚举左子树的长度
//
//}

// 输入一个二叉树 和 一个整数，打印出二叉树中节点值的和等于输入整数所有的路径

func main() {
	rootNode := &TreeNode{
		ID:  1,
		Val: 1,
	}
	node1 := &TreeNode{
		ID:  2,
		Val: -1,
	}
	node2 := &TreeNode{
		ID:  3,
		Val: 2,
	}
	node3 := &TreeNode{
		ID:  4,
		Val: 10,
	}
	rootNode.Left = node1
	rootNode.Right = node2
	node1.Right = node3

	node1.Parent = rootNode
	node2.Parent = rootNode
	node3.Parent = node1

	//maxDept := binary_tree.GetMaxDepth(root)
	//fmt.Println(maxDept)

	//nums := binary_tree.NumsOfKLevelTreeNode(root, 2)
	//fmt.Println(nums)

	//isBalance, depth := IsBalanceTree(root)
	//fmt.Println(isBalance, depth)

	//fmt.Println(treeMaxDistance(root))

	depthMap := make(map[*TreeNode]int, 0)
	TreeDepth(rootNode, 0, depthMap)
	lcaNode := lca(node1, node3, depthMap)
	fmt.Println(lcaNode.ID)
}
