/*
 * @Description: avltree
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-13 20:05:07
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 15:11:22
 */
package avltree

import (
	. "collections/common"
	"fmt"
)

type AVLTreeNode[T any] struct {
	Value  T
	height int
	count  int
	left   *AVLTreeNode[T]
	right  *AVLTreeNode[T]
	parent *AVLTreeNode[T]
}

type AVLTreeIterator[T any] *AVLTreeNode[T]

func (avlnode *AVLTreeNode[T]) GetHeight() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.height
}

func (avlnode *AVLTreeNode[T]) GetCount() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.count
}

func (avlnode *AVLTreeNode[T]) GetPre() *AVLTreeNode[T] {
	if avlnode.parent == nil {
		if avlnode.left == nil {
			return nil
		}
		rightnode := avlnode.left
		for rightnode.right != nil {
			rightnode = rightnode.right
		}
		return rightnode
	} else if avlnode.parent.right == avlnode {
		if avlnode.left == nil {
			return avlnode.parent
		}
		rightnode := avlnode.left
		for rightnode.right != nil {
			rightnode = rightnode.right
		}
		return rightnode
	} else {
		if avlnode.left != nil {
			rightnode := avlnode.left
			for rightnode.right != nil {
				rightnode = rightnode.right
			}
			return rightnode
		}
		curnode := avlnode
		parentnode := avlnode.parent
		for parentnode.left == curnode {
			curnode = parentnode
			parentnode = parentnode.parent
			if parentnode == nil {
				return nil
			}
		}
		return parentnode
	}
}

func (avlnode *AVLTreeNode[T]) GetNext() *AVLTreeNode[T] {
	if avlnode.parent == nil {
		if avlnode.right == nil {
			return nil
		}
		leftnode := avlnode.right
		for leftnode.left != nil {
			leftnode = leftnode.left
		}
		return leftnode
	} else if avlnode.parent.left == avlnode {
		if avlnode.right == nil {
			return avlnode.parent
		}
		leftnode := avlnode.right
		for leftnode.left != nil {
			leftnode = leftnode.left
		}
		return leftnode
	} else {
		if avlnode.right != nil {
			leftnode := avlnode.right
			for leftnode.left != nil {
				leftnode = leftnode.left
			}
			return leftnode
		}

		curnode := avlnode
		parentnode := avlnode.parent
		for parentnode.right == curnode {
			curnode = parentnode
			parentnode = parentnode.parent
			if parentnode == nil {
				return nil
			}
		}
		return parentnode
	}
}

func (avlnode *AVLTreeNode[T]) getBalanceFactor() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.left.GetHeight() - avlnode.right.GetHeight()
}

func (avlnode *AVLTreeNode[T]) inOrder(fn func(value *T)) {
	if avlnode == nil {
		return
	}
	avlnode.left.inOrder(fn)
	fn(&avlnode.Value)
	avlnode.right.inOrder(fn)
}

func (avlnode *AVLTreeNode[T]) findNode(value *T, rules, equal func(a, b *T) bool) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	if equal(value, &avlnode.Value) {
		return avlnode
	} else if rules(value, &avlnode.Value) {
		return avlnode.left.findNode(value, rules, equal)
	} else {
		return avlnode.right.findNode(value, rules, equal)
	}
}

func (avlnode *AVLTreeNode[T]) orderNode(value *T, rules, equal func(a, b *T) bool) int {
	if avlnode == nil {
		return 0
	}
	if equal(value, &avlnode.Value) {
		return avlnode.left.GetCount() + 1
	} else if rules(value, &avlnode.Value) {
		return avlnode.left.orderNode(value, rules, equal)
	} else {
		count := avlnode.right.orderNode(value, rules, equal)
		if count > 0 {
			return count + avlnode.left.GetCount() + 1
		}
		return 0
	}
}

func (avlnode *AVLTreeNode[T]) atNode(index, preCount int) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	count := avlnode.left.GetCount() + preCount
	if index == count {
		return avlnode
	} else if index < count {
		return avlnode.left.atNode(index, preCount)
	} else {
		return avlnode.right.atNode(index, count+1)
	}
}

