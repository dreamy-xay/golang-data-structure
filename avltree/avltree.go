/*
 * @Description: avltree
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-13 20:05:07
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-20 17:17:34
 */
package avltree

import (
	"fmt"
)

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type AVLTreeNode struct {
	Value  interface{}
	height int
	count  int
	left   *AVLTreeNode
	right  *AVLTreeNode
	parent *AVLTreeNode
}

func (avlnode *AVLTreeNode) GetHeight() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.height
}

func (avlnode *AVLTreeNode) GetCount() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.count
}

func (avlnode *AVLTreeNode) GetPre() *AVLTreeNode {
	if avlnode.parent == nil {
		rightnode := avlnode.left
		for rightnode.right != nil {
			rightnode = rightnode.right
		}
		return rightnode
	} else if avlnode.parent.right == avlnode {
		if avlnode.left != nil {
			return avlnode.left
		}
		return avlnode.parent
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

func (avlnode *AVLTreeNode) GetNext() *AVLTreeNode {
	if avlnode.parent == nil {
		leftnode := avlnode.right
		for leftnode.left != nil {
			leftnode = leftnode.left
		}
		return leftnode
	} else if avlnode.parent.left == avlnode {
		if avlnode.right != nil {
			return avlnode.right
		}
		return avlnode.parent
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

func (avlnode *AVLTreeNode) getBalanceFactor() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.left.GetHeight() - avlnode.right.GetHeight()
}

func (avlnode *AVLTreeNode) inOrder(fn func(value *interface{})) {
	if avlnode == nil {
		return
	}
	avlnode.left.inOrder(fn)
	fn(&avlnode.Value)
	avlnode.right.inOrder(fn)
}

func (avlnode *AVLTreeNode) findNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *AVLTreeNode {
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

func (avlnode *AVLTreeNode) orderNode(value *interface{}, rules, equal func(a, b *interface{}) bool) int {
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

func (avlnode *AVLTreeNode) atNode(index, preCount int) *AVLTreeNode {
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

func (avlnode *AVLTreeNode) lowerBoundNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *AVLTreeNode {
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

func (avlnode *AVLTreeNode) upperBoundNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *AVLTreeNode {
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

func (avlnode *AVLTreeNode) leftRotate() *AVLTreeNode {
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

func (avlnode *AVLTreeNode) rightRotate() *AVLTreeNode {
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

func (avlnode *AVLTreeNode) balanceNode() *AVLTreeNode {
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

func (avlnode *AVLTreeNode) insertNode(node, parent *AVLTreeNode, success *bool, rules, equal func(a, b *interface{}) bool) *AVLTreeNode {
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

func (avlnode *AVLTreeNode) eraseNode(value *interface{}, success *bool, rules, equal func(a, b *interface{}) bool) *AVLTreeNode {
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

type AVLTree struct {
	root  *AVLTreeNode
	rules func(a, b *interface{}) bool
	equal func(a, b *interface{}) bool
}

func NewAVLTree(rules, equal func(a, b *interface{}) bool) AVLTree {
	return AVLTree{
		root:  nil,
		rules: rules,
		equal: equal,
	}
}

func (avl *AVLTree) Init(arr []interface{}, rules, equal func(a, b *interface{}) bool) {
	avl.Clear()
	avl.rules = rules
	avl.equal = equal
	for _, elem := range arr {
		avl.Insert(elem)
	}
}

func (avl *AVLTree) Size() int {
	return avl.root.GetCount()
}

func (avl *AVLTree) ForSize() int {
	size := 0
	avl.root.inOrder(func(_ *interface{}) {
		size += 1
	})
	return size
}

func (avl *AVLTree) Empty() bool {
	return avl.root == nil
}

func (avl *AVLTree) Height() int {
	return avl.root.GetHeight()
}

func (avl *AVLTree) Insert(value interface{}) *AVLTreeNode {
	if avl.root == nil {
		avl.root = &AVLTreeNode{
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

func (avl *AVLTree) Erase(value interface{}) {
	avl.EraseByNode(&avl.root, value)
}

func (avl *AVLTree) Find(value interface{}) (*AVLTreeNode, bool) {
	return avl.FindByNode(avl.root, value)
}

func (avl *AVLTree) Order(value interface{}) int {
	return avl.OrderByNode(avl.root, value)
}

func (avl *AVLTree) At(index int) *AVLTreeNode {
	return avl.AtByNode(avl.root, index)
}

func (avl *AVLTree) LowerBound(value interface{}) *AVLTreeNode {
	return avl.LowerBoundByNode(avl.root, value)
}

func (avl *AVLTree) UpperBound(value interface{}) *AVLTreeNode {
	return avl.UpperBoundByNode(avl.root, value)
}

func (avl *AVLTree) Begin() *AVLTreeNode {
	if avl.root == nil {
		return nil
	}
	node := avl.root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (avl *AVLTree) End() *AVLTreeNode {
	if avl.root == nil {
		return nil
	}
	node := avl.root
	for node.right != nil {
		node = node.right
	}
	return node
}

func (avl *AVLTree) Front() (interface{}, bool) {
	if avl.root == nil {
		return nil, false
	}
	node := avl.root
	for node.left != nil {
		node = node.left
	}
	return node.Value, true
}

func (avl *AVLTree) Back() (interface{}, bool) {
	if avl.root == nil {
		return nil, false
	}
	node := avl.root
	for node.right != nil {
		node = node.right
	}
	return node.Value, true
}

func (avl *AVLTree) Clear() {
	avl.ClearByNode(&avl.root)
}

func (avl *AVLTree) IsAVLTree() bool {
	return avl.IsAVLTreeByNode(avl.root)
}

func (avl *AVLTree) ForEach(fn func(value *interface{})) {
	avl.root.inOrder(fn)
}

func (avl *AVLTree) String() string {
	str := "AVLTree {\n"
	avl.root.inOrder(func(value *interface{}) {
		str += "    " + fmt.Sprint(*value) + ",\n"
	})
	str += "}"
	return str
}

func (avl *AVLTree) IsAVLTreeByNode(avlnode *AVLTreeNode) bool {
	if avlnode == nil {
		return true
	}
	if Abs(avlnode.getBalanceFactor()) > 1 || !((avlnode.left == nil || avl.rules(&avlnode.left.Value, &avlnode.Value)) && (avlnode.right == nil || !avl.rules(&avlnode.right.Value, &avlnode.Value))) {
		return false
	}
	return avl.IsAVLTreeByNode(avlnode.left) && avl.IsAVLTreeByNode(avlnode.right)
}

func (avl *AVLTree) FindByNode(avlnode *AVLTreeNode, value interface{}) (*AVLTreeNode, bool) {
	if avlnode == nil {
		return nil, false
	}
	node := avlnode.findNode(&value, avl.rules, avl.equal)
	if node == nil {
		return nil, false
	}
	return node, true
}

func (avl *AVLTree) OrderByNode(avlnode *AVLTreeNode, value interface{}) int {
	if avlnode == nil {
		return -1
	}
	return avlnode.orderNode(&value, avl.rules, avl.equal) - 1
}

func (avl *AVLTree) AtByNode(avlnode *AVLTreeNode, index int) *AVLTreeNode {
	if avlnode == nil || index < 0 || index >= avl.Size() {
		return nil
	}
	return avlnode.atNode(index, 0)
}

func (avl *AVLTree) LowerBoundByNode(avlnode *AVLTreeNode, value interface{}) *AVLTreeNode {
	if avlnode == nil {
		return nil
	}
	return avlnode.lowerBoundNode(&value, avl.rules, avl.equal)
}

func (avl *AVLTree) UpperBoundByNode(avlnode *AVLTreeNode, value interface{}) *AVLTreeNode {
	if avlnode == nil {
		return nil
	}
	return avlnode.upperBoundNode(&value, avl.rules, avl.equal)
}

func (avl *AVLTree) InsertByNode(avlnode **AVLTreeNode, value interface{}) *AVLTreeNode {
	if *avlnode == nil {
		return nil
	}
	node := &AVLTreeNode{
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

func (avl *AVLTree) EraseByNode(avlnode **AVLTreeNode, value interface{}) {
	if *avlnode == nil {
		return
	}
	success := false
	*avlnode = (*avlnode).eraseNode(&value, &success, avl.rules, avl.equal)
}

func (avl *AVLTree) ClearByNode(avlnode **AVLTreeNode) {
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
