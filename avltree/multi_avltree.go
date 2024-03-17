/*
 * @Description: multi_avltree
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2023-08-07 16:15:14
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 15:36:00
 */
package avltree

import (
	. "collections/common"
	"fmt"
)

type MultiAVLTreeNode[T any] struct {
	Value                T                           // 值
	valueList            *AVLTreeMap[uint, T]        // equal函数判断相同值形成的列表（此处使用AVLTreeMap增速）
	currentIterativeNode *AVLTreeNode[Pair[uint, T]] // 当前迭代节点
	valueListCounter     uint                        // 相同值形成的列表所对应map key的计数器
	height               int                         // 所在位置相比于整棵树的高度
	count                int                         // 所在位置作为子树根节点 子树的节点总数量
	left                 *MultiAVLTreeNode[T]
	right                *MultiAVLTreeNode[T]
	parent               *MultiAVLTreeNode[T]
}

func (avlnode *MultiAVLTreeNode[T]) updateCurrentIterativeNode(node *AVLTreeNode[Pair[uint, T]]) {
	avlnode.currentIterativeNode = node                      // 更新【当前迭代节点】
	avlnode.Value = avlnode.currentIterativeNode.Value.Value // 更新值
}

func (avlnode *MultiAVLTreeNode[T]) GetHeight() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.height
}

func (avlnode *MultiAVLTreeNode[T]) GetCount() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.count
}

func (avlnode *MultiAVLTreeNode[T]) GetPreDifferent() *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	// 计算上一个节点
	prenode := func() *MultiAVLTreeNode[T] {
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
	}()

	// 重置新节点的迭代索引数
	if prenode != nil {
		prenode.updateCurrentIterativeNode(prenode.valueList.End()) // 更新迭代节点状态
	}

	return prenode
}

func (avlnode *MultiAVLTreeNode[T]) GetNextDifferent() *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	// 计算下一个节点
	nextnode := func() *MultiAVLTreeNode[T] {
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
	}()

	// 重置新节点的迭代索引数
	if nextnode != nil {
		nextnode.updateCurrentIterativeNode(nextnode.valueList.Begin()) // 更新迭代节点状态
	}

	return nextnode
}

func (avlnode *MultiAVLTreeNode[T]) GetPre() *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	// 相同的Value未迭代完全
	if avlnode.currentIterativeNode.GetPre() != nil {
		avlnode.updateCurrentIterativeNode(avlnode.currentIterativeNode.GetPre()) // 更新迭代节点状态
		return avlnode
	}

	// 相同的Value迭代完了，直接迭代上一个不同于avlnode.Value的Value所对应的节点
	return avlnode.GetPreDifferent()
}

func (avlnode *MultiAVLTreeNode[T]) GetNext() *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	// 相同的Value未迭代完全
	if avlnode.currentIterativeNode.GetNext() != nil {
		avlnode.updateCurrentIterativeNode(avlnode.currentIterativeNode.GetNext()) // 更新迭代节点状态
		return avlnode
	}

	// 相同的Value迭代完了，直接迭代下一个不同于avlnode.Value的Value所对应的节点
	return avlnode.GetNextDifferent()
}

func (avlnode *MultiAVLTreeNode[T]) getBalanceFactor() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.left.GetHeight() - avlnode.right.GetHeight()
}

func (avlnode *MultiAVLTreeNode[T]) inOrder(fn func(value *T)) {
	if avlnode == nil {
		return
	}
	avlnode.left.inOrder(fn)
	// 迭代相同值
	for valueListNode := avlnode.valueList.Begin(); valueListNode != nil; valueListNode = valueListNode.GetNext() {
		// 此处不更新迭代节点、不更新迭代值（avlnode 不会被外部使用）
		fn(&valueListNode.Value.Value)
	}
	avlnode.right.inOrder(fn)
}

func (avlnode *MultiAVLTreeNode[T]) findNode(value *T, rules, equal func(a, b *T) bool) *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	if equal(value, &avlnode.Value) {
		avlnode.updateCurrentIterativeNode(avlnode.valueList.Begin()) // 更新迭代节点状态
		return avlnode
	} else if rules(value, &avlnode.Value) {
		return avlnode.left.findNode(value, rules, equal)
	} else {
		return avlnode.right.findNode(value, rules, equal)
	}
}

