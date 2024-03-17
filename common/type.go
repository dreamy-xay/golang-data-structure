/*
 * @Description: common type
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2023-08-06 15:01:45
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 16:30:25
 */

package common

import "fmt"

type Range struct {
	Start int
	End   int
}

func (rg *Range) Check() *Range {
	if rg.Start >= rg.End {
		panic("Start must be less than End.")
	}
	return rg
}

func (rg *Range) Limit(min, max int) *Range {
	rg.Start = Max(rg.Start, min)
	rg.End = Min(rg.End, max)
	return rg
}

func (rg *Range) String() string {
	return fmt.Sprintf("[%v, %v)", rg.Start, rg.End)
}

type Pair[KEY any, VALUE any] struct {
	Key   KEY
	Value VALUE
}

func (pair *Pair[KEY, VALUE]) String() string {
	return fmt.Sprintf("{%v: %v}", pair.Key, pair.Value)
}

type Iterator[T any] struct {
	HasNext func() bool
	Next    func() *T
}
