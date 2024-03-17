/*
 * @Description: queue
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 12:43:34
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-08-10 16:49:23
 */
package queue

import "fmt"

type QueueNode[T any] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T any] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (queue *Queue[T]) Init(arr []T) {
	queue.Clear()
	for _, elem := range arr {
		queue.Push(elem)
	}
}

func (queue *Queue[T]) Size() int {
	return queue.size
}

func (queue *Queue[T]) Empty() bool {
	return queue.size == 0
}

func (queue *Queue[T]) Push(value T) {
	if queue.head == nil {
		queue.head = &QueueNode[T]{
			value: value,
			next:  nil,
		}
		queue.tail = queue.head
	} else {
		newnode := QueueNode[T]{
			value: value,
			next:  nil,
		}
		queue.tail.next = &newnode
		queue.tail = &newnode
	}
	queue.size += 1
}

func (queue *Queue[T]) Pop() any {
	if queue.head == nil {
		return nil
	}
	first := queue.head.value
	if queue.head == queue.tail {
		queue.head = nil
		queue.tail = nil
	} else {
		queue.head = queue.head.next
	}
	queue.size -= 1
	return first
}

func (queue *Queue[T]) Front() (any, bool) {
	if queue.head == nil {
		return nil, false
	}
	return queue.head.value, true
}

func (queue *Queue[T]) Clear() {
	queue.size = 0
	queue.tail = nil
	if queue.head != nil {
		var next *QueueNode[T]
		for queue.head.next != nil {
			next = queue.head.next
			queue.head.next = nil
			queue.head = next
		}
		queue.head = nil
	}
}

func (queue *Queue[T]) String() string {
	str := "Queue {\n"
	for node := queue.head; node != nil; node = node.next {
		str += "    " + fmt.Sprint(node.value) + ",\n"
	}
	str += "}"
	return str
}
