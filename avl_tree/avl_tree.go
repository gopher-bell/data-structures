package main

import (
	"fmt"
)

type node struct {
	value       int
	left, right *node
	height      int
}

func getHeight(n *node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func getBalance(n *node) int {
	if n == nil {
		return 0
	}
	return getHeight(n.left) - getHeight(n.right)
}

func rotateRight(n *node) *node {
	x := n.left
	y := x.right
	x.right = n
	n.left = y

	n.height = 1 + max(getHeight(n.left), getHeight(n.right))
	x.height = 1 + max(getHeight(x.left), getHeight(x.right))

	return x
}

func rotateLeft(n *node) *node {
	x := n.right
	y := x.left
	x.left = n
	n.right = y

	n.height = 1 + max(getHeight(n.left), getHeight(n.right))
	x.height = 1 + max(getHeight(x.left), getHeight(x.right))

	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insertNode(n *node, value int) *node {
	if n == nil {
		return &node{value: value, height: 1}
	}

	if value < n.value {
		n.left = insertNode(n.left, value)
	} else {
		n.right = insertNode(n.right, value)
	}

	n.height = 1 + max(getHeight(n.left), getHeight(n.right))
	balance := getBalance(n)

	if balance > 1 && value < n.left.value {
		return rotateRight(n)
	}

	if balance < -1 && value > n.right.value {
		return rotateLeft(n)
	}

	if balance > 1 && value > n.left.value {
		n.left = rotateLeft(n.left)
		return rotateRight(n)
	}

	if balance < -1 && value < n.right.value {
		n.right = rotateRight(n.right)
		return rotateLeft(n)
	}

	return n
}

func traverseInOrder(n *node) {
	if n != nil {
		traverseInOrder(n.left)
		fmt.Printf("%d ", n.value)
		traverseInOrder(n.right)
	}
}

func main() {
	var root *node
	root = insertNode(root, 10)
	root = insertNode(root, 20)
	root = insertNode(root, 30)
	root = insertNode(root, 40)
	root = insertNode(root, 50)
	root = insertNode(root, 25)

	fmt.Println("In-order traversal of AVL tree:")
	traverseInOrder(root)
}
