package tree

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root   *Node
	nCount int
}

func GenerateTree(count int) *AVLTree {
	t := &AVLTree{nCount: 0}
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		key := rand.Intn(1000)
		t.Insert(key)
	}
	return t
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
	t.nCount++
}

func (t *AVLTree) ToMermaid() string {
	var tree string
	stack := []*Node{t.Root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Left != nil {
			tree += fmt.Sprintf("%v--> %v\n", current.Key, current.Left.Key)
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			tree += fmt.Sprintf("%v--> %v\n", current.Key, current.Right.Key)
			stack = append(stack, current.Right)
		}
	}

	return "{{< mermaid >}}\ngraph TD\n" + tree + "\n{{< /mermaid >}}"
}

func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1
	return y
}

func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1
	return x
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	balanceFactor := getBalance(node)

	if balanceFactor > 1 {
		if key < node.Left.Key {
			return rightRotate(node)
		} else if key > node.Left.Key {
			node.Left = leftRotate(node.Left)
			return rightRotate(node)
		}
	}

	if balanceFactor < -1 {
		if key > node.Right.Key {
			return leftRotate(node)
		} else if key < node.Right.Key {
			node.Right = rightRotate(node.Right)
			leftRotate(node)
		}
	}

	return node
}
