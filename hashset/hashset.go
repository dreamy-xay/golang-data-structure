/*
 * @Description: hashset
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 21:51:20
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 16:06:55
 */
package hashset

import (
	. "collections/common"
	"fmt"
)

type HashSet[T comparable] struct {
	elements map[T]bool
}

func NewHashSet[T comparable]() HashSet[T] {
	return HashSet[T]{elements: make(map[T]bool)}
}

func NewHashSetCapacity[T comparable](capacity int) *HashSet[T] {
	return &HashSet[T]{elements: make(map[T]bool, capacity)}
}

func (hashset *HashSet[T]) Init(arr []T) {
	hashset.Clear()
	for _, elem := range arr {
		hashset.elements[elem] = true
	}
}

func (hashset *HashSet[T]) Empty() bool {
	return len(hashset.elements) == 0
}

func (hashset *HashSet[T]) Size() int {
	return len(hashset.elements)
}

func (hashset *HashSet[T]) Clear() {
	for key := range hashset.elements {
		delete(hashset.elements, key)
	}
}

func (hashset *HashSet[T]) Insert(key T) {
	hashset.elements[key] = true
}

func (hashset *HashSet[T]) Erase(key T) {
	delete(hashset.elements, key)
}

func (hashset *HashSet[T]) Find(key T) (T, bool) {
	_, ok := (hashset.elements)[key]
	return key, ok
}

func (hashset *HashSet[T]) Contains(key T) bool {
	_, ok := hashset.elements[key]
	return ok
}

func (hashset *HashSet[T]) Keys() []T {
	keys := make([]T, 0, len(hashset.elements))
	for key := range hashset.elements {
		keys = append(keys, key)
	}
	return keys
}

func (hashset *HashSet[T]) ForEach(fn func(key *T)) {
	for key, value := range hashset.elements {
		if value {
			fn(&key)
		}
	}
}

func (hashset *HashSet[T]) ToMap() map[T]bool {
	newMap := make(map[T]bool, len(hashset.elements))
	for key, value := range hashset.elements {
		newMap[key] = value
	}
	return newMap
}

func (hashset *HashSet[T]) GetIterator() *Iterator[T] {
	setList := []T{}
	for key := range hashset.elements {
		setList = append(setList, key)
	}
	current := -1
	listLen := len(setList)
	return NewIterator[T](func() bool {
		return current+1 < listLen
	}, func() *T {
		current += 1
		if current < listLen {
			return &setList[current]
		} else {
			return nil
		}
	})
}

func (hashset *HashSet[T]) String() string {
	str := "HashSet {\n"
	for key := range hashset.elements {
		str += "    " + fmt.Sprint(key) + ",\n"
	}
	str += "}"
	return str
}
