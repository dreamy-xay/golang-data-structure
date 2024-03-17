/*
 * @Description: hashmap
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 19:55:07
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2023-09-05 16:24:00
 */

package hashmap

import (
	. "collections/common"
	"fmt"
)

type HashMap[KEY comparable, VALUE any] map[KEY]VALUE

func NewHashMap[KEY comparable, VALUE any]() HashMap[KEY, VALUE] {
	return make(map[KEY]VALUE)
}

func NewHashMapCapacity[KEY comparable, VALUE any](capacity int) HashMap[KEY, VALUE] {
	return make(map[KEY]VALUE, capacity)
}

func (hashmap *HashMap[KEY, VALUE]) Init(arr []Pair[KEY, VALUE]) {
	hashmap.Clear()
	for _, elem := range arr {
		(*hashmap)[elem.Key] = elem.Value
	}
}

func (hashmap *HashMap[KEY, VALUE]) Empty() bool {
	return len(*hashmap) == 0
}

func (hashmap *HashMap[KEY, VALUE]) Size() int {
	return len(*hashmap)
}

func (hashmap *HashMap[KEY, VALUE]) Clear() {
	for k := range *hashmap {
		delete(*hashmap, k)
	}
}

func (hashmap *HashMap[KEY, VALUE]) Insert(key KEY, value VALUE) {
	(*hashmap)[key] = value
}

func (hashmap *HashMap[KEY, VALUE]) InsertPair(pair Pair[KEY, VALUE]) {
	(*hashmap)[pair.Key] = pair.Value
}

func (hashmap *HashMap[KEY, VALUE]) Erase(key KEY) {
	delete(*hashmap, key)
}

func (hashmap *HashMap[KEY, VALUE]) Find(key KEY) (*Pair[KEY, VALUE], bool) {
	value, ok := (*hashmap)[key]
	return &Pair[KEY, VALUE]{Key: key, Value: value}, ok
}

func (hashmap *HashMap[KEY, VALUE]) Contains(key KEY) bool {
	_, ok := (*hashmap)[key]
	return ok
}

func (hashmap *HashMap[KEY, VALUE]) Keys() []KEY {
	keys := make([]KEY, 0, len(*hashmap))
	for k := range *hashmap {
		keys = append(keys, k)
	}
	return keys
}

func (hashmap *HashMap[KEY, VALUE]) ForEach(fn func(key *KEY, value *VALUE)) {
	for key, value := range *hashmap {
		fn(&key, &value)
	}
}

func (hashmap *HashMap[KEY, VALUE]) ForEachPair(fn func(pair *Pair[KEY, VALUE])) {
	for key, value := range *hashmap {
		fn(&Pair[KEY, VALUE]{Key: key, Value: value})
	}
}

func (hashmap *HashMap[KEY, VALUE]) GetIterator() *Iterator[Pair[KEY, VALUE]] {
	pairList := []Pair[KEY, VALUE]{}
	for key, value := range *hashmap {
		pairList = append(pairList, Pair[KEY, VALUE]{Key: key, Value: value})
	}
	current := -1
	listLen := len(pairList)
	return NewIterator[Pair[KEY, VALUE]](func() bool {
		return current+1 < listLen
	}, func() *Pair[KEY, VALUE] {
		current += 1
		if current < listLen {
			return &pairList[current]
		} else {
			return nil
		}

	})
}

func (hashmap *HashMap[KEY, VALUE]) String() string {
	str := "HashMap {\n"
	for key, value := range *hashmap {
		str += "    " + fmt.Sprint(key) + ": " + fmt.Sprint(value) + ",\n"
	}
	str += "}"
	return str
}