func (avlnode *AVLTreeNode[T]) lowerBoundNode(value *T, rules, equal func(a, b *T) bool) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	if equal(value, &avlnode.Value) {
		return avlnode
	} else if rules(value, &avlnode.Value) {
		node := avlnode.left.lowerBoundNode(value, rules, equal)
		if node == nil {
			return avlnode
		}
		return node
	} else {
		return avlnode.right.lowerBoundNode(value, rules, equal)
	}
}

func (avlnode *AVLTreeNode[T]) upperBoundNode(value *T, rules, equal func(a, b *T) bool) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	if equal(value, &avlnode.Value) {
		leftnode := avlnode.right
		if leftnode == nil {
			return nil
		}
		for leftnode.left != nil {
			leftnode = leftnode.left
		}
		return leftnode
	} else if rules(value, &avlnode.Value) {
		node := avlnode.left.upperBoundNode(value, rules, equal)
		if node == nil {
			return avlnode
		}
		return node
	} else {
		return avlnode.right.upperBoundNode(value, rules, equal)
	}
}

func (avlnode *AVLTreeNode[T]) leftRotate() *AVLTreeNode[T] {
	/*
		 avlnode
		 /    \
		t4    mid
			 /  \
			t3  right
				 / \
			    t2 t1
	*/
	// update left right
	mid := avlnode.right
	t3 := mid.left
	mid.left = avlnode
	avlnode.right = t3

	// uodate parent
	mid.parent = avlnode.parent
	avlnode.parent = mid
	if t3 != nil {
		t3.parent = avlnode
	}

	// update count
	mid.count = avlnode.count
	avlnode.count = avlnode.left.GetCount() + avlnode.right.GetCount() + 1

	// update height
	avlnode.height = Max(avlnode.left.GetHeight(), avlnode.right.GetHeight()) + 1
	mid.height = Max(mid.left.GetHeight(), mid.right.GetHeight()) + 1
	return mid
}

func (avlnode *AVLTreeNode[T]) rightRotate() *AVLTreeNode[T] {
	/*
			   avlnode
			   /     \
			  mid    t4
			 /  \
		   left t3
		   / \
		  t1 t2
	*/
	// update left right
	mid := avlnode.left
	t3 := mid.right
	mid.right = avlnode
	avlnode.left = t3

	// uodate parent
	mid.parent = avlnode.parent
	avlnode.parent = mid
	if t3 != nil {
		t3.parent = avlnode
	}

	// update count
	mid.count = avlnode.count
	avlnode.count = avlnode.left.GetCount() + avlnode.right.GetCount() + 1

	// update height
	avlnode.height = Max(avlnode.left.GetHeight(), avlnode.right.GetHeight()) + 1
	mid.height = Max(mid.left.GetHeight(), mid.right.GetHeight()) + 1
	return mid
}

func (avlnode *AVLTreeNode[T]) balanceNode() *AVLTreeNode[T] {
	// update height
	avlnode.height = Max(avlnode.left.GetHeight(), avlnode.right.GetHeight()) + 1

	// update count
	avlnode.count = avlnode.left.GetCount() + avlnode.right.GetCount() + 1

	// rotate balance
	balanceFactor := avlnode.getBalanceFactor()
	if balanceFactor > 1 {
		leftBalanceFactor := avlnode.left.getBalanceFactor()
		if leftBalanceFactor > 0 {
			return avlnode.rightRotate()
		}
		if leftBalanceFactor < 0 {
			avlnode.left = avlnode.left.leftRotate()
			return avlnode.rightRotate()
		}
	}
	if balanceFactor < -1 {
		rightBalanceFactor := avlnode.right.getBalanceFactor()
		if rightBalanceFactor < 0 {
			return avlnode.leftRotate()
		}
		if rightBalanceFactor > 0 {
			avlnode.right = avlnode.right.rightRotate()
			return avlnode.leftRotate()
		}
	}
	return avlnode
}

func (avlnode *AVLTreeNode[T]) insertNode(node, parent *AVLTreeNode[T], success *bool, rules, equal func(a, b *T) bool) *AVLTreeNode[T] {
	// insert
	if avlnode == nil {
		*success = true
		node.parent = parent
		return node
	}
	if equal(&node.Value, &avlnode.Value) {
		return avlnode
	} else if rules(&node.Value, &avlnode.Value) {
		avlnode.left = avlnode.left.insertNode(node, avlnode, success, rules, equal)
	} else {
		avlnode.right = avlnode.right.insertNode(node, avlnode, success, rules, equal)
	}

	// update height and balance
	if *success {
		return avlnode.balanceNode()
	} else {
		return avlnode
	}
}

