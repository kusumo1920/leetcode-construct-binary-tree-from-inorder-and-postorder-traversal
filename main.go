package main

import "fmt"

func main() {
	var mapper map[int]int
	fmt.Println(mapper == nil)
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	output := buildTree(inorder, postorder)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	mapper := make(map[int]int)
	postIdx := len(postorder) - 1
	var recursiveFn func(int, int) *TreeNode
	recursiveFn = func(inLeft int, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}
		rootVal := postorder[postIdx]
		root := &TreeNode{Val: rootVal}
		index := mapper[rootVal]
		postIdx--
		root.Right = recursiveFn(index+1, inRight)
		root.Left = recursiveFn(inLeft, index-1)
		return root
	}
	for i := 0; i < len(inorder); i++ {
		mapper[inorder[i]] = i
	}
	return recursiveFn(0, len(inorder)-1)
}
