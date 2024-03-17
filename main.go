/*
 * @Description:
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 20:57:22
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-20 17:29:21
 */
package main

import (
	"collections/avltree"
	"collections/hashmap"
	"collections/hashset"
	"collections/list"
	"collections/priorityqueue"
	"collections/queue"
	"collections/rbtree"
	"collections/rbtree_at"
	"collections/stack"
	"collections/vector"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var less = func(a, b *interface{}) bool {
		return (*a).(int) < (*b).(int)
	}
	var greater = func(a, b *interface{}) bool {
		return (*a).(int) > (*b).(int)
	}
	var equal = func(a, b *interface{}) bool {
		return (*a).(int) == (*b).(int)
	}

	// hashmap test
	mp := hashmap.NewHashMap()
	mp.Insert("hello", 0)
	mp.Insert("world", 0)
	fmt.Println(mp.Find("world"))
	fmt.Println(mp.String())

	// hashset test
	st := hashset.NewHashSet()
	st.Insert("9873987")
	st.Insert(903920323)
	st.Insert("sahskjahskajsh")
	st.ForEach(func(key *interface{}) {
		fmt.Println(*key)
	})
	fmt.Println(st.String())

	// priorityqueue
	pq := priorityqueue.PriorityQueue{}
	pq.Init([]interface{}{}, less)
	pq.Push(9)
	pq.Push(67)
	pq.Push(78)
	pq.Push(-1)
	pq.Push(3)
	pq.Push(56)
	fmt.Println(pq.String())
	fmt.Println(pq.Empty())
	pq.Pop()
	for !pq.Empty() {
		fmt.Println(pq.Pop())
	}

	// stack
	stk := stack.Stack{}
	stk.Init([]interface{}{0, 3, 6, 8, 9, 11})
	haspush := true
	fmt.Println(stk.String())
	for !stk.Empty() {
		if top, ok := stk.Top(); haspush && ok && top == 6 {
			stk.Push(3728)
			haspush = false
		}
		fmt.Println(stk.Pop())
	}

	// queue
	qu := queue.NewQueue()
	qu.Init([]interface{}{4, 8, 2, 1})
	fmt.Println(qu.Size())
	qu.Push(6565)
	qu.Push(222)
	fmt.Println(qu.String())
	for !qu.Empty() {
		front, ok := qu.Front()
		fmt.Println(ok, front, qu.Size())
		qu.Pop()
	}

	// list
	lt := list.NewList()
	lt.Init([]interface{}{1, 6, 0, 32, 5})
	lt.PushFront(90)
	lt.PopBack()
	lt.PushBack(78)
	fmt.Println(lt.String())
	fmt.Println(lt.Size())
	lt.Reverse()
	fmt.Println(lt.String())
	lt.Sort(less)
	fmt.Println(lt.String())

	// vector
	vec := vector.NewVectorCapacity(10)
	vec.Init([]interface{}{1, 5, 9, 2, 0, 3})
	fmt.Println(vec.String())
	vec.Sort(less)
	fmt.Println(vec.String())
	vec.Reverse()
	fmt.Println(vec.String())
	vec.Insert(0, 111)
	fmt.Println(vec.String())
	vec.Insert(5, 234)
	fmt.Println(vec.String())
	vec.Insert(vec.Size(), 999)
	fmt.Println(vec.String())
	vec.Erase(0)
	fmt.Println(vec.String())
	vec.Erase(vec.Size() - 1)
	fmt.Println(vec.String())
	vec.Erase(3)
	fmt.Println(vec.String())
	vec.PopBack()
	fmt.Println(vec.String())

	// avltree
	tr := avltree.NewAVLTree(less, equal)
	tr.Init([]interface{}{1, 7, 0, 9, 23, -1, 45, -9, 21}, greater, equal)
	fmt.Println(tr.String())

	fmt.Println(1, ":", tr.Order(1))
	fmt.Println(7, ":", tr.Order(7))
	fmt.Println(0, ":", tr.Order(0))
	fmt.Println(9, ":", tr.Order(9))
	fmt.Println(23, ":", tr.Order(23))
	fmt.Println(-1, ":", tr.Order(-1))
	fmt.Println(45, ":", tr.Order(45))
	fmt.Println(-9, ":", tr.Order(-9))
	fmt.Println(21, ":", tr.Order(21))
	fmt.Println(-2, ":", tr.Order(-2))
	fmt.Println(49, ":", tr.Order(49))
	fmt.Println(-30, ":", tr.Order(-30))

	fmt.Println(0, ":", tr.At(0))
	fmt.Println(1, ":", tr.At(1))
	fmt.Println(2, ":", tr.At(2))
	fmt.Println(3, ":", tr.At(3))
	fmt.Println(4, ":", tr.At(4))
	fmt.Println(5, ":", tr.At(5))
	fmt.Println(6, ":", tr.At(6))
	fmt.Println(7, ":", tr.At(7))
	fmt.Println(8, ":", tr.At(8))
	fmt.Println(9, ":", tr.At(9))
	fmt.Println(10, ":", tr.At(10))
	fmt.Println(-1, ":", tr.At(-1))

	fmt.Println("lowerbound 0:", tr.LowerBound(0))
	fmt.Println("lowerbound 44:", tr.LowerBound(44))
	fmt.Println("lowerbound 12:", tr.LowerBound(12))
	fmt.Println("lowerbound -5:", tr.LowerBound(-5))
	fmt.Println("lowerbound -100:", tr.LowerBound(-100))
	fmt.Println("lowerbound 27:", tr.LowerBound(27))
	fmt.Println("lowerbound 6:", tr.LowerBound(6))
	fmt.Println("lowerbound 9:", tr.LowerBound(9))
	fmt.Println("lowerbound 22:", tr.LowerBound(22))
	fmt.Println("lowerbound 333:", tr.LowerBound(333))
	fmt.Println("lowerbound 46:", tr.LowerBound(46))
	fmt.Println("lowerbound 3:", tr.LowerBound(3))

	fmt.Println("upperbound 0:", tr.UpperBound(0))
	fmt.Println("upperbound 44:", tr.UpperBound(44))
	fmt.Println("upperbound 12:", tr.UpperBound(12))
	fmt.Println("upperbound -5:", tr.UpperBound(-5))
	fmt.Println("upperbound -100:", tr.UpperBound(-100))
	fmt.Println("upperbound 27:", tr.UpperBound(27))
	fmt.Println("upperbound 6:", tr.UpperBound(6))
	fmt.Println("upperbound 9:", tr.UpperBound(9))
	fmt.Println("upperbound 22:", tr.UpperBound(22))
	fmt.Println("upperbound 333:", tr.UpperBound(333))
	fmt.Println("upperbound 46:", tr.UpperBound(46))
	fmt.Println("upperbound 3:", tr.UpperBound(3))

	node, ok := tr.Find(23)
	fmt.Println(tr.Size(), tr.ForSize(), ok, node)
	for it := tr.Begin(); it != nil; it = it.GetNext() {
		fmt.Println(it.Value)
	}
	tr.Erase(1)
	for it := tr.End(); it != nil; it = it.GetPre() {
		fmt.Println(it.Value)
	}
	fmt.Println(tr.Size())
	vec.Clear()
	fmt.Println("load...")
	rand.Seed(time.Now().Unix())
	trNum := 10000000
	for i := 0; i < trNum; i += 1 {
		vec.PushBack(rand.Intn(10000))
	}
	vec_rb := vector.NewVector()
	vec_rb.Init(vec)
	fmt.Println("run...")
	start := time.Now()
	tr.Init(vec, less, equal)
	cost := time.Since(start)
	fmt.Printf("avl cost=[%s]\n", cost)
	fmt.Println("ok...", tr.ForSize(), tr.Size())

	start = time.Now()
	vec.Sort(less)
	cost = time.Since(start)
	fmt.Printf("sort cost=[%s]\n", cost)

	start = time.Now()
	isOk := true
	for i := 0; i < trNum; i += 1 {
		vecValue := vec[i]
		trValue := tr.At(i).Value
		if vecValue != trValue {
			fmt.Println(vecValue, trValue, vecValue == trValue)
			isOk = false
			break
		}
	}
	cost = time.Since(start)
	fmt.Printf("vec tr compare cost=[%s]\n", cost)
	if isOk {
		fmt.Println("all ok...")
	} else {
		fmt.Println("error!!!")
	}

	// rbtree
	rbtr := rbtree.NewRBTree(less, equal)
	rbtr.Init([]interface{}{1, 7, 0, 9, 23, -1, 45, -9, 21}, greater, equal)
	fmt.Println(rbtr.Size())
	rbtr.ForEach(func(value *interface{}) {
		fmt.Println(*value)
		// time.Sleep(time.Duration(2) * time.Second)
	})
	fmt.Println(rbtr.String())

	fmt.Println("run...")
	start = time.Now()
	rbtr.Init(vec_rb, less, equal)
	cost = time.Since(start)
	fmt.Printf("rb cost=[%s]\n", cost)
	fmt.Println("ok...", rbtr.ForSize(), rbtr.Size())
	vec_rb_i := 0
	isOk = true
	rbtr.ForEach(func(value *interface{}) {
		if *value != vec[vec_rb_i] {
			isOk = false
		}
		vec_rb_i += 1
	})
	fmt.Println("isOk:", isOk)

	fmt.Println("lowerbound 0:", rbtr.LowerBound(0))
	fmt.Println("lowerbound 44:", rbtr.LowerBound(44))
	fmt.Println("lowerbound 12:", rbtr.LowerBound(12))
	fmt.Println("lowerbound -5:", rbtr.LowerBound(-5))
	fmt.Println("lowerbound -100:", rbtr.LowerBound(-100))
	fmt.Println("lowerbound 27:", rbtr.LowerBound(27))
	fmt.Println("lowerbound 6:", rbtr.LowerBound(6))
	fmt.Println("lowerbound 9:", rbtr.LowerBound(9))
	fmt.Println("lowerbound 22:", rbtr.LowerBound(22))
	fmt.Println("lowerbound 333:", rbtr.LowerBound(333))
	fmt.Println("lowerbound 46:", rbtr.LowerBound(46))
	fmt.Println("lowerbound 3:", rbtr.LowerBound(3))

	fmt.Println("upperbound 0:", rbtr.UpperBound(0))
	fmt.Println("upperbound 44:", rbtr.UpperBound(44))
	fmt.Println("upperbound 12:", rbtr.UpperBound(12))
	fmt.Println("upperbound -5:", rbtr.UpperBound(-5))
	fmt.Println("upperbound -100:", rbtr.UpperBound(-100))
	fmt.Println("upperbound 27:", rbtr.UpperBound(27))
	fmt.Println("upperbound 6:", rbtr.UpperBound(6))
	fmt.Println("upperbound 9:", rbtr.UpperBound(9))
	fmt.Println("upperbound 22:", rbtr.UpperBound(22))
	fmt.Println("upperbound 333:", rbtr.UpperBound(333))
	fmt.Println("upperbound 46:", rbtr.UpperBound(46))
	fmt.Println("upperbound 3:", rbtr.UpperBound(3))

	// rbtree_at
	rbtr_at := rbtree_at.NewRBTree(greater, equal)
	rbtr_at.Init([]interface{}{1, 7, 0, 9, 23, -1, 45, -9, 21}, less, equal)
	fmt.Println(rbtr_at.Size(), rbtr_at.ForSize())
	rbtr_at.ForEach(func(value *interface{}) {
		fmt.Println(*value)
		// time.Sleep(time.Duration(2) * time.Second)
	})
	fmt.Println(rbtr_at.String())
	fmt.Println(1, ":", rbtr_at.Order(1))
	fmt.Println(7, ":", rbtr_at.Order(7))
	fmt.Println(0, ":", rbtr_at.Order(0))
	fmt.Println(9, ":", rbtr_at.Order(9))
	fmt.Println(23, ":", rbtr_at.Order(23))
	fmt.Println(-1, ":", rbtr_at.Order(-1))
	fmt.Println(45, ":", rbtr_at.Order(45))
	fmt.Println(-9, ":", rbtr_at.Order(-9))
	fmt.Println(21, ":", rbtr_at.Order(21))
	fmt.Println(-2, ":", rbtr_at.Order(-2))
	fmt.Println(49, ":", rbtr_at.Order(49))
	fmt.Println(-30, ":", rbtr_at.Order(-30))

	fmt.Println(0, ":", rbtr_at.At(0))
	fmt.Println(1, ":", rbtr_at.At(1))
	fmt.Println(2, ":", rbtr_at.At(2))
	fmt.Println(3, ":", rbtr_at.At(3))
	fmt.Println(4, ":", rbtr_at.At(4))
	fmt.Println(5, ":", rbtr_at.At(5))
	fmt.Println(6, ":", rbtr_at.At(6))
	fmt.Println(7, ":", rbtr_at.At(7))
	fmt.Println(8, ":", rbtr_at.At(8))
	fmt.Println(9, ":", rbtr_at.At(9))
	fmt.Println(10, ":", rbtr_at.At(10))
	fmt.Println(-1, ":", rbtr_at.At(-1))

	fmt.Println("run...")
	start = time.Now()
	rbtr_at.Init(vec_rb, less, equal)
	cost = time.Since(start)
	fmt.Printf("rb cost=[%s]\n", cost)
	fmt.Println("ok...", rbtr_at.ForSize(), rbtr_at.Size())
	vec_rb_i = 0
	isOk = true
	rbtr_at.ForEach(func(value *interface{}) {
		if *value != vec[vec_rb_i] {
			isOk = false
		}
		vec_rb_i += 1
	})
	fmt.Println("isOk:", isOk)
}
