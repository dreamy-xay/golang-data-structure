/*
 * @Description: RBTree
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-16 21:43:01
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-08-07 14:10:32
 */
package rbtree

import "fmt"

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

func (color *Color) String() string {
	if *color == RED {
		return "RED"
	} else {
		return "BLACK"
	}
}

type RBTreeNode[T any] struct {
	Value  T
	color  Color
	left   *RBTreeNode[T]
	right  *RBTreeNode[T]
	parent *RBTreeNode[T]
}

func (rbnode *RBTreeNode[T]) GetColor() Color {
	if rbnode == nil {
		return BLACK
	}
	return rbnode.color
}

func (rbnode *RBTreeNode[T]) GetPre() *RBTreeNode[T] {
	if rbnode.parent == nil {
		rightnode := rbnode.left
		for rightnode.right != nil {
			rightnode = rightnode.right
		}
		return rightnode
	} else if rbnode.parent.right == rbnode {
		if rbnode.left != nil {
			return rbnode.left
		}
		return rbnode.parent
	} else {
		if rbnode.left != nil {
			rightnode := rbnode.left
			for rightnode.right != nil {
				rightnode = rightnode.right
			}
			return rightnode
		}
		curnode := rbnode
		parentnode := rbnode.parent
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

func (rbnode *RBTreeNode[T]) GetNext() *RBTreeNode[T] {
	if rbnode.parent == nil {
		leftnode := rbnode.right
		for leftnode.left != nil {
			leftnode = leftnode.left
		}
		return leftnode
	} else if rbnode.parent.left == rbnode {
		if rbnode.right != nil {
			return rbnode.right
		}
		return rbnode.parent
	} else {
		if rbnode.right != nil {
			leftnode := rbnode.right
			for leftnode.left != nil {
				leftnode = leftnode.left
			}
			return leftnode
		}
		curnode := rbnode
		parentnode := rbnode.parent
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

func (rbnode *RBTreeNode[T]) inOrder(fn func(value *T)) {
	if rbnode == nil {
		return
	}
	rbnode.left.inOrder(fn)
	fn(&rbnode.Value)
	rbnode.right.inOrder(fn)
}

func (rbnode *RBTreeNode[T]) findNode(value *T, rules, equal func(a, b *T) bool) *RBTreeNode[T] {
	if rbnode == nil {
		return nil
	}
	if equal(value, &rbnode.Value) {
		return rbnode
	} else if rules(value, &rbnode.Value) {
		return rbnode.left.findNode(value, rules, equal)
	} else {
		return rbnode.right.findNode(value, rules, equal)
	}
}

func (rbnode *RBTreeNode[T]) lowerBoundNode(value *T, rules, equal func(a, b *T) bool) *RBTreeNode[T] {
	if rbnode == nil {
		return nil
	}
	if equal(value, &rbnode.Value) {
		return rbnode
	} else if rules(value, &rbnode.Value) {
		node := rbnode.left.lowerBoundNode(value, rules, equal)
		if node == nil {
			return rbnode
		}
		return node
	} else {
		return rbnode.right.lowerBoundNode(value, rules, equal)
	}
}

func (rbnode *RBTreeNode[T]) upperBoundNode(value *T, rules, equal func(a, b *T) bool) *RBTreeNode[T] {
	if rbnode == nil {
		return nil
	}
	if equal(value, &rbnode.Value) {
		leftnode := rbnode.right
		if leftnode == nil {
			return nil
		}
		for leftnode.left != nil {
			leftnode = leftnode.left
		}
		return leftnode
	} else if rules(value, &rbnode.Value) {
		node := rbnode.left.upperBoundNode(value, rules, equal)
		if node == nil {
			return rbnode
		}
		return node
	} else {
		return rbnode.right.upperBoundNode(value, rules, equal)
	}
}

func (rbnode *RBTreeNode[T]) leftRotate() *RBTreeNode[T] {
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
	mid := rbnode.right
	t3 := mid.left
	mid.left = rbnode
	rbnode.right = t3

	// uodate parent
	mid.parent = rbnode.parent
	rbnode.parent = mid
	if t3 != nil {
		t3.parent = rbnode
	}

	return mid
}

func (rbnode *RBTreeNode[T]) rightRotate() *RBTreeNode[T] {
	/*
			    rbnode
			   /     \
			  mid    t4
			 /  \
		   left t3
		   / \
		  t1 t2
	*/
	// update left right
	mid := rbnode.left
	t3 := mid.right
	mid.right = rbnode
	rbnode.left = t3

	// uodate parent
	mid.parent = rbnode.parent
	rbnode.parent = mid
	if t3 != nil {
		t3.parent = rbnode
	}

	return mid
}

func (rbnode *RBTreeNode[T]) insertNode(value *T, root **RBTreeNode[T], rules, equal func(a, b *T) bool) *RBTreeNode[T] {
	// define
	var (
		pnode   *RBTreeNode[T] = nil
		curnode *RBTreeNode[T] = rbnode
		node    *RBTreeNode[T] = nil
	)

	// find position
	for curnode != nil {
		if equal(value, &curnode.Value) {
			return node
		} else if rules(value, &curnode.Value) {
			pnode = curnode
			curnode = curnode.left
		} else {
			pnode = curnode
			curnode = curnode.right
		}
	}

	// insert
	node = &RBTreeNode[T]{
		Value:  *value,
		color:  RED,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	curnode = node
	curnode.parent = pnode
	if rules(value, &pnode.Value) {
		pnode.left = curnode
	} else {
		pnode.right = curnode
	}

	// update color and balancd
	for pnode.GetColor() == RED {
		gfnode := pnode.parent
		if pnode == gfnode.left {
			unode := gfnode.right
			if unode.GetColor() == RED { // parent, uncle = RED, RED
				pnode.color = BLACK
				unode.color = BLACK
				gfnode.color = RED

				curnode = gfnode
				pnode = curnode.parent
			} else {
				if curnode == pnode.right {
					pnode = pnode.leftRotate()
					gfnode.left = pnode
				}

				gffnode := gfnode.parent
				isLeft := gffnode != nil && gffnode.left == gfnode
				gfnode = gfnode.rightRotate()
				if gffnode == nil {
					*root = gfnode
				} else {
					if isLeft {
						gffnode.left = gfnode
					} else {
						gffnode.right = gfnode
					}
				}

				// update color
				pnode.color = BLACK
				gfnode.color = RED

				break
			}
		} else {
			unode := gfnode.left
			if unode.GetColor() == RED { // uncle, parent = RED, RED
				pnode.color = BLACK
				unode.color = BLACK
				gfnode.color = RED

				curnode = gfnode
				pnode = curnode.parent
			} else {
				if curnode == pnode.left {
					pnode = pnode.rightRotate()
					gfnode.right = pnode
				}

				gffnode := gfnode.parent
				isLeft := gffnode != nil && gffnode.left == gfnode
				gfnode = gfnode.leftRotate()
				if gffnode == nil {
					*root = gfnode
				} else {
					if isLeft {
						gffnode.left = gfnode
					} else {
						gffnode.right = gfnode
					}
				}

				// update color
				pnode.color = BLACK
				gfnode.color = RED

				break
			}
		}
	}

	// root is black
	(*root).color = BLACK

	return node
}

func (rbnode *RBTreeNode[T]) eraseNode(value *T, rules, equal func(a, b *T) bool) *RBTreeNode[T] {
	// update color and balance
	return rbnode
}

func (rb *RBTree[T]) deleteNode(node *RBTreeNode[T]) {
	// 获取节点
	var parent, gfnode, gffnode *RBTreeNode[T]
	var isLeft bool
	var curnode *RBTreeNode[T]

	if node.left == nil || node.right == nil {
		curnode = node
	} else {
		curnode = node.right
		for curnode.left != nil {
			curnode = curnode.left
		}
	}

	if curnode.left != nil {
		parent = curnode.parent
		gfnode = curnode
		isLeft = true
	} else {
		parent = curnode.parent
		gfnode = curnode.right
		isLeft = false
	}

	// 获取后继节点
	if gfnode.right != nil {
		gffnode = gfnode.right
		for gffnode.left != nil {
			gffnode = gffnode.left
		}
	} else {
		gffnode = gfnode
	}

	// 删除节点
	if parent == nil {
		rb.root = gffnode
	} else {
		if isLeft {
			parent.left = gffnode
		} else {
			parent.right = gffnode
		}
	}

	if gffnode != nil {
		gffnode.parent = parent
	}
}

type RBTree[T any] struct {
	root  *RBTreeNode[T]
	size  int
	rules func(a, b *T) bool
	equal func(a, b *T) bool
}

func NewRBTree[T any](rules, equal func(a, b *T) bool) RBTree[T] {
	return RBTree[T]{
		root:  nil,
		rules: rules,
		equal: equal,
	}
}

func (rb *RBTree[T]) Init(arr []T, rules, equal func(a, b *T) bool) {
	rb.Clear()
	rb.rules = rules
	rb.equal = equal
	for _, elem := range arr {
		rb.Insert(elem)
	}
}

func (rb *RBTree[T]) Size() int {
	return rb.size
}

func (rb *RBTree[T]) ForSize() int {
	size := 0
	rb.root.inOrder(func(_ *T) {
		size += 1
	})
	return size
}

func (rb *RBTree[T]) Empty() bool {
	return rb.root == nil
}

func (rb *RBTree[T]) Insert(value T) *RBTreeNode[T] {
	if rb.root == nil {
		rb.root = &RBTreeNode[T]{
			Value:  value,
			color:  BLACK,
			parent: nil,
			left:   nil,
			right:  nil,
		}
		rb.size += 1
		return rb.root
	}
	return rb.InsertByNode(rb.root, value)
}

func (rb *RBTree[T]) Erase(value T) {
	rb.EraseByNode(&rb.root, value)
}

func (rb *RBTree[T]) Find(value T) (*RBTreeNode[T], bool) {
	return rb.FindByNode(rb.root, value)
}

func (rb *RBTree[T]) LowerBound(value T) *RBTreeNode[T] {
	return rb.LowerBoundByNode(rb.root, value)
}

func (rb *RBTree[T]) UpperBound(value T) *RBTreeNode[T] {
	return rb.UpperBoundByNode(rb.root, value)
}

func (rb *RBTree[T]) Begin() *RBTreeNode[T] {
	if rb.root == nil {
		return nil
	}
	node := rb.root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (rb *RBTree[T]) End() *RBTreeNode[T] {
	if rb.root == nil {
		return nil
	}
	node := rb.root
	for node.right != nil {
		node = node.right
	}
	return node
}

func (rb *RBTree[T]) Front() (any, bool) {
	if rb.root == nil {
		return nil, false
	}
	node := rb.root
	for node.left != nil {
		node = node.left
	}
	return node.Value, true
}

func (rb *RBTree[T]) Back() (any, bool) {
	if rb.root == nil {
		return nil, false
	}
	node := rb.root
	for node.right != nil {
		node = node.right
	}
	return node.Value, true
}

func (rb *RBTree[T]) Clear() {
	rb.ClearByNode(&rb.root)
}

func (rb *RBTree[T]) ForEach(fn func(value *T)) {
	rb.root.inOrder(fn)
}

func (rb *RBTree[T]) String() string {
	str := "RBTree {\n"
	rb.root.inOrder(func(value *T) {
		str += "    " + fmt.Sprint(*value) + ",\n"
	})
	str += "}"
	return str
}

func (rb *RBTree[T]) FindByNode(rbnode *RBTreeNode[T], value T) (*RBTreeNode[T], bool) {
	if rbnode == nil {
		return nil, false
	}
	node := rbnode.findNode(&value, rb.rules, rb.equal)
	if node == nil {
		return nil, false
	}
	return node, true
}

func (rb *RBTree[T]) LowerBoundByNode(rbnode *RBTreeNode[T], value T) *RBTreeNode[T] {
	if rbnode == nil {
		return nil
	}
	return rbnode.lowerBoundNode(&value, rb.rules, rb.equal)
}

func (rb *RBTree[T]) UpperBoundByNode(rbnode *RBTreeNode[T], value T) *RBTreeNode[T] {
	if rbnode == nil {
		return nil
	}
	return rbnode.upperBoundNode(&value, rb.rules, rb.equal)
}

func (rb *RBTree[T]) InsertByNode(rbnode *RBTreeNode[T], value T) *RBTreeNode[T] {
	if rbnode == nil {
		return nil
	}
	node := rbnode.insertNode(&value, &rb.root, rb.rules, rb.equal)
	if node != nil {
		rb.size += 1
	}
	return node

}

func (rb *RBTree[T]) EraseByNode(rbnode **RBTreeNode[T], value T) {
	if *rbnode == nil {
		return
	}
	*rbnode = (*rbnode).eraseNode(&value, rb.rules, rb.equal)
}

func (rb *RBTree[T]) ClearByNode(rbnode **RBTreeNode[T]) {
	if *rbnode == nil {
		return
	}
	rb.ClearByNode(&(*rbnode).left)
	rb.ClearByNode(&(*rbnode).right)
	rb.size -= 1
	(*rbnode).color = BLACK
	(*rbnode).left = nil
	(*rbnode).right = nil
	(*rbnode).parent = nil
	*rbnode = nil
}
