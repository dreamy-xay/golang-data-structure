/*
 * @Description: queue
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 12:43:34
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-12 15:33:17
 */
package queue

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type Queue struct {
	head *node
	tail *node
	size int
}

func NewQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (queue *Queue) Init(arr []interface{}) {
	queue.Clear()
	for _, elem := range arr {
		queue.Push(elem)
	}
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) Empty() bool {
	return queue.size == 0
}

func (queue *Queue) Push(value interface{}) {
	if queue.head == nil {
		queue.head = &node{
			value: value,
			next:  nil,
		}
		queue.tail = queue.head
	} else {
		newnode := node{
			value: value,
			next:  nil,
		}
		queue.tail.next = &newnode
		queue.tail = &newnode
	}
	queue.size += 1
}

func (queue *Queue) Pop() interface{} {
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

func (queue *Queue) Front() (interface{}, bool) {
	if queue.head == nil {
		return nil, false
	}
	return queue.head.value, true
}

func (queue *Queue) Clear() {
	queue.size = 0
	queue.tail = nil
	if queue.head != nil {
		var next *node
		for queue.head.next != nil {
			next = queue.head.next
			queue.head.next = nil
			queue.head = next
		}
		queue.head = nil
	}
}

func (queue *Queue) String() string {
	str := "Queue {\n"
	for node := queue.head; node != nil; node = node.next {
		str += "    " + fmt.Sprint(node.value) + ",\n"
	}
	str += "}"
	return str
}
