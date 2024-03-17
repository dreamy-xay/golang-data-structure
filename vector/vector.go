/*
 * @Description: vector
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 15:47:27
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-14 14:44:26
 */
package vector

import (
	"fmt"
	"sort"
)

type Vector []interface{}

func NewVector() Vector {
	return make(Vector, 0)
}

func NewVectorCapacity(capacity int) Vector {
	return make(Vector, 0, capacity)
}

func (vec *Vector) Init(arr []interface{}) {
	vec.Clear()
	*vec = append(*vec, arr...)
}

func (vec *Vector) Size() int {
	return len(*vec)
}

func (vec *Vector) Empty() bool {
	return len(*vec) == 0
}

func (vec *Vector) At(index int) (interface{}, bool) {
	n := len(*vec)
	if index >= n || index < 0 {
		return nil, false
	}
	return (*vec)[index], true
}

func (vec *Vector) Insert(index int, value interface{}) {
	n := len(*vec)
	if index > n || index < 0 {
		return
	}
	*vec = append(*vec, 0)
	copy((*vec)[index+1:], (*vec)[index:])
	(*vec)[index] = value
}

func (vec *Vector) PushBack(value interface{}) {
	*vec = append(*vec, value)
}

func (vec *Vector) PushFront(value interface{}) {
	vec.Insert(0, value)
}

func (vec *Vector) Erase(index int) interface{} {
	n := len(*vec)
	if index >= n || index < 0 {
		return nil
	}
	target := (*vec)[index]
	*vec = append((*vec)[:index], (*vec)[index+1:]...)
	return target
}

func (vec *Vector) PopBack() interface{} {
	n := len(*vec) - 1
	if n < 0 {
		return nil
	}
	last := (*vec)[n]
	*vec = (*vec)[:n]
	return last
}

func (vec *Vector) PopFront() interface{} {
	n := len(*vec)
	if n <= 0 {
		return nil
	}
	first := (*vec)[0]
	*vec = (*vec)[1:]
	return first
}

func (vec *Vector) Front() (interface{}, bool) {
	return vec.At(0)
}

func (vec *Vector) Back() (interface{}, bool) {
	return vec.At(len(*vec) - 1)
}

func (vec *Vector) Reverse() {
	n := len(*vec)
	part := n / 2
	n -= 1
	for i := 0; i < part; i += 1 {
		(*vec)[i], (*vec)[n-i] = (*vec)[n-i], (*vec)[i]
	}
}

func (vec *Vector) Sort(fn func(a, b *interface{}) bool) {
	sort.Slice((*vec), func(i, j int) bool {
		return fn(&(*vec)[i], &(*vec)[j])
	})
}

func (vec *Vector) Clear() {
	*vec = (*vec)[:0]
}

func (vec *Vector) String() string {
	str := "Vector {\n"
	for _, elem := range *vec {
		str += "    " + fmt.Sprint(elem) + ",\n"
	}
	str += "}"
	return str
}
