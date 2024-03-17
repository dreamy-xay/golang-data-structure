/*
 * @Description: stack
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-11 22:03:42
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-12 15:43:09
 */
package stack

import "fmt"

type Stack struct {
	elements []interface{}
}

func NewStack() Stack {
	return Stack{
		elements: make([]interface{}, 0),
	}
}

func NewStackCapacity(capacity int) *Stack {
	return &Stack{
		elements: make([]interface{}, 0, capacity),
	}
}

func (stack *Stack) Init(arr []interface{}) {
	stack.Clear()
	stack.elements = append(stack.elements, arr...)
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}

func (stack *Stack) Empty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.elements = append(stack.elements, value)
}

func (stack *Stack) Pop() interface{} {
	n := len(stack.elements) - 1
	if n < 0 {
		return nil
	}
	last := stack.elements[n]
	stack.elements = stack.elements[:n]
	return last
}

func (stack *Stack) Top() (interface{}, bool) {
	n := len(stack.elements)
	if n == 0 {
		return nil, false
	}
	return stack.elements[n-1], true
}

func (stack *Stack) Clear() {
	stack.elements = stack.elements[:0]
}

func (stack *Stack) String() string {
	str := "Stack {\n"
	for _, elem := range stack.elements {
		str += "    " + fmt.Sprint(elem) + ",\n"
	}
	str += "}"
	return str
}