func (avlnode *AVLTreeNode[T]) eraseNode(value *T, success *bool, rules, equal func(a, b *T) bool) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	node := avlnode
	if equal(value, &avlnode.Value) {
		*success = true
		if avlnode.left == nil {
			node = avlnode.right
			if node != nil {
				node.parent = avlnode.parent
			}
			avlnode.right = nil
			avlnode.parent = nil
			avlnode.height = 1
			avlnode.count = 1
		} else if avlnode.right == nil {
			node = avlnode.left
			if node != nil {
				node.parent = avlnode.parent
			}
			avlnode.left = nil
			avlnode.parent = nil
			avlnode.height = 1
			avlnode.count = 1
		} else {
			node = avlnode.right
			for node.left != nil {
				node = node.left
			}
			node.right = avlnode.right.eraseNode(&node.Value, success, rules, equal)
			node.left = avlnode.left
			node.parent = avlnode.parent
			node.left.parent = node
			if node.right != nil {
				node.right.parent = node
			}
			avlnode.left = nil
			avlnode.right = nil
			avlnode.parent = nil
			avlnode.height = 1
			avlnode.count = 1
		}
	} else if rules(value, &avlnode.Value) {
		avlnode.left = avlnode.left.eraseNode(value, success, rules, equal)
	} else {
		avlnode.right = avlnode.right.eraseNode(value, success, rules, equal)
	}

	if node == nil {
		return nil
	}
	// update height and balance
	if *success {
		return node.balanceNode()
	} else {
		return node
	}
}

type AVLTree[T any] struct {
	root  *AVLTreeNode[T]
	rules func(a, b *T) bool
	equal func(a, b *T) bool
}

func NewAVLTree[T any](rules, equal func(a, b *T) bool) *AVLTree[T] {
	return &AVLTree[T]{
		root:  nil,
		rules: rules,
		equal: equal,
	}
}

func (avl *AVLTree[T]) Init(arr []T, rules, equal func(a, b *T) bool) {
	avl.Clear()
	avl.rules = rules
	avl.equal = equal
	for _, elem := range arr {
		avl.Insert(elem)
	}
}

func (avl *AVLTree[T]) Size() int {
	return avl.root.GetCount()
}

func (avl *AVLTree[T]) ForSize() int {
	size := 0
	avl.root.inOrder(func(_ *T) {
		size += 1
	})
	return size
}

func (avl *AVLTree[T]) Empty() bool {
	return avl.root == nil
}

func (avl *AVLTree[T]) Height() int {
	return avl.root.GetHeight()
}

func (avl *AVLTree[T]) Insert(value T) *AVLTreeNode[T] {
	if avl.root == nil {
		avl.root = &AVLTreeNode[T]{
			Value:  value,
			height: 1,
			count:  1,
			parent: nil,
			left:   nil,
			right:  nil,
		}
		return avl.root
	}
	return avl.InsertByNode(&avl.root, value)
}

func (avl *AVLTree[T]) Erase(value T) {
	avl.EraseByNode(&avl.root, value)
}

func (avl *AVLTree[T]) Find(value T) (*AVLTreeNode[T], bool) {
	return avl.FindByNode(avl.root, value)
}

func (avl *AVLTree[T]) Contains(value T) bool {
	_, ok := avl.FindByNode(avl.root, value)
	return ok
}

func (avl *AVLTree[T]) Order(value T) int {
	return avl.OrderByNode(avl.root, value)
}

func (avl *AVLTree[T]) At(index int) *AVLTreeNode[T] {
	return avl.AtByNode(avl.root, index)
}

func (avl *AVLTree[T]) LowerBound(value T) *AVLTreeNode[T] {
	return avl.LowerBoundByNode(avl.root, value)
}

func (avl *AVLTree[T]) UpperBound(value T) *AVLTreeNode[T] {
	return avl.UpperBoundByNode(avl.root, value)
}

