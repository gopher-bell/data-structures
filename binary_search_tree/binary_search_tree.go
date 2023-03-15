package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

func (n *Node) PreOrder() {
	if n != nil {
		fmt.Println(n.Key)
		n.Left.PreOrder()
		n.Right.PreOrder()
	}

}

func main() {
	tree := &Node{Key: 200}
	tree.Insert(100)
	tree.Insert(500)
	tree.Insert(300)
	tree.Insert(10)
	tree.Insert(1)
	tree.PreOrder()

	fmt.Println(tree.Search(1))
	fmt.Println(tree.Search(55))
}
