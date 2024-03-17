/*
 * @Description:
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 20:57:22
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 20:35:47
 */
package main

import (
	"collections/avltree"
	. "collections/common"
	"collections/hashmap"
	"collections/hashset"
	"collections/list"
	"collections/priorityqueue"
	"collections/queue"
	"collections/rbtree"
	"collections/stack"
	"collections/vector"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("start test >>>>>>>>>>>>>>>>>>>>>")

	// hashmap test
	fmt.Println("\nhashmap test >>>>>>>>>>>>>>>>>>>>>")
	mp := hashmap.NewHashMap[string, int]()
	mp.Insert("hello", 0)
	mp.Insert("world", 1)
	fmt.Println(mp.Find("world"))
	fmt.Println(MinElementByFunction(mp.GetIterator(), func(a, b *Pair[string, int]) bool {
		return a.Value < b.Value
	}))
	fmt.Println(MaxElementByFunction(mp.GetIterator(), func(a, b *Pair[string, int]) bool {
		return a.Value > b.Value
	}))
	fmt.Println(&mp)

	// hashset test
	fmt.Println("\nhashset test >>>>>>>>>>>>>>>>>>>>>")
	st := hashset.NewHashSet[string]()
	st.Insert("9873987")
	st.Insert("903920323")
	st.Insert("sahskjahskajsh")
	st.ForEach(func(key *string) {
		fmt.Println(*key)
	})
	st_iter := st.GetIterator()
	for st_iter.HasNext() {
		fmt.Println(*st_iter.Next())
	}
	fmt.Println(&st)
	fmt.Println("min string =", *MinElementByFunction(st.GetIterator(), Less[string]))
	fmt.Println(st.ToMap())

	// priorityqueue
	fmt.Println("\npriorityqueue test >>>>>>>>>>>>>>>>>>>>>")
	pq := priorityqueue.PriorityQueue[int]{}
	pq.Init([]int{}, Greater[int])
	pq.Push(9)
	pq.Push(67)
	pq.Push(78)
	pq.Push(-1)
	pq.Push(3)
	pq.Push(56)
	fmt.Println(&pq)
	fmt.Println(pq.Empty())
	pq.Pop()
	for !pq.Empty() {
		fmt.Println(pq.Pop())
	}

	// stack
	fmt.Println("\nstack test >>>>>>>>>>>>>>>>>>>>>")
	stk := stack.Stack[int]{}
	stk.Init([]int{0, 3, 6, 8, 9, 11})
	haspush := true
	fmt.Println(&stk)
	for !stk.Empty() {
		if top, ok := stk.Top(); haspush && ok && top == 6 {
			stk.Push(3728)
			haspush = false
		}
		fmt.Println(stk.Pop())
	}

	// queue
	fmt.Println("\nqueue test >>>>>>>>>>>>>>>>>>>>>")
	qu := queue.NewQueue[int]()
	qu.Init([]int{4, 8, 2, 1, 7878, 4564})
	fmt.Println(qu.Size())
	qu.Clear()
	front, ok := qu.Front()
	fmt.Println(front, ok, qu.Size())
	qu.Push(6565)
	qu.Push(222)
	qu.Push(3232)
	qu.Push(69905)
	fmt.Println(&qu)
	for !qu.Empty() {
		front, ok := qu.Front()
		fmt.Println(ok, front, qu.Size())
		qu.Pop()
	}

	// list
	fmt.Println("\nlist test >>>>>>>>>>>>>>>>>>>>>")
	lt := list.NewList[int]()
	lt.Init([]int{1, 6, 0, 32, 5})
	for it := lt.Begin(); it != nil; it = it.GetNext() {
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	for it := lt.End(); it != nil; it = it.GetPre() {
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	fmt.Println("sum =", Accumulate(lt.GetIterator(), 0))
	lt.PushFront(90)
	lt.PopBack()
	lt.PushBack(78)
	fmt.Println(lt)
	fmt.Println(lt.Size())
	lt.Reverse()
	fmt.Println(lt)
	lt.Sort(Less[int])
	fmt.Println(lt)

	// vector
	fmt.Println("\nvector test >>>>>>>>>>>>>>>>>>>>>")
	vec := vector.NewVectorCapacity[int](10)
	vec.Init([]int{1, 5, 9, 2, 0, 3, 9, 5, 2, 5, 0, 5})
	fmt.Println(&vec)
	vec.Sort(Greater[int])
	fmt.Println(&vec)
	fmt.Println("5=", vec.LowerBound(5, Greater[int]), "9=", vec.LowerBound(9, Greater[int]), "10=", vec.LowerBound(-1, Greater[int]))
	fmt.Println("5=", vec.UpperBound(5, Greater[int]), "9=", vec.UpperBound(9, Greater[int]), "10=", vec.UpperBound(-1, Greater[int]))
	vec.Unique(Equal[int])
	fmt.Println(&vec)
	fmt.Println(vec.BinarySearch(3, Greater[int], Equal[int]), vec.BinarySearch(10, Greater[int], Equal[int]))
	vec.Reverse()
	fmt.Println(&vec)
	fmt.Println("5=", vec.LowerBound(5, Less[int]), "9=", vec.LowerBound(9, Less[int]), "10=", vec.LowerBound(10, Less[int]))
	fmt.Println("5=", vec.UpperBound(5, Less[int]), "9=", vec.UpperBound(9, Less[int]), "10=", vec.UpperBound(10, Less[int]))
	vec.Insert(0, 111)
	fmt.Println(&vec)
	vec.Insert(5, 234)
	fmt.Println(&vec)
	vec.Insert(vec.Size(), 999)
	fmt.Println(&vec)
	vec.Erase(0)
	fmt.Println(&vec)
	vec.Erase(vec.Size() - 1)
	fmt.Println(&vec)
	vec.Erase(3)
	fmt.Println(&vec)
	vec.PopBack()
	fmt.Println(&vec)

	// avltree
	fmt.Println("\navltree test >>>>>>>>>>>>>>>>>>>>>")
	tr := avltree.NewAVLTree[int](Less[int], Equal[int])
	tr.Init([]int{1, 7, 0, 9, 23, -1, 45, -9, 21}, Greater[int], Equal[int])
	fmt.Println(tr)

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
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	tr.Erase(1)
	for it := tr.End(); it != nil; it = it.GetPre() {
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	fmt.Println(tr.Size())

	// sort test
	fmt.Println("avltree sort test >>>>>>>>>>>>>>>>>>>>>")
	vec.Clear()
	tr.Clear()

	// 生成数据
	fmt.Println("generate random data...")
	rand.Seed(time.Now().Unix())
	trNum := 100000
	randSet := hashset.NewHashSetCapacity[int](trNum)
	for i := 0; i < trNum; i += 1 {
		for true {
			num := rand.Intn(trNum)
			if !randSet.Contains(num) {
				randSet.Insert(num)
				vec.PushBack(num)
				break
			}

		}
	}
	// 缓存红黑树测速随机数据
	vec_rb := vector.NewVector[int]()
	vec_rb.Init(vec)
	fmt.Println("run...")

	// tr
	start := time.Now()
	tr.Init(vec, Less[int], Equal[int])
	cost := time.Since(start)
	fmt.Printf("avl cost=[%s]\n", cost)
	fmt.Println("avl ok...", tr.ForSize(), tr.Size())

	// quick sort
	start = time.Now()
	vec.Sort(Less[int])
	cost = time.Since(start)
	fmt.Printf("sort cost=[%s]\n", cost)

	// result compare
	fmt.Println("result compare...")
	start = time.Now()
	isOk := true
	trit := tr.Begin()
	for i := 0; i < trNum; i += 1 {
		vecValue := vec[i]
		trValue := tr.At(i).Value
		if trValue != trit.Value {
			fmt.Println("it error >>> ", i, trit.Value, trValue)
		}
		if vecValue != trValue {
			fmt.Println("tr error >>> ", i, vecValue, trValue, vecValue == trValue)
			isOk = false
			break
		}
		trit = trit.GetNext()
	}
	cost = time.Since(start)
	fmt.Printf("vec tr compare cost=[%s]\n", cost)
	if isOk {
		fmt.Println("all ok...")
	} else {
		fmt.Println("error!!!")
	}

	// avltreemap
	fmt.Println("\navltreemap test >>>>>>>>>>>>>>>>>>>>>")
	avlmp := avltree.AVLTreeMap[string, int]{}
	avlmp.Init([]Pair[string, int]{}, Less[string], Equal[string])
	avlmp.Insert("hello", 212)
	avlmp.Insert("hellol", -1)
	avlmp.Insert("ello ", 5454)
	avlmp.Insert("world", 3)
	avlmp.Insert("ajgshja", 3)
	fmt.Println(avlmp.Find("world"))
	fmt.Println(avlmp.Find("hello "))
	fmt.Println(&avlmp)

	// rbtree
	fmt.Println("\nrbtree test >>>>>>>>>>>>>>>>>>>>>")
	rbtr := rbtree.NewRBTree[int](Less[int], Equal[int])
	rbtr.Init([]int{1, 7, 0, 9, 23, -1, 45, -9, 21}, Greater[int], Equal[int])
	fmt.Println(rbtr.Size())
	rbtr.ForEach(func(value *int) {
		fmt.Println(*value)
		// time.Sleep(time.Duration(2) * time.Second)
	})
	fmt.Println(&rbtr)

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

	// sort test
	fmt.Println("rbtree sort test >>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("run...")
	rbtr.Clear()

	// tr
	start = time.Now()
	rbtr.Init(vec_rb, Less[int], Equal[int])
	cost = time.Since(start)
	fmt.Printf("rb cost=[%s]\n", cost)
	fmt.Println("ok...", rbtr.ForSize(), rbtr.Size())
	vec_rb_i := 0
	isOk = true
	rbtr.ForEach(func(value *int) {
		if *value != vec[vec_rb_i] {
			isOk = false
		}
		vec_rb_i += 1
	})
	fmt.Println("isOk:", isOk)

	// multiavltree
	fmt.Println("\nmultiavltree test >>>>>>>>>>>>>>>>>>>>>")
	mtr := avltree.NewMultiAVLTree[int](Less[int], Equal[int])
	mtr.Init([]int{1, 7, 0, 9, 23, -1, 45, -9, 21, 0, 1, 0, 9, 21, 0}, Greater[int], Equal[int])
	fmt.Println(mtr)
	fmt.Println(1, ":", mtr.Order(1))
	fmt.Println(7, ":", mtr.Order(7))
	fmt.Println(0, ":", mtr.Order(0))
	fmt.Println(9, ":", mtr.Order(9))
	fmt.Println(23, ":", mtr.Order(23))
	fmt.Println(-1, ":", mtr.Order(-1))
	fmt.Println(45, ":", mtr.Order(45))
	fmt.Println(-9, ":", mtr.Order(-9))
	fmt.Println(21, ":", mtr.Order(21))
	fmt.Println(-2, ":", mtr.Order(-2))
	fmt.Println(49, ":", mtr.Order(49))
	fmt.Println(-30, ":", mtr.Order(-30))

	fmt.Println(0, ":", mtr.At(0))
	fmt.Println(1, ":", mtr.At(1))
	fmt.Println(2, ":", mtr.At(2))
	fmt.Println(3, ":", mtr.At(3))
	fmt.Println(4, ":", mtr.At(4))
	fmt.Println(4, ":", mtr.At(5))
	fmt.Println(6, ":", mtr.At(6))
	fmt.Println(7, ":", mtr.At(7))
	fmt.Println(8, ":", mtr.At(8))
	fmt.Println(9, ":", mtr.At(9))
	fmt.Println(10, ":", mtr.At(10))
	fmt.Println(11, ":", mtr.At(11))
	fmt.Println(12, ":", mtr.At(12))
	fmt.Println(13, ":", mtr.At(13))
	fmt.Println(14, ":", mtr.At(14))
	fmt.Println(15, ":", mtr.At(15))
	fmt.Println(-1, ":", mtr.At(-1))

	fmt.Println("lowerbound 0:", mtr.LowerBound(0))
	fmt.Println("lowerbound 44:", mtr.LowerBound(44))
	fmt.Println("lowerbound 12:", mtr.LowerBound(12))
	fmt.Println("lowerbound -5:", mtr.LowerBound(-5))
	fmt.Println("lowerbound -100:", mtr.LowerBound(-100))
	fmt.Println("lowerbound 27:", mtr.LowerBound(27))
	fmt.Println("lowerbound 6:", mtr.LowerBound(6))
	fmt.Println("lowerbound 9:", mtr.LowerBound(9))
	fmt.Println("lowerbound 22:", mtr.LowerBound(22))
	fmt.Println("lowerbound 333:", mtr.LowerBound(333))
	fmt.Println("lowerbound 46:", mtr.LowerBound(46))
	fmt.Println("lowerbound 3:", mtr.LowerBound(3))

	fmt.Println("upperbound 0:", mtr.UpperBound(0))
	fmt.Println("upperbound 44:", mtr.UpperBound(44))
	fmt.Println("upperbound 12:", mtr.UpperBound(12))
	fmt.Println("upperbound -5:", mtr.UpperBound(-5))
	fmt.Println("upperbound -100:", mtr.UpperBound(-100))
	fmt.Println("upperbound 27:", mtr.UpperBound(27))
	fmt.Println("upperbound 6:", mtr.UpperBound(6))
	fmt.Println("upperbound 9:", mtr.UpperBound(9))
	fmt.Println("upperbound 22:", mtr.UpperBound(22))
	fmt.Println("upperbound 333:", mtr.UpperBound(333))
	fmt.Println("upperbound 46:", mtr.UpperBound(46))
	fmt.Println("upperbound 3:", mtr.UpperBound(3))

	mnode, mok := mtr.Find(23)
	fmt.Println(mtr.Size(), mtr.ForSize(), mok, mnode)
	for it := mtr.Begin(); it != nil; it = it.GetNext() {
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	findNode, _ := mtr.Find(1)
	mtr.Erase(findNode)
	for it := mtr.End(); it != nil; it = it.GetPre() {
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	fmt.Println(mtr.Size(), mtr.ForSize())
	mtr.EraseAll(0)
	for it := mtr.End(); it != nil; it = it.GetPre() {
		fmt.Print(it.Value, " ")
	}
	fmt.Println()
	fmt.Println(mtr.Size(), mtr.ForSize())

	// sort test
	fmt.Println("multiavltree sort test >>>>>>>>>>>>>>>>>>>>>")
	vec.Clear()
	mtr.Clear()

	// 生成数据
	fmt.Println("generate random data...")
	rand.Seed(time.Now().Unix())
	mtrNum := 1000000
	for i := 0; i < mtrNum; i += 1 {
		vec.PushBack(rand.Intn(mtrNum / 10000))
	}

	// mtr
	fmt.Println("run...")
	start = time.Now()
	mtr.Init(vec, Less[int], Equal[int])
	cost = time.Since(start)
	fmt.Printf("multiavl cost=[%s]\n", cost)
	fmt.Println("multiavl ok...", mtr.ForSize(), mtr.Size())

	// quick sort
	start = time.Now()
	vec.Sort(Less[int])
	cost = time.Since(start)
	fmt.Printf("sort cost=[%s]\n", cost)

	// result compare
	fmt.Println("result compare...")
	start = time.Now()
	isOk = true
	mtrit := mtr.End()
	for i := mtrNum - 1; i >= 0; i -= 1 {
		vecValue := vec[i]
		// mtrValue := mtrit.Value
		mtrValue := mtr.At(i).Value
		if mtrValue != mtrit.Value {
			fmt.Println("mit error >>> ", i, mtrit.Value, mtrValue)
		}
		if vecValue != mtrValue {
			fmt.Println("error >>> ", i, vecValue, mtrValue, vecValue == mtrValue)
			isOk = false
			break
		}
		mtrit = mtrit.GetPre()
	}
	cost = time.Since(start)
	fmt.Printf("vec tr compare cost=[%s]\n", cost)
	if isOk {
		fmt.Println("all ok...")
	} else {
		fmt.Println("error!!!")
	}

	// multiavltreemap
	fmt.Println("\nmultiavltreemap test >>>>>>>>>>>>>>>>>>>>>")
	mavlmp := avltree.MultiAVLTreeMap[string, int]{}
	mavlmp.Init([]Pair[string, int]{}, Less[string], Equal[string])
	mavlmp.Insert("hello", 212)
	mavlmp.Insert("hello", 8473684)
	mavlmp.Insert("hello", 44532)
	mavlmp.Insert("hello", 98654)
	mavlmp.Insert("hellol", -1)
	mavlmp.Insert("ello ", 5454)
	mavlmp.Insert("world", 3)
	mavlmp.Insert("ajgshja", 3)
	fmt.Println(mavlmp.Find("world"))
	fmt.Println(mavlmp.Find("hello"))
	fmt.Println(mavlmp.Find("hello "))
	fmt.Println(&mavlmp)

	fmt.Println(0, ":", mavlmp.At(0))
	fmt.Println(1, ":", mavlmp.At(1))
	fmt.Println(2, ":", mavlmp.At(2))
	fmt.Println(3, ":", mavlmp.At(3))
	fmt.Println(4, ":", mavlmp.At(4))
	fmt.Println(4, ":", mavlmp.At(5))
	fmt.Println(6, ":", mavlmp.At(6))
	fmt.Println(7, ":", mavlmp.At(7))
	fmt.Println(8, ":", mavlmp.At(8))
	fmt.Println(-1, ":", mavlmp.At(-1))

	iter, _ := mavlmp.Find("hello")
	iterList := vector.Vector[*avltree.MultiAVLTreeNode[Pair[string, int]]]{}
	for i := mavlmp.Count("hello"); i > 0; i -= 1 {
		iterList.PushBack(iter)
		iter = iter.GetNext()
	}
	fmt.Println(iterList.Size())
	mavlmp.Erase(iterList[1])
	fmt.Println(&mavlmp)
	mavlmp.Erase(iterList[2])
	fmt.Println(&mavlmp)
	mavlmp.Erase(iterList[3])
	fmt.Println(&mavlmp)

	// common
	fmt.Println("\ncommon test >>>>>>>>>>>>>>>>>>>>>")
	var arr_test vector.Vector[int] = []int{2781, 3, 7, 0, 2, -1, 6, -2187, 21, 34, 672, 23}
	fmt.Println(*MaxElement(arr_test.GetIterator()), GetIndexByArrayPointer(arr_test, MaxElement(arr_test.GetIterator())))
	fmt.Println(*MinElement(arr_test.GetIterator()), GetIndexByArrayPointer(arr_test, MinElement(arr_test.GetIterator())))
	fmt.Println(*MaxElementByFunction(arr_test.GetIterator(), Greater[int]))
	fmt.Println(*MinElementByFunction(arr_test.GetIterator(), Less[int]))
	fmt.Println(Accumulate(arr_test.GetIterator(), 0))

	// // rbtree_at
	// rbtr_at := rbtree_at.NewRBTree(greater, equal)
	// rbtr_at.Init([]interface{}{1, 7, 0, 9, 23, -1, 45, -9, 21}, less, equal)
	// fmt.Println(rbtr_at.Size(), rbtr_at.ForSize())
	// rbtr_at.ForEach(func(value *interface{}) {
	// 	fmt.Println(*value)
	// })
	// fmt.Println(&rbtr_at)
	// fmt.Println(1, ":", rbtr_at.Order(1))
	// fmt.Println(7, ":", rbtr_at.Order(7))
	// fmt.Println(0, ":", rbtr_at.Order(0))
	// fmt.Println(9, ":", rbtr_at.Order(9))
	// fmt.Println(23, ":", rbtr_at.Order(23))
	// fmt.Println(-1, ":", rbtr_at.Order(-1))
	// fmt.Println(45, ":", rbtr_at.Order(45))
	// fmt.Println(-9, ":", rbtr_at.Order(-9))
	// fmt.Println(21, ":", rbtr_at.Order(21))
	// fmt.Println(-2, ":", rbtr_at.Order(-2))
	// fmt.Println(49, ":", rbtr_at.Order(49))
	// fmt.Println(-30, ":", rbtr_at.Order(-30))

	// fmt.Println(0, ":", rbtr_at.At(0))
	// fmt.Println(1, ":", rbtr_at.At(1))
	// fmt.Println(2, ":", rbtr_at.At(2))
	// fmt.Println(3, ":", rbtr_at.At(3))
	// fmt.Println(4, ":", rbtr_at.At(4))
	// fmt.Println(5, ":", rbtr_at.At(5))
	// fmt.Println(6, ":", rbtr_at.At(6))
	// fmt.Println(7, ":", rbtr_at.At(7))
	// fmt.Println(8, ":", rbtr_at.At(8))
	// fmt.Println(9, ":", rbtr_at.At(9))
	// fmt.Println(10, ":", rbtr_at.At(10))
	// fmt.Println(-1, ":", rbtr_at.At(-1))

	// fmt.Println("run...")
	// start = time.Now()
	// rbtr_at.Init(vec_rb, less, equal)
	// cost = time.Since(start)
	// fmt.Printf("rb cost=[%s]\n", cost)
	// fmt.Println("ok...", rbtr_at.ForSize(), rbtr_at.Size())
	// vec_rb_i = 0
	// isOk = true
	// rbtr_at.ForEach(func(value *interface{}) {
	// 	if *value != vec[vec_rb_i] {
	// 		isOk = false
	// 	}
	// 	vec_rb_i += 1
	// })
	// fmt.Println("isOk:", isOk)
}