func (avl *AVLTree[T]) Begin() *AVLTreeNode[T] {
	if avl.root == nil {
		return nil
	}
	node := avl.root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (avl *AVLTree[T]) End() *AVLTreeNode[T] {
	if avl.root == nil {
		return nil
	}
	node := avl.root
	for node.right != nil {
		node = node.right
	}
	return node
}

func (avl *AVLTree[T]) Front() (any, bool) {
	if avl.root == nil {
		return nil, false
	}
	node := avl.root
	for node.left != nil {
		node = node.left
	}
	return node.Value, true
}

func (avl *AVLTree[T]) Back() (any, bool) {
	if avl.root == nil {
		return nil, false
	}
	node := avl.root
	for node.right != nil {
		node = node.right
	}
	return node.Value, true
}

func (avl *AVLTree[T]) Clear() {
	avl.ClearByNode(&avl.root)
}

func (avl *AVLTree[T]) IsAVLTree() bool {
	return avl.IsAVLTreeByNode(avl.root)
}

func (avl *AVLTree[T]) ForEach(fn func(value *T)) {
	avl.root.inOrder(fn)
}

func (avl *AVLTree[T]) GetIterator() *Iterator[T] {
	current := avl.Begin()
	isIterated := false
	return NewIterator[T](func() bool {
		return current.GetNext() != nil
	}, func() *T {
		if isIterated {
			current = current.GetNext()
		} else {
			isIterated = true
		}
		return &current.Value
	})
}

func (avl *AVLTree[T]) String() string {
	str := "AVLTree {\n"
	avl.root.inOrder(func(value *T) {
		str += "    " + fmt.Sprint(*value) + ",\n"
	})
	str += "}"
	return str
}

func (avl *AVLTree[T]) IsAVLTreeByNode(avlnode *AVLTreeNode[T]) bool {
	if avlnode == nil {
		return true
	}
	if Abs(avlnode.getBalanceFactor()) > 1 || !((avlnode.left == nil || avl.rules(&avlnode.left.Value, &avlnode.Value)) && (avlnode.right == nil || !avl.rules(&avlnode.right.Value, &avlnode.Value))) {
		return false
	}
	return avl.IsAVLTreeByNode(avlnode.left) && avl.IsAVLTreeByNode(avlnode.right)
}

func (avl *AVLTree[T]) FindByNode(avlnode *AVLTreeNode[T], value T) (*AVLTreeNode[T], bool) {
	if avlnode == nil {
		return nil, false
	}
	node := avlnode.findNode(&value, avl.rules, avl.equal)
	if node == nil {
		return nil, false
	}
	return node, true
}

func (avl *AVLTree[T]) OrderByNode(avlnode *AVLTreeNode[T], value T) int {
	if avlnode == nil {
		return -1
	}
	return avlnode.orderNode(&value, avl.rules, avl.equal) - 1
}

func (avl *AVLTree[T]) AtByNode(avlnode *AVLTreeNode[T], index int) *AVLTreeNode[T] {
	if avlnode == nil || index < 0 || index >= avl.Size() {
		return nil
	}
	return avlnode.atNode(index, 0)
}

func (avl *AVLTree[T]) LowerBoundByNode(avlnode *AVLTreeNode[T], value T) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	return avlnode.lowerBoundNode(&value, avl.rules, avl.equal)
}

func (avl *AVLTree[T]) UpperBoundByNode(avlnode *AVLTreeNode[T], value T) *AVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	return avlnode.upperBoundNode(&value, avl.rules, avl.equal)
}

func (avl *AVLTree[T]) InsertByNode(avlnode **AVLTreeNode[T], value T) *AVLTreeNode[T] {
	if *avlnode == nil {
		return nil
	}
	node := &AVLTreeNode[T]{
		Value:  value,
		height: 1,
		count:  1,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	success := false
	*avlnode = (*avlnode).insertNode(node, (*avlnode).parent, &success, avl.rules, avl.equal)
	if success {
		return node
	}
	return nil
}

func (avl *AVLTree[T]) EraseByNode(avlnode **AVLTreeNode[T], value T) {
	if *avlnode == nil {
		return
	}
	success := false
	*avlnode = (*avlnode).eraseNode(&value, &success, avl.rules, avl.equal)
}

func (avl *AVLTree[T]) ClearByNode(avlnode **AVLTreeNode[T]) {
	if *avlnode == nil {
		return
	}
	avl.ClearByNode(&(*avlnode).left)
	avl.ClearByNode(&(*avlnode).right)
	(*avlnode).height = 1
	(*avlnode).count = 1
	(*avlnode).left = nil
	(*avlnode).right = nil
	(*avlnode).parent = nil
	*avlnode = nil
}
