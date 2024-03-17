/*
 * @Description: stack
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-11 22:03:42
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-08-06 16:07:42
 */
package stack

import "fmt"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

func NewStackCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0, capacity),
	}
}

func (stack *Stack[T]) Init(arr []T) {
	stack.Clear()
	stack.elements = append(stack.elements, arr...)
}

func (stack *Stack[T]) Size() int {
	return len(stack.elements)
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack[T]) Push(value T) {
	stack.elements = append(stack.elements, value)
}

func (stack *Stack[T]) Pop() any {
	n := len(stack.elements) - 1
	if n < 0 {
		return nil
	}
	last := stack.elements[n]
	stack.elements = stack.elements[:n]
	return last
}

func (stack *Stack[T]) Top() (any, bool) {
	n := len(stack.elements)
	if n == 0 {
		return nil, false
	}
	return stack.elements[n-1], true
}

func (stack *Stack[T]) Clear() {
	stack.elements = stack.elements[:0]
}

func (stack *Stack[T]) String() string {
	str := "Stack {\n"
	for _, elem := range stack.elements {
		str += "    " + fmt.Sprint(elem) + ",\n"
	}
	str += "}"
	return str
}
