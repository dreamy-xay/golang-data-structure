/*
 * @Description: vector
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 15:47:27
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 21:06:29
 */
package vector

import (
	. "collections/common"
	"fmt"
	"sort"
)

type Vector[T any] []T

func NewVector[T any]() Vector[T] {
	return make(Vector[T], 0)
}

func NewVectorCapacity[T any](capacity int) Vector[T] {
	return make(Vector[T], 0, capacity)
}

func (vec *Vector[T]) Init(arr []T) {
	vec.Clear()
	*vec = append(*vec, arr...)
}

func (vec *Vector[T]) Size() int {
	return len(*vec)
}

func (vec *Vector[T]) Empty() bool {
	return len(*vec) == 0
}

func (vec *Vector[T]) At(index int) (any, bool) {
	n := len(*vec)
	if index >= n || index < 0 {
		return nil, false
	}
	return (*vec)[index], true
}

func (vec *Vector[T]) Insert(index int, value T) {
	n := len(*vec)
	if index > n || index < 0 {
		return
	}
	*vec = append(*vec, value)
	copy((*vec)[index+1:], (*vec)[index:])
	(*vec)[index] = value
}

func (vec *Vector[T]) PushBack(value T) {
	*vec = append(*vec, value)
}

func (vec *Vector[T]) PushFront(value T) {
	vec.Insert(0, value)
}

func (vec *Vector[T]) Erase(index int) any {
	n := len(*vec)
	if index >= n || index < 0 {
		return nil
	}
	target := (*vec)[index]
	*vec = append((*vec)[:index], (*vec)[index+1:]...)
	return target
}

func (vec *Vector[T]) PopBack() any {
	n := len(*vec) - 1
	if n < 0 {
		return nil
	}
	last := (*vec)[n]
	*vec = (*vec)[:n]
	return last
}

func (vec *Vector[T]) PopFront() any {
	n := len(*vec)
	if n <= 0 {
		return nil
	}
	first := (*vec)[0]
	*vec = (*vec)[1:]
	return first
}

func (vec *Vector[T]) Front() (any, bool) {
	return vec.At(0)
}

func (vec *Vector[T]) Back() (any, bool) {
	return vec.At(len(*vec) - 1)
}

func (vec *Vector[T]) Reverse() {
	n := len(*vec)
	part := n / 2
	n -= 1
	for i := 0; i < part; i += 1 {
		(*vec)[i], (*vec)[n-i] = (*vec)[n-i], (*vec)[i]
	}
}

func (vec *Vector[T]) Sort(fn func(a, b *T) bool) {
	sort.Slice((*vec), func(i, j int) bool {
		return fn(&(*vec)[i], &(*vec)[j])
	})
}

func (vec *Vector[T]) Unique(equal func(a, b *T) bool) {
	n := len(*vec)
	i := 0
	for j := 0; j < n-1; j += 1 {
		if !equal(&(*vec)[j], &(*vec)[j+1]) { // 如果后边的不等于前面的就把后边的往前添加
			i += 1
			(*vec)[i] = (*vec)[j+1] // 把后边的往前添加
		}
	}
	*vec = (*vec)[:i+1]
}

func (vec *Vector[T]) BinarySearch(target T, rules, equal func(a, b *T) bool) int {
	return vec.BinarySearchByRange(NewRange(0, len(*vec)), target, rules, equal)
}

func (vec *Vector[T]) LowerBound(target T, rules func(a, b *T) bool) int {
	return vec.LowerBoundByRange(NewRange(0, len(*vec)), target, rules)
}

func (vec *Vector[T]) UpperBound(target T, rules func(a, b *T) bool) int {
	return vec.UpperBoundByRange(NewRange(0, len(*vec)), target, rules)
}

func (vec *Vector[T]) BinarySearchByRange(rg Range, target T, rules, equal func(a, b *T) bool) int {
	rg.Limit(0, len(*vec)) // 范围修正

	var mid int         // 中间值
	left := rg.Start    // 左边界
	right := rg.End - 1 // 右边界

	for left <= right { // 循环条件
		mid = left + ((right - left) >> 1) // 取中间值

		if equal(&(*vec)[mid], &target) { // 相等就返回
			return mid
		}
		if rules(&target, &(*vec)[mid]) { // 目标值小于中间值就向左找
			right = mid - 1
		} else { // 目标值大于中间值就向右找
			left = mid + 1
		}
	}
	return -1
}

func (vec *Vector[T]) LowerBoundByRange(rg Range, target T, rules func(a, b *T) bool) int {
	rg.Limit(0, len(*vec)) // 范围修正

	var mid int         // 中间值
	left := rg.Start    // 左边界
	right := rg.End - 1 // 右边界

	for left <= right { // 循环条件
		mid = left + ((right - left) >> 1) // 取中间值

		if rules(&(*vec)[mid], &target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left < rg.End && left > -1 { // 判断是否越界
		return left
	} else { // 越界就返回-1
		return -1
	}
}

func (vec *Vector[T]) UpperBoundByRange(rg Range, target T, rules func(a, b *T) bool) int {
	rg.Limit(0, len(*vec)) // 范围修正

	var mid int         // 中间值
	left := rg.Start    // 左边界
	right := rg.End - 1 // 右边界

	for left <= right { // 循环条件
		mid = left + ((right - left) >> 1) // 取中间值

		if rules(&target, &(*vec)[mid]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left < rg.End && left > -1 { // 判断是否越界
		return left
	} else { // 越界就返回-1
		return -1
	}
}

func (vec *Vector[T]) Clear() {
	*vec = (*vec)[:0]
}

func (vec *Vector[T]) Fill(start, end int, value T) {
	if start >= end {
		return
	}
	start = Max(start, 0)
	end = Min(end, len(*vec))
	for i := start; i < end; i += 1 {
		(*vec)[i] = value
	}
}

func (vec *Vector[T]) GetIterator() *Iterator[T] {
	current := -1
	vectorLen := len(*vec)
	return NewIterator[T](func() bool {
		return current+1 < vectorLen
	}, func() *T {
		current += 1
		if current < vectorLen {
			return &(*vec)[current]
		} else {
			return nil
		}
	})
}

func (vec *Vector[T]) String() string {
	str := "Vector {\n"
	for _, elem := range *vec {
		str += "    " + fmt.Sprint(elem) + ",\n"
	}
	str += "}"
	return str
}
