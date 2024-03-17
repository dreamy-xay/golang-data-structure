/*
 * @Description: RBTree
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-16 21:43:01
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-17 22:47:22
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

type RBTreeNode struct {
	Value  interface{}
	color  Color
	left   *RBTreeNode
	right  *RBTreeNode
	parent *RBTreeNode
}

func (rbnode *RBTreeNode) GetColor() Color {
	if rbnode == nil {
		return BLACK
	}
	return rbnode.color
}

func (rbnode *RBTreeNode) GetPre() *RBTreeNode {
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

func (rbnode *RBTreeNode) GetNext() *RBTreeNode {
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

func (rbnode *RBTreeNode) inOrder(fn func(value *interface{})) {
	if rbnode == nil {
		return
	}
	rbnode.left.inOrder(fn)
	fn(&rbnode.Value)
	rbnode.right.inOrder(fn)
}

func (rbnode *RBTreeNode) findNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *RBTreeNode {
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

func (rbnode *RBTreeNode) lowerBoundNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *RBTreeNode {
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

func (rbnode *RBTreeNode) upperBoundNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *RBTreeNode {
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

func (rbnode *RBTreeNode) leftRotate() *RBTreeNode {
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

func (rbnode *RBTreeNode) rightRotate() *RBTreeNode {
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

func insertNode(value *interface{}, rbnode *RBTreeNode, root **RBTreeNode, rules, equal func(a, b *interface{}) bool) *RBTreeNode {
	// define
	var (
		pnode   *RBTreeNode = nil
		curnode *RBTreeNode = rbnode
		node    *RBTreeNode = nil
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
	node = &RBTreeNode{
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

func (rbnode *RBTreeNode) eraseNode(value *interface{}, rules, equal func(a, b *interface{}) bool) *RBTreeNode {
	// update color and balance
	return rbnode
}

type RBTree struct {
	root  *RBTreeNode
	size  int
	rules func(a, b *interface{}) bool
	equal func(a, b *interface{}) bool
}

func NewRBTree(rules, equal func(a, b *interface{}) bool) RBTree {
	return RBTree{
		root:  nil,
		rules: rules,
		equal: equal,
	}
}

func (rb *RBTree) Init(arr []interface{}, rules, equal func(a, b *interface{}) bool) {
	rb.Clear()
	rb.rules = rules
	rb.equal = equal
	for _, elem := range arr {
		rb.Insert(elem)
	}
}

func (rb *RBTree) Size() int {
	return rb.size
}

func (rb *RBTree) ForSize() int {
	size := 0
	rb.root.inOrder(func(_ *interface{}) {
		size += 1
	})
	return size
}

func (rb *RBTree) Empty() bool {
	return rb.root == nil
}

func (rb *RBTree) Insert(value interface{}) *RBTreeNode {
	if rb.root == nil {
		rb.root = &RBTreeNode{
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

func (rb *RBTree) Erase(value interface{}) {
	rb.EraseByNode(&rb.root, value)
}

func (rb *RBTree) Find(value interface{}) (*RBTreeNode, bool) {
	return rb.FindByNode(rb.root, value)
}

func (rb *RBTree) LowerBound(value interface{}) *RBTreeNode {
	return rb.LowerBoundByNode(rb.root, value)
}

func (rb *RBTree) UpperBound(value interface{}) *RBTreeNode {
	return rb.UpperBoundByNode(rb.root, value)
}

func (rb *RBTree) Begin() *RBTreeNode {
	if rb.root == nil {
		return nil
	}
	node := rb.root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (rb *RBTree) End() *RBTreeNode {
	if rb.root == nil {
		return nil
	}
	node := rb.root
	for node.right != nil {
		node = node.right
	}
	return node
}

func (rb *RBTree) Front() (interface{}, bool) {
	if rb.root == nil {
		return nil, false
	}
	node := rb.root
	for node.left != nil {
		node = node.left
	}
	return node.Value, true
}

func (rb *RBTree) Back() (interface{}, bool) {
	if rb.root == nil {
		return nil, false
	}
	node := rb.root
	for node.right != nil {
		node = node.right
	}
	return node.Value, true
}

func (rb *RBTree) Clear() {
	rb.ClearByNode(&rb.root)
}

func (rb *RBTree) ForEach(fn func(value *interface{})) {
	rb.root.inOrder(fn)
}

func (rb *RBTree) String() string {
	str := "RBTree {\n"
	rb.root.inOrder(func(value *interface{}) {
		str += "    " + fmt.Sprint(*value) + ",\n"
	})
	str += "}"
	return str
}

func (rb *RBTree) FindByNode(rbnode *RBTreeNode, value interface{}) (*RBTreeNode, bool) {
	if rbnode == nil {
		return nil, false
	}
	node := rbnode.findNode(&value, rb.rules, rb.equal)
	if node == nil {
		return nil, false
	}
	return node, true
}

func (rb *RBTree) LowerBoundByNode(rbnode *RBTreeNode, value interface{}) *RBTreeNode {
	if rbnode == nil {
		return nil
	}
	return rbnode.lowerBoundNode(&value, rb.rules, rb.equal)
}

func (rb *RBTree) UpperBoundByNode(rbnode *RBTreeNode, value interface{}) *RBTreeNode {
	if rbnode == nil {
		return nil
	}
	return rbnode.upperBoundNode(&value, rb.rules, rb.equal)
}

func (rb *RBTree) InsertByNode(rbnode *RBTreeNode, value interface{}) *RBTreeNode {
	if rbnode == nil {
		return nil
	}
	node := insertNode(&value, rbnode, &rb.root, rb.rules, rb.equal)
	if node != nil {
		rb.size += 1
	}
	return node

}

func (rb *RBTree) EraseByNode(rbnode **RBTreeNode, value interface{}) {
	if *rbnode == nil {
		return
	}
	*rbnode = (*rbnode).eraseNode(&value, rb.rules, rb.equal)
}

func (rb *RBTree) ClearByNode(rbnode **RBTreeNode) {
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
