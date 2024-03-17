/*
 * @Description: multi_avltree_map
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2023-08-07 13:26:34
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-08-10 20:51:39
 */
package avltree

import (
	. "collections/common"
	"fmt"
)

type MultiAVLTreeMap[KEY any, VALUE any] struct {
	*MultiAVLTree[Pair[KEY, VALUE]]
}

func NewMultiAVLTreeMap[KEY any, VALUE any](rules, equal func(a, b *KEY) bool) *MultiAVLTreeMap[KEY, VALUE] {
	return &MultiAVLTreeMap[KEY, VALUE]{
		MultiAVLTree: &MultiAVLTree[Pair[KEY, VALUE]]{
			root: nil,
			rules: func(a, b *Pair[KEY, VALUE]) bool {
				return rules(&a.Key, &b.Key)
			},
			equal: func(a, b *Pair[KEY, VALUE]) bool {
				return equal(&a.Key, &b.Key)
			}},
	}
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Init(arr []Pair[KEY, VALUE], rules, equal func(a, b *KEY) bool) {
	if treemap.MultiAVLTree == nil {
		treemap.MultiAVLTree = &MultiAVLTree[Pair[KEY, VALUE]]{
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
	// arr 插入
	for _, elem := range arr {
		treemap.InsertPair(elem)
	}
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Insert(key KEY, value VALUE) *MultiAVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.MultiAVLTree.Insert(Pair[KEY, VALUE]{Key: key, Value: value})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) InsertPair(pair Pair[KEY, VALUE]) *MultiAVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.MultiAVLTree.Insert(pair)
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) EraseAll(key KEY) {
	treemap.MultiAVLTree.EraseAll(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Find(key KEY) (*MultiAVLTreeNode[Pair[KEY, VALUE]], bool) {
	return treemap.MultiAVLTree.Find(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Contains(key KEY) bool {
	return treemap.MultiAVLTree.Contains(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Order(key KEY) int {
	return treemap.MultiAVLTree.Order(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) LowerBound(key KEY) *MultiAVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.MultiAVLTree.LowerBound(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) UpperBound(key KEY) *MultiAVLTreeNode[Pair[KEY, VALUE]] {
	return treemap.MultiAVLTree.UpperBound(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Count(key KEY) int {
	return treemap.MultiAVLTree.Count(Pair[KEY, VALUE]{Key: key})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) Keys() []KEY {
	keys := make([]KEY, 0, treemap.Size())
	treemap.MultiAVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		keys = append(keys, pair.Key)
	})
	return keys
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) ForEach(fn func(key *KEY, value *VALUE)) {
	treemap.MultiAVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		fn(&pair.Key, &pair.Value)
	})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) ForEachPair(fn func(pair *Pair[KEY, VALUE])) {
	treemap.MultiAVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		fn(pair)
	})
}

func (treemap *MultiAVLTreeMap[KEY, VALUE]) String() string {
	str := "MultiAVLTreeMap {\n"
	treemap.MultiAVLTree.ForEach(func(pair *Pair[KEY, VALUE]) {
		str += "    " + fmt.Sprint(pair.Key) + ": " + fmt.Sprint(pair.Value) + ",\n"
	})
	str += "}"
	return str
}
