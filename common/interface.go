/*
 * @Description: common interface
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2023-08-06 15:01:45
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 14:24:45
 */

package common

// Comparable 代表所有可判断是否相等不等的类型
type Comparable interface {
	Ordered | Complex | Boolean
}

// Ordered 代表所有可比大小排序的类型
type Ordered interface {
	Number | String
}

type Number interface {
	Integer | Float
}

type Complex interface {
	~complex64 | ~complex128
}

type Boolean interface {
	~bool
}

type String interface {
	~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}
