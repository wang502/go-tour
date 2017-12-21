package main

import (
	"fmt"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
	order int
}

func (root *treeNode) insert(val int) *treeNode {
	return root.insertHelper(val, 0)
}

func (root *treeNode) insertHelper(val, order int) *treeNode {
	if root == nil {
		return &treeNode{val, nil, nil, order}
	}
	if val <= root.val {
		root.order++
		root.left = root.left.insertHelper(val, order)
		root.right.incrementOrder()
	} else {
		root.right = root.right.insertHelper(val, root.order+1)
	}
	return root
}

func (root *treeNode) incrementOrder() {
	if root == nil {
		return
	}
	root.order++
	root.left.incrementOrder()
	root.right.incrementOrder()
}

func (root *treeNode) inorder() {
	if root == nil {
		return
	}
	root.left.inorder()
	fmt.Printf("val: %d order: %d\n", root.val, root.order)
	root.right.inorder()
}

func (root *treeNode) preorder() {
	if root == nil {
		return
	}
	fmt.Printf("tree's preorder traversal is: ")
	stack := []*treeNode{root}
	for len(stack) != 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curNode.val)
		if curNode.right != nil {
			stack = append(stack, curNode.right)
		}
		if curNode.left != nil {
			stack = append(stack, curNode.left)
		}
	}
	fmt.Printf("\n")
}

func ksmallestBST(root *treeNode, k int) int {
	if root == nil {
		return -1
	}
	if root.order+1 == k {
		return root.val
	} else if root.order+1 < k {
		return ksmallestBST(root.right, k)
	} else {
		return ksmallestBST(root.left, k)
	}
}

func (root *treeNode) ceil(num int) int {
	if root == nil {
		return -1
	}
	if root.val == num {
		return root.val
	} else if root.val < num {
		return root.right.ceil(num)
	} else {
		tmp := root.left.ceil(num)
		if tmp > num {
			return tmp
		}
		return root.val
	}
}

func (root *treeNode) isComplete() bool {
	queue := make([]*treeNode, 0)
	queue = append(queue, root)
	flag := false
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode.left != nil {
			if flag {
				return false
			}
			queue = append(queue, curNode.left)
		} else {
			flag = true
		}
		if curNode.right != nil {
			if flag {
				return false
			}
			queue = append(queue, curNode.right)
		} else {
			flag = true
		}
	}
	return true
}

func morrisInorder(root *treeNode) []int {
	arr := make([]int, 0)
	curNode := root
	for curNode != nil {
		if curNode.left != nil {
			rightMost := curNode.left
			for rightMost.right != nil && rightMost.right != curNode {
				rightMost = rightMost.right
			}
			if rightMost.right == curNode {
				rightMost.right = nil
				arr = append(arr, curNode.val)
				curNode = curNode.right
			} else {
				rightMost.right = curNode
				curNode = curNode.left
			}
		} else {
			arr = append(arr, curNode.val)
			curNode = curNode.right
		}
	}
	return arr
}

func morrisPreorder(root *treeNode) []int {
	arr := make([]int, 0)
	curNode := root
	for curNode != nil {
		if curNode.left != nil {
			rightMost := curNode.left
			for rightMost.right != nil && rightMost.right != curNode {
				rightMost = rightMost.right
			}
			if rightMost.right == curNode {
				rightMost.right = nil
				curNode = curNode.right
			} else {
				arr = append(arr, curNode.val)
				rightMost.right = curNode
				curNode = curNode.left
			}
		} else {
			arr = append(arr, curNode.val)
			curNode = curNode.right
		}
	}
	return arr
}

func main() {
	root := &treeNode{10, nil, nil, 0}
	root = root.insert(5)
	root = root.insert(20)
	root = root.insert(1)
	root = root.insert(7)
	root = root.insert(15)
	root = root.insert(16)
	root = root.insert(30)
	root.inorder()

	fmt.Printf("3rd smallest: %d\n", ksmallestBST(root, 3))
	fmt.Printf("6th smallest: %d\n", ksmallestBST(root, 6))

	fmt.Printf("14's ceil value is %d\n", root.ceil(14))
	fmt.Printf("0's ceil value is %d\n", root.ceil(0))
	fmt.Printf("30's ceil value is %d\n", root.ceil(30))
	fmt.Printf("40's ceil value is %d\n", root.ceil(40))

	fmt.Printf("tree is complete? %t\n", root.isComplete())

	root = root.insert(13)
	root = root.insert(0)
	root = root.insert(4)
	root = root.insert(6)
	root = root.insert(9)
	fmt.Printf("tree is complete? %t\n", root.isComplete())

	root.inorder()
	root.preorder()

	fmt.Printf("tree's morris inorder traversal is: %v\n", morrisInorder(root))
	fmt.Printf("tree's morris preorder traversal is: %v\n", morrisPreorder(root))
}
