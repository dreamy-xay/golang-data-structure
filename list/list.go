/*
 * @Description: list
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 13:43:38
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 16:24:49
 */
package list

import (
	. "collections/common"
	"fmt"
	"sort"
)

type ListNode[T any] struct {
	Value T
	pre   *ListNode[T]
	next  *ListNode[T]
}

func (listnode *ListNode[T]) GetPre() *ListNode[T] {
	if listnode == nil {
		return nil
	}
	return listnode.pre
}

func (listnode *ListNode[T]) GetNext() *ListNode[T] {
	if listnode == nil {
		return nil
	}
	return listnode.next
}

type List[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
	size int
}

func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (list *List[T]) Init(arr []T) {
	list.Clear()
	for _, elem := range arr {
		list.PushBack(elem)
	}
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) ForSize() int {
	size := 0
	for node := list.head; node != nil; node = node.next {
		size += 1
	}
	return size
}

func (list *List[T]) Empty() bool {
	return list.size == 0
}

func (list *List[T]) Insert(value T, position *ListNode[T], isNext bool) *ListNode[T] {
	list.size += 1
	if position == nil {
		list.head = &ListNode[T]{
			Value: value,
			pre:   nil,
			next:  nil,
		}
		list.tail = list.head
		return list.head
	}
	if isNext {
		if position.next != nil {
			pnext := position.next
			position.next = &ListNode[T]{
				Value: value,
				pre:   position,
				next:  pnext,
			}
			pnext.pre = position.next
		} else {
			position.next = &ListNode[T]{
				Value: value,
				pre:   position,
				next:  nil,
			}
			list.tail = position.next
		}
		return position.next
	} else {
		if position.pre != nil {
			ppre := position.pre
			position.pre = &ListNode[T]{
				Value: value,
				pre:   ppre,
				next:  position,
			}
			ppre.next = position.pre
		} else {
			position.pre = &ListNode[T]{
				Value: value,
				pre:   nil,
				next:  position,
			}
			list.head = position.pre
		}
		return position.pre
	}
}

func (list *List[T]) PushBack(value T) *ListNode[T] {
	return list.Insert(value, list.tail, true)
}

func (list *List[T]) PushFront(value T) *ListNode[T] {
	return list.Insert(value, list.head, false)
}

func (list *List[T]) Erase(position *ListNode[T]) {
	list.size -= 1
	if position.pre != nil && position.next != nil {
		position.pre.next = position.next.next
		position.pre = nil
		position.next = nil

	} else if position.pre != nil {
		list.tail = position.pre
		position.pre.next = nil
		position.pre = nil
	} else if position.next != nil {
		list.head = position.next
		position.next.pre = nil
		position.next = nil
	} else {
		list.head = nil
		list.tail = nil
		list.size = 0
	}
}

func (list *List[T]) PopBack() {
	list.Erase(list.tail)
}

func (list *List[T]) PopFront() {
	list.Erase(list.head)
}

func (list *List[T]) Begin() *ListNode[T] {
	return list.head
}

func (list *List[T]) End() *ListNode[T] {
	return list.tail
}

func (list *List[T]) Front() (any, bool) {
	if list.head == nil {
		return nil, false
	}
	return list.head.Value, true
}

func (list *List[T]) Back() (any, bool) {
	if list.tail == nil {
		return nil, false
	}
	return list.tail.Value, true
}

func (list *List[T]) Reverse() {
	var next *ListNode[T]
	node := list.head
	for node != nil {
		next = node.next
		node.next = node.pre
		node.pre = next
		node = next
	}
	node = list.head
	list.head = list.tail
	list.tail = node
}

func (list *List[T]) Sort(fn func(a, b *T) bool) {
	vec := make([]T, 0, list.ForSize())
	for node := list.head; node != nil; node = node.next {
		vec = append(vec, node.Value)
	}
	sort.Slice(vec, func(i, j int) bool {
		return fn(&vec[i], &vec[j])
	})
	i := 0
	for node := list.head; node != nil; node = node.next {
		node.Value = vec[i]
		i += 1
	}
}

func (list *List[T]) Clear() {
	list.size = 0
	list.tail = nil
	if list.head != nil {
		var next *ListNode[T]
		for list.head.next != nil {
			next = list.head.next
			list.head.next = nil
			list.head = next
		}
		list.head = nil
	}
}

func (list *List[T]) ForEach(fn func(index int, elem *T)) {
	index := 0
	for node := list.head; node != nil; node = node.next {
		fn(index, &node.Value)
		index += 1
	}
}

func (list *List[T]) GetIterator() *Iterator[T] {
	current := list.Begin()
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

func (list *List[T]) String() string {
	str := "List {\n"
	for node := list.head; node != nil; node = node.next {
		str += "    " + fmt.Sprint(node.Value) + ",\n"
	}
	str += "}"
	return str
}
