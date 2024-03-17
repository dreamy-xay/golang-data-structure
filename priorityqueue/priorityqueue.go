/*
 * @Description: priorityqueue
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 22:08:21
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-16 14:58:48
 */
package priorityqueue

import "fmt"

type PriorityQueue struct {
	elements []interface{}
	rules    func(a, b *interface{}) bool
}

func NewPriorityQueue(rules func(a, b *interface{}) bool) PriorityQueue {
	return PriorityQueue{
		elements: make([]interface{}, 0),
		rules:    rules,
	}
}

func NewPriorityQueueCapacity(rules func(a, b *interface{}) bool, capacity int) *PriorityQueue {
	return &PriorityQueue{
		elements: make([]interface{}, 0, capacity),
		rules:    rules,
	}
}

func (pq *PriorityQueue) Init(arr []interface{}, rules func(a, b *interface{}) bool) {
	n := len(arr)
	pq.Clear()
	pq.rules = rules
	pq.elements = append(pq.elements, arr...)
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.elements = append(pq.elements, x)
	pq.up(len(pq.elements) - 1)
}

func (pq *PriorityQueue) Pop() interface{} {
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

func (pq *PriorityQueue) Top() (interface{}, bool) {
	if pq.Empty() {
		return nil, false
	}
	return pq.elements[0], true
}

func (pq *PriorityQueue) Size() int {
	return len(pq.elements)
}

func (pq *PriorityQueue) Empty() bool {
	return len(pq.elements) == 0
}

func (pq *PriorityQueue) Clear() {
	pq.elements = pq.elements[:0]
}

func (pq *PriorityQueue) String() string {
	str := "PriorityQueue {\n"
	for _, elem := range pq.elements {
		str += "    " + fmt.Sprint(elem) + ",\n"
	}
	str += "}"
	return str
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.rules(&pq.elements[j], &pq.elements[i]) {
			break
		}
		pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
		j = i
	}
}

func (pq *PriorityQueue) down(i0, n int) bool {
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