func (avlnode *MultiAVLTreeNode[T]) orderNode(value *T, rules, equal func(a, b *T) bool) int {
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
			return count + avlnode.left.GetCount() + avlnode.valueList.Size() // avlnode.valueList.Size() 为equal函数判断相同的值数量
		}
		return 0
	}
}

func (avlnode *MultiAVLTreeNode[T]) atNode(index, preCount int) *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	count := avlnode.left.GetCount() + preCount
	if index < count {
		return avlnode.left.atNode(index, preCount)
	} else if index >= count+avlnode.valueList.Size() {
		return avlnode.right.atNode(index, count+avlnode.valueList.Size())
	} else {
		avlnode.updateCurrentIterativeNode(avlnode.valueList.At(index - count)) // 更新迭代节点状态
		return avlnode
	}
}

func (avlnode *MultiAVLTreeNode[T]) lowerBoundNode(value *T, rules, equal func(a, b *T) bool) *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	if equal(value, &avlnode.Value) {
		avlnode.updateCurrentIterativeNode(avlnode.valueList.Begin()) // 更新迭代节点状态
		return avlnode
	} else if rules(value, &avlnode.Value) {
		node := avlnode.left.lowerBoundNode(value, rules, equal)
		if node == nil {
			avlnode.updateCurrentIterativeNode(avlnode.valueList.Begin()) // 更新迭代节点状态
			return avlnode
		}
		node.updateCurrentIterativeNode(node.valueList.Begin()) // 更新迭代节点状态
		return node
	} else {
		return avlnode.right.lowerBoundNode(value, rules, equal)
	}
}

func (avlnode *MultiAVLTreeNode[T]) upperBoundNode(value *T, rules, equal func(a, b *T) bool) *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	if equal(value, &avlnode.Value) {
		// 此处未更新【当前迭代节点】，因为GetNextDifferent函数已经自动更新了【当前迭代节点】
		return avlnode.GetNextDifferent()
	} else if rules(value, &avlnode.Value) {
		node := avlnode.left.upperBoundNode(value, rules, equal)
		if node == nil {
			avlnode.updateCurrentIterativeNode(avlnode.valueList.Begin()) // 更新迭代节点状态
			return avlnode
		}
		node.updateCurrentIterativeNode(node.valueList.Begin()) // 更新迭代节点状态
		return node
	} else {
		return avlnode.right.upperBoundNode(value, rules, equal)
	}
}

func (avlnode *MultiAVLTreeNode[T]) leftRotate() *MultiAVLTreeNode[T] {
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
	avlnode.count = avlnode.left.GetCount() + avlnode.right.GetCount() + avlnode.valueList.Size()

	// update height
	avlnode.height = Max(avlnode.left.GetHeight(), avlnode.right.GetHeight()) + 1
	mid.height = Max(mid.left.GetHeight(), mid.right.GetHeight()) + 1
	return mid
}

func (avlnode *MultiAVLTreeNode[T]) rightRotate() *MultiAVLTreeNode[T] {
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
	avlnode.count = avlnode.left.GetCount() + avlnode.right.GetCount() + avlnode.valueList.Size()

	// update height
	avlnode.height = Max(avlnode.left.GetHeight(), avlnode.right.GetHeight()) + 1
	mid.height = Max(mid.left.GetHeight(), mid.right.GetHeight()) + 1
	return mid
}

