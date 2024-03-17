/*
 * @Description: common function
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2023-08-06 15:01:45
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-06 17:33:44
 */

package common

import "unsafe"

func NewRange(start, end int) Range {
	if start >= end {
		panic("Start must be less than End.")
	}
	return Range{Start: start, End: end}
}

func NewPair[KEY any, VALUE any](key KEY, value VALUE) *Pair[KEY, VALUE] {
	return &Pair[KEY, VALUE]{Key: key, Value: value}
}

func NewIterator[T any](hasNext func() bool, next func() *T) *Iterator[T] {
	return &Iterator[T]{HasNext: hasNext, Next: next}
}

func Less[T Ordered](a, b *T) bool {
	return *a < *b
}

func Greater[T Ordered](a, b *T) bool {
	return *a > *b
}

func Equal[T Comparable](a, b *T) bool {
	return *a == *b
}

func LessEqual[T Ordered](a, b *T) bool {
	return *a <= *b
}

func GreaterEqual[T Ordered](a, b *T) bool {
	return *a >= *b
}

func Abs[T Number](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func Max[T Number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MaxElement[T Ordered](iterator *Iterator[T]) *T {
	// 定义变量current和max，初始值为nil
	var current *T
	var max *T = nil
	// 遍历迭代器
	for iterator.HasNext() {
		// 获取当前迭代器的下一个元素
		current = iterator.Next()
		// 如果max为空，或者当前元素比max小，则更新max
		if max == nil || *max < *current {
			max = current
		}
	}
	// 返回max
	return max
}

func MinElement[T Ordered](iterator *Iterator[T]) *T {
	// 定义变量current，并赋值为空
	var current *T
	// 定义变量min，并赋值为nil
	var min *T = nil
	// 遍历迭代器
	for iterator.HasNext() {
		// 获取当前迭代器的下一个元素
		current = iterator.Next()
		// 如果min为空，或者当前元素小于min，则将当前元素赋值给min
		if min == nil || *min > *current {
			min = current
		}
	}
	// 返回min
	return min
}

func MaxElementByFunction[T any](iterator *Iterator[T], compare func(a, b *T) bool) *T {
	var current *T
	var max *T = nil
	for iterator.HasNext() {
		current = iterator.Next()
		if max == nil || compare(current, max) {
			max = current
		}
	}
	return max
}

func MinElementByFunction[T any](iterator *Iterator[T], compare func(a, b *T) bool) *T {
	var current *T
	var min *T = nil
	for iterator.HasNext() {
		current = iterator.Next()
		if min == nil || compare(current, min) {
			min = current
		}
	}
	return min
}

func Accumulate[T Number](iterator *Iterator[T], value T) T {
	// 初始化求和值
	sum := value
	// 循环遍历求和值
	for iterator.HasNext() {
		sum += *iterator.Next()
	}
	// 返回求和值
	return sum
}

func GetIndexByArrayPointer[T any](arr []T, pointer *T) int {
	if len(arr) == 0 || pointer == nil { // 数组为空或指针为空
		return -1
	}
	size := unsafe.Sizeof(arr[0])             // 数组元素大小
	start := uintptr(unsafe.Pointer(&arr[0])) // 数组起始位置
	end := uintptr(unsafe.Pointer(pointer))   // 数组指针位置
	return int((end - start) / size)
}
