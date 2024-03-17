/*
 * @Description: avltree_map
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2023-08-07 13:26:34
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-08-10 21:06:05
 */
package avltree

import (
	. "collections/common"
	"fmt"
)

type AVLTreeMap[KEY any, VALUE any] struct {
	*AVLTree[Pair[KEY, VALUE]]
}

func NewAVLTreeMap[KEY any, VALUE any](rules, equal func(a, b *KEY) bool) *AVLTreeMap[KEY, VALUE] {
	return &AVLTreeMap[KEY, VALUE]{
		AVLTree: &AVLTree[Pair[KEY, VALUE]]{
			root: nil,
			rules: func(a, b *Pair[KEY, VALUE]) bool {
				return rules(&a.Key, &b.Key)
			},
			equal: func(a, b *Pair[KEY, VALUE]) bool {
				return equal(&a.Key, &b.Key)
			},
		},
	}
}

func (treemap *AVLTreeMap[KEY, VALUE]) Init(arr []Pair[KEY, VALUE], rules, equal func(a, b *KEY) bool) {
	if treemap.AVLTree == nil {
		treemap.AVLTree = &AVLTree[Pair[KEY, VALUE]]{
			root: nil,
			rules: func(a, b *Pair[KEY, VALUE]) bool {
				return rules(&a.Key, &b.Key)
			},
			equal: func(a, b *Pair[KEY, VALUE]) bool {
				return equal(&a.Key, &b.Key)
			},
		}
	} else {
		treemap.Clear()
		treemap.rules = func(a, b *Pair[KEY, VALUE]) bool {
			return rules(&a.Key, &b.Key)
		}
		treemap.equal = func(a, b *Pair[KEY, VALUE]) bool {
			return equal(&a.Key, &b.Key)
		}
	}
	for _, elem := range arr {
		treemap.InsertPair(elem)
	}
}

func (treemap *AVLTreeMap[KEY, VALUE]) Insert(key KEY, value VALUE) *AVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.AVLTree.Insert(Pair[KEY, VALUE]{Key: key, Value: value})
}

func (treemap *AVLTreeMap[KEY, VALUE]) InsertPair(pair Pair[KEY, VALUE]) *AVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.AVLTree.Insert(pair)
}

func (treemap *AVLTreeMap[KEY, VALUE]) Erase(key KEY) {
	treemap.AVLTree.Erase(Pair[KEY, VALUE]{Key: key})
}

func (treemap *AVLTreeMap[KEY, VALUE]) Find(key KEY) (*AVLTreeNode[Pair[KEY, VALUE]], bool) {
	return treemap.AVLTree.Find(Pair[KEY, VALUE]{Key: key})
}

func (treemap *AVLTreeMap[KEY, VALUE]) Contains(key KEY) bool {
	return treemap.AVLTree.Contains(Pair[KEY, VALUE]{Key: key})
}

func (treemap *AVLTreeMap[KEY, VALUE]) Order(key KEY) int {
	return treemap.AVLTree.Order(Pair[KEY, VALUE]{Key: key})
}

func (treemap *AVLTreeMap[KEY, VALUE]) LowerBound(key KEY) *AVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.AVLTree.LowerBound(Pair[KEY, VALUE]{Key: key})
}

func (treemap *AVLTreeMap[KEY, VALUE]) UpperBound(key KEY) *AVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.AVLTree.UpperBound(Pair[KEY, VALUE]{Key: key})
}

func (treemap *AVLTreeMap[KEY, VALUE]) Keys() []KEY {
	keys := make([]KEY, 0, treemap.Size())
	treemap.AVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		keys = append(keys, pair.Key)
	})
	return keys
}

func (treemap *AVLTreeMap[KEY, VALUE]) ForEach(fn func(key *KEY, value *VALUE)) {
	treemap.AVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		fn(&pair.Key, &pair.Value)
	})
}

func (treemap *AVLTreeMap[KEY, VALUE]) ForEachPair(fn func(pair *Pair[KEY, VALUE])) {
	treemap.AVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		fn(pair)
	})
}

func (treemap *AVLTreeMap[KEY, VALUE]) String() string {
	str := "AVLTreeMap {\n"
	treemap.AVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		str += "    " + fmt.Sprint(pair.Key) + ": " + fmt.Sprint(pair.Value) + ",\n"
	})
	str += "}"
	return str
}
