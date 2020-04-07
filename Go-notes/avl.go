// src: https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
package main

import (
	"fmt"
	"math/rand"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type node struct {
	val, height int
	left, right *node
}

type AVLTree struct {
	root *node
}

func (at *AVLTree) Insert(val int) {
	at.root = insertR(at.root, val)
}

func getMaxHeight(root *node) int {
	height := 0
	if root.left != nil {
		height = max(height, root.left.height)
	}
	if root.right != nil {
		height = max(height, root.right.height)
	}
	return height
}

func getBalance(root *node) int {
	if root == nil {
		return 0
	}

	balance := 0
	if root.left != nil {
		balance -= root.left.height
	}
	if root.right != nil {
		balance += root.right.height
	}
	return balance
}

func leftRotate(x *node) *node {
	y := x.right
	T2 := y.left

	// Perform rotation
	y.left = x
	x.right = T2

	// Update heights
	x.height = getMaxHeight(x) + 1
	y.height = getMaxHeight(y) + 1

	return y
}

func rightRotate(y *node) *node {
	x := y.left
	T2 := x.right

	// Perform rotation
	x.right = y
	y.left = T2

	// Update heights
	x.height = getMaxHeight(x) + 1
	y.height = getMaxHeight(y) + 1

	// Return new root
	return x
}

func insertR(root *node, val int) *node {
	if root == nil {
		return &node{val, 1, nil, nil}
	} else if val == root.val {
		return root
	} else if val < root.val {
		root.left = insertR(root.left, val)
	} else if val > root.val {
		root.right = insertR(root.right, val)
	}

	root.height = getMaxHeight(root) + 1

	balance := getBalance(root)

	if balance > 1 && root.left != nil && val < root.left.val { // Left Left Case
		return rightRotate(root)
	}

	if balance < -1 && root.right != nil && val > root.right.val { // Right Right Case
		return leftRotate(root)
	}

	if balance > 1 && root.left != nil && val > root.left.val { // Left Right Case
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if balance < -1 && root.right != nil && val < root.right.val { // Right Left Case
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

func (at *AVLTree) Search(val int) bool {
	tmp := at.root
	for tmp != nil {
		if tmp.val == val {
			return true
		} else if val < tmp.val {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	return false
}

func (at *AVLTree) InOrder() {
	inOrder(at.root)
}

func inOrder(root *node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.val)
	inOrder(root.right)
}

func shuffle(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	vals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffle(vals)
	fmt.Println(vals)

	tree := AVLTree{}

	for i := 0; i <= 10; i++ {
		tree.Insert(i)
	}

	tree.InOrder()
}
