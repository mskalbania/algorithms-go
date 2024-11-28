package other

import "cmp"

type BinaryTree[T cmp.Ordered] struct {
	root *Node[T]
}

type Node[T cmp.Ordered] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func (t *BinaryTree[T]) Insert(value T) {
	t.root = insertNode(t.root, value)
}

func insertNode[T cmp.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{data: value}
	}
	if value < node.data {
		node.left = insertNode(node.left, value)
	} else {
		node.right = insertNode(node.right, value)
	}
	return node
}

func (t *BinaryTree[T]) Search(value T) bool {
	return searchNode(t.root, value)
}

func searchNode[T cmp.Ordered](node *Node[T], value T) bool {
	if node == nil {
		return false
	}
	if value == node.data {
		return true
	}
	if value < node.data {
		return searchNode(node.left, value)
	}
	return searchNode(node.right, value)
}

func (t *BinaryTree[T]) Delete(value T) {
	t.root = deleteNode(t.root, value)
}

func deleteNode[T cmp.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}
	if value == node.data {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		node.data = minValue(node.right)
		node.right = deleteNode(node.right, node.data)
	} else if value < node.data {
		node.left = deleteNode(node.left, value)
	} else {
		node.right = deleteNode(node.right, value)
	}
	return node
}

func minValue[T cmp.Ordered](node *Node[T]) T {
	for node.left != nil {
		node = node.left
	}
	return node.data
}
