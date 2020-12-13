package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

1、找根节点 前序遍历的第一个节点是根节点
2、找左子数 中序遍历的左子树在根节点左边
3、找右子数 中序遍历的右子数在根节点右边

得出解题：
递归的方式 分别创建 左子树和右子树


*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 || len(preorder) != len(inorder) {
		return nil
	}
	// 中序遍历根节点index
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			rootIndex = i
		}
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func main() {
	preorder, inorder := []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