func (avlnode *MultiAVLTreeNode[T]) balanceNode() *MultiAVLTreeNode[T] {
	// update height
	avlnode.height = Max(avlnode.left.GetHeight(), avlnode.right.GetHeight()) + 1

	// update count
	avlnode.count = avlnode.left.GetCount() + avlnode.right.GetCount() + avlnode.valueList.Size()

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

func (avlnode *MultiAVLTreeNode[T]) insertNode(node, parent *MultiAVLTreeNode[T], rules, equal func(a, b *T) bool) *MultiAVLTreeNode[T] {
	// insert
	if avlnode == nil {
		// 新建的node
		node.parent = parent
		// 更新node迭代列表和状态
		node.valueList = NewAVLTreeMap[uint, T](Less[uint], Equal[uint])
		node.currentIterativeNode = node.valueList.Insert(node.valueListCounter, node.Value) // 迭代列表插入值并更新【当前迭代节点】
		node.valueListCounter += 1                                                           // 相同值形成的列表所对应map key的计数器 增加1
		return node
	}
	if equal(&node.Value, &avlnode.Value) {
		// 更新子树节点数量
		avlnode.count += 1
		// 相同则只更新node迭代列表和状态
		avlnode.updateCurrentIterativeNode(avlnode.valueList.Insert(avlnode.valueListCounter, node.Value)) // 迭代列表插入值并更新【当前迭代节点】
		avlnode.valueListCounter += 1                                                                      // 相同值形成的列表所对应map key的计数器 增加1
		*node = *avlnode                                                                                   // 更新插入节点位置
		return avlnode
	} else if rules(&node.Value, &avlnode.Value) {
		avlnode.left = avlnode.left.insertNode(node, avlnode, rules, equal)
	} else {
		avlnode.right = avlnode.right.insertNode(node, avlnode, rules, equal)
	}

	// update height and balance
	return avlnode.balanceNode()
}

func (avlnode *MultiAVLTreeNode[T]) eraseNode(node *MultiAVLTreeNode[T], all bool, rules, equal func(a, b *T) bool) *MultiAVLTreeNode[T] {
	// all 表示是否删除全部（清空全部相同数），否则删除一个

	if avlnode == nil {
		return nil
	}

	curnode := avlnode
	if equal(&node.Value, &avlnode.Value) {
		// 如果不是删除全部则数量减少1就行（数量减少1后等于0相当于删除该节点）
		if !all && avlnode.valueList.Size() > 1 {
			// 判断 node 合理性
			if node.valueList != avlnode.valueList { // node不合理，迭代节点和列表无效
				return avlnode
			}

			// 更新子树节点数量
			avlnode.count -= 1
			// 只更新node迭代列表和状态
			if node.currentIterativeNode.GetNext() != nil {
				avlnode.updateCurrentIterativeNode(node.currentIterativeNode.GetNext()) // 更新迭代节点状态
			} else {
				avlnode.updateCurrentIterativeNode(avlnode.valueList.Begin()) // 更新迭代节点状态
			}
			avlnode.valueList.Erase(node.currentIterativeNode.Value.Key) // 更新node迭代列表

			return avlnode
		}

		if avlnode.left == nil {
			curnode = avlnode.right
			if curnode != nil {
				curnode.parent = avlnode.parent
			}
			avlnode.right = nil
			avlnode.parent = nil
		} else if avlnode.right == nil {
			curnode = avlnode.left
			if curnode != nil {
				curnode.parent = avlnode.parent
			}
			avlnode.left = nil
			avlnode.parent = nil
		} else {
			curnode = avlnode.right
			for curnode.left != nil {
				curnode = curnode.left
			}
			curnode.right = avlnode.right.eraseNode(curnode, all, rules, equal)
			curnode.left = avlnode.left
			curnode.parent = avlnode.parent
			curnode.left.parent = curnode
			if curnode.right != nil {
				curnode.right.parent = curnode
			}
			avlnode.left = nil
			avlnode.right = nil
			avlnode.parent = nil
		}
	} else if rules(&node.Value, &avlnode.Value) {
		avlnode.left = avlnode.left.eraseNode(node, all, rules, equal)
	} else {
		avlnode.right = avlnode.right.eraseNode(node, all, rules, equal)
	}

	if curnode == nil {
		return nil
	}
	// update height and balance
	return curnode.balanceNode()
}

type MultiAVLTree[T any] struct {
	root  *MultiAVLTreeNode[T]
	rules func(a, b *T) bool
	equal func(a, b *T) bool
}

func NewMultiAVLTree[T any](rules, equal func(a, b *T) bool) *MultiAVLTree[T] {
	return &MultiAVLTree[T]{
		root:  nil,
		rules: rules,
		equal: equal,
	}
}

func (avl *MultiAVLTree[T]) Init(arr []T, rules, equal func(a, b *T) bool) {
	avl.Clear()
	avl.rules = rules
	avl.equal = equal
	for _, elem := range arr {
		avl.Insert(elem)
	}
}

func (avl *MultiAVLTree[T]) Size() int {
	return avl.root.GetCount()
}

func (avl *MultiAVLTree[T]) ForSize() int {
	size := 0
	avl.root.inOrder(func(_ *T) {
		size += 1
	})
	return size
}

func (avl *MultiAVLTree[T]) Empty() bool {
	return avl.root == nil
}

func (avl *MultiAVLTree[T]) Height() int {
	return avl.root.GetHeight()
}

func (avl *MultiAVLTree[T]) Insert(value T) *MultiAVLTreeNode[T] {
	if avl.root == nil {
		// 新建迭代列表并更新当前迭代节点
		valueList := NewAVLTreeMap[uint, T](Less[uint], Equal[uint])
		valueList.Insert(0, value)

		avl.root = &MultiAVLTreeNode[T]{
			Value:                value,
			valueList:            valueList,
			currentIterativeNode: valueList.Begin(),
			valueListCounter:     1,
			height:               1,
			count:                1,
			parent:               nil,
			left:                 nil,
			right:                nil,
		}
		return avl.root
	}
	return avl.InsertByNode(&avl.root, value)
}

func (avl *MultiAVLTree[T]) Erase(node *MultiAVLTreeNode[T]) {
	avl.EraseByNode(&avl.root, node)
}

func (avl *MultiAVLTree[T]) EraseAll(value T) {
	avl.EraseAllByNode(&avl.root, value)
}

func (avl *MultiAVLTree[T]) Find(value T) (*MultiAVLTreeNode[T], bool) {
	return avl.FindByNode(avl.root, value)
}

func (avl *MultiAVLTree[T]) Contains(value T) bool {
	_, ok := avl.FindByNode(avl.root, value)
	return ok
}

func (avl *MultiAVLTree[T]) Order(value T) int {
	return avl.OrderByNode(avl.root, value)
}

func (avl *MultiAVLTree[T]) At(index int) *MultiAVLTreeNode[T] {
	return avl.AtByNode(avl.root, index)
}

func (avl *MultiAVLTree[T]) LowerBound(value T) *MultiAVLTreeNode[T] {
	return avl.LowerBoundByNode(avl.root, value)
}

func (avl *MultiAVLTree[T]) UpperBound(value T) *MultiAVLTreeNode[T] {
	return avl.UpperBoundByNode(avl.root, value)
}

func (avl *MultiAVLTree[T]) Begin() *MultiAVLTreeNode[T] {
	if avl.root == nil {
		return nil
	}
	node := avl.root
	for node.left != nil {
		node = node.left
	}
	node.updateCurrentIterativeNode(node.valueList.Begin()) // 更新迭代节点状态
	return node
}

func (avl *MultiAVLTree[T]) End() *MultiAVLTreeNode[T] {
	if avl.root == nil {
		return nil
	}
	node := avl.root
	for node.right != nil {
		node = node.right
	}
	node.updateCurrentIterativeNode(node.valueList.End()) // 更新迭代节点状态
	return node
}

func (avl *MultiAVLTree[T]) Front() (any, bool) {
	if avl.root == nil {
		return nil, false
	}
	node := avl.root
	for node.left != nil {
		node = node.left
	}
	return node.Value, true
}

func (avl *MultiAVLTree[T]) Back() (any, bool) {
	if avl.root == nil {
		return nil, false
	}
	node := avl.root
	for node.right != nil {
		node = node.right
	}
	return node.Value, true
}

func (avl *MultiAVLTree[T]) Clear() {
	avl.ClearByNode(&avl.root)
}

func (avl *MultiAVLTree[T]) Count(value T) int {
	return avl.CountByNode(avl.root, value)
}

func (avl *MultiAVLTree[T]) IsAVLTree() bool {
	return avl.IsAVLTreeByNode(avl.root)
}

func (avl *MultiAVLTree[T]) ForEach(fn func(value *T)) {
	avl.root.inOrder(fn)
}

func (avl *MultiAVLTree[T]) GetIterator() *Iterator[T] {
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

func (avl *MultiAVLTree[T]) String() string {
	str := "MultiAVLTree {\n"
	avl.root.inOrder(func(value *T) {
		str += "    " + fmt.Sprint(*value) + ",\n"
	})
	str += "}"
	return str
}

func (avl *MultiAVLTree[T]) IsAVLTreeByNode(avlnode *MultiAVLTreeNode[T]) bool {
	if avlnode == nil {
		return true
	}
	if Abs(avlnode.getBalanceFactor()) > 1 || !((avlnode.left == nil || avl.rules(&avlnode.left.Value, &avlnode.Value)) && (avlnode.right == nil || !avl.rules(&avlnode.right.Value, &avlnode.Value))) {
		return false
	}
	return avl.IsAVLTreeByNode(avlnode.left) && avl.IsAVLTreeByNode(avlnode.right)
}

func (avl *MultiAVLTree[T]) FindByNode(avlnode *MultiAVLTreeNode[T], value T) (*MultiAVLTreeNode[T], bool) {
	if avlnode == nil {
		return nil, false
	}
	node := avlnode.findNode(&value, avl.rules, avl.equal)
	if node == nil {
		return nil, false
	}
	return node, true
}

func (avl *MultiAVLTree[T]) OrderByNode(avlnode *MultiAVLTreeNode[T], value T) int {
	if avlnode == nil {
		return -1
	}
	return avlnode.orderNode(&value, avl.rules, avl.equal) - 1
}

func (avl *MultiAVLTree[T]) AtByNode(avlnode *MultiAVLTreeNode[T], index int) *MultiAVLTreeNode[T] {
	if avlnode == nil || index < 0 || index >= avl.Size() {
		return nil
	}
	return avlnode.atNode(index, 0)
}

func (avl *MultiAVLTree[T]) LowerBoundByNode(avlnode *MultiAVLTreeNode[T], value T) *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	return avlnode.lowerBoundNode(&value, avl.rules, avl.equal)
}

func (avl *MultiAVLTree[T]) UpperBoundByNode(avlnode *MultiAVLTreeNode[T], value T) *MultiAVLTreeNode[T] {
	if avlnode == nil {
		return nil
	}
	return avlnode.upperBoundNode(&value, avl.rules, avl.equal)
}

func (avl *MultiAVLTree[T]) InsertByNode(avlnode **MultiAVLTreeNode[T], value T) *MultiAVLTreeNode[T] {
	if *avlnode == nil {
		return nil
	}
	node := &MultiAVLTreeNode[T]{
		Value:                value,
		valueList:            nil,
		currentIterativeNode: nil,
		valueListCounter:     0,
		height:               1,
		count:                1,
		left:                 nil,
		right:                nil,
		parent:               nil,
	}
	*avlnode = (*avlnode).insertNode(node, (*avlnode).parent, avl.rules, avl.equal)
	return node
}

func (avl *MultiAVLTree[T]) EraseByNode(avlnode **MultiAVLTreeNode[T], node *MultiAVLTreeNode[T]) {
	if *avlnode == nil || node == nil {
		return
	}
	*avlnode = (*avlnode).eraseNode(node, false, avl.rules, avl.equal)
}

func (avl *MultiAVLTree[T]) EraseAllByNode(avlnode **MultiAVLTreeNode[T], value T) {
	if *avlnode == nil {
		return
	}
	*avlnode = (*avlnode).eraseNode(&MultiAVLTreeNode[T]{Value: value}, true, avl.rules, avl.equal)
}

func (avl *MultiAVLTree[T]) ClearByNode(avlnode **MultiAVLTreeNode[T]) {
	if *avlnode == nil {
		return
	}
	avl.ClearByNode(&(*avlnode).left)
	avl.ClearByNode(&(*avlnode).right)
	(*avlnode).left = nil
	(*avlnode).right = nil
	(*avlnode).parent = nil
	*avlnode = nil
}

func (avl *MultiAVLTree[T]) CountByNode(avlnode *MultiAVLTreeNode[T], value T) int {
	if avlnode == nil {
		return 0
	}
	node := avlnode.findNode(&value, avl.rules, avl.equal)
	if node == nil {
		return 0
	}
	return node.valueList.Size()
}
