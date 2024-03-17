/*
 * @Description: priorityqueue
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 22:08:21
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-08-06 15:59:19
 */
package priorityqueue

import "fmt"

type PriorityQueue[T any] struct {
	elements []T
	rules    func(a, b *T) bool
}

func NewPriorityQueue[T any](rules func(a, b *T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		elements: make([]T, 0),
		rules:    rules,
	}
}

func NewPriorityQueueCapacity[T any](rules func(a, b *T) bool, capacity int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		elements: make([]T, 0, capacity),
		rules:    rules,
	}
}

func (pq *PriorityQueue[T]) Init(arr []T, rules func(a, b *T) bool) {
	n := len(arr)
	pq.Clear()
	pq.rules = rules
	pq.elements = append(pq.elements, arr...)
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.elements = append(pq.elements, x)
	pq.up(len(pq.elements) - 1)
}

func (pq *PriorityQueue[T]) Pop() any {
	n := len(pq.elements) - 1
	if n < 0 {
		return nil
	}
	pq.elements[0], pq.elements[n] = pq.elements[n], pq.elements[0]
	pq.down(0, n)
	last := pq.elements[n]
	pq.elements = pq.elements[:n]
	return last
}

func (pq *PriorityQueue[T]) Top() (any, bool) {
	if pq.Empty() {
		return nil, false
	}
	return pq.elements[0], true
}

func (pq *PriorityQueue[T]) Size() int {
	return len(pq.elements)
}

func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.elements) == 0
}

func (pq *PriorityQueue[T]) Clear() {
	pq.elements = pq.elements[:0]
}

func (pq *PriorityQueue[T]) String() string {
	str := "PriorityQueue {\n"
	for _, elem := range pq.elements {
		str += "    " + fmt.Sprint(elem) + ",\n"
	}
	str += "}"
	return str
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.rules(&pq.elements[j], &pq.elements[i]) {
			break
		}
		pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
		j = i
	}
}

func (pq *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.rules(&pq.elements[j2], &pq.elements[j1]) {
			j = j2
		}
		if !pq.rules(&pq.elements[j], &pq.elements[i]) {
			break
		}
		pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
		i = j
	}
	return i > i0
}
