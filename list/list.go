/*
 * @Description: list
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 13:43:38
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-14 14:45:40
 */
package list

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Value interface{}
	pre   *ListNode
	next  *ListNode
}

func (listnode *ListNode) GetPre() *ListNode {
	return listnode.pre
}

func (listnode *ListNode) GetNext() *ListNode {
	return listnode.next
}

type List struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (list *List) Init(arr []interface{}) {
	list.Clear()
	for _, elem := range arr {
		list.PushBack(elem)
	}
}

func (list *List) Size() int {
	return list.size
}

func (list *List) ForSize() int {
	size := 0
	for node := list.head; node != nil; node = node.next {
		size += 1
	}
	return size
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Insert(value interface{}, position *ListNode, isNext bool) *ListNode {
	list.size += 1
	if position == nil {
		list.head = &ListNode{
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
			position.next = &ListNode{
				Value: value,
				pre:   position,
				next:  pnext,
			}
			pnext.pre = position.next
		} else {
			position.next = &ListNode{
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
			position.pre = &ListNode{
				Value: value,
				pre:   ppre,
				next:  position,
			}
			ppre.next = position.pre
		} else {
			position.pre = &ListNode{
				Value: value,
				pre:   nil,
				next:  position,
			}
			list.head = position.pre
		}
		return position.pre
	}
}

func (list *List) PushBack(value interface{}) *ListNode {
	return list.Insert(value, list.tail, true)
}

func (list *List) PushFront(value interface{}) *ListNode {
	return list.Insert(value, list.head, false)
}

func (list *List) Erase(position *ListNode) {
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

func (list *List) PopBack() {
	list.Erase(list.tail)
}

func (list *List) PopFront() {
	list.Erase(list.head)
}

func (list *List) Begin() *ListNode {
	return list.head
}

func (list *List) End() *ListNode {
	return list.tail
}

func (list *List) Front() (interface{}, bool) {
	if list.head == nil {
		return nil, false
	}
	return list.head.Value, true
}

func (list *List) Back() (interface{}, bool) {
	if list.tail == nil {
		return nil, false
	}
	return list.tail.Value, true
}

func (list *List) Reverse() {
	var next *ListNode
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

func (list *List) Sort(fn func(a, b *interface{}) bool) {
	vec := make([]interface{}, 0, list.ForSize())
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

func (list *List) Clear() {
	list.size = 0
	list.tail = nil
	if list.head != nil {
		var next *ListNode
		for list.head.next != nil {
			next = list.head.next
			list.head.next = nil
			list.head = next
		}
		list.head = nil
	}
}

func (list *List) ForEach(fn func(index int, elem *interface{})) {
	index := 0
	for node := list.head; node != nil; node = node.next {
		fn(index, &node.Value)
		index += 1
	}
}

func (list *List) String() string {
	str := "List {\n"
	for node := list.head; node != nil; node = node.next {
		str += "    " + fmt.Sprint(node.Value) + ",\n"
	}
	str += "}"
	return str
}
