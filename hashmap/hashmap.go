/*
 * @Description: hashmap
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 19:55:07
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-14 12:45:48
 */

package hashmap

import "fmt"

type Pair struct {
	Key   interface{}
	Value interface{}
}

type HashMap map[interface{}]interface{}

func NewHashMap() HashMap {
	return make(map[interface{}]interface{})
}

func NewHashMapCapacity(capacity int) HashMap {
	return make(map[interface{}]interface{}, capacity)
}

func (hashmap *HashMap) Init(arr []Pair) {
	hashmap.Clear()
	for _, elem := range arr {
		(*hashmap)[elem.Key] = elem.Value
	}
}

func (hashmap *HashMap) Empty() bool {
	return len(*hashmap) == 0
}

func (hashmap *HashMap) Size() int {
	return len(*hashmap)
}

func (hashmap *HashMap) Clear() {
	for k := range *hashmap {
		delete(*hashmap, k)
	}
}

func (hashmap *HashMap) Insert(key, value interface{}) {
	(*hashmap)[key] = value
}

func (hashmap *HashMap) InsertPair(pair Pair) {
	(*hashmap)[pair.Key] = pair.Value
}

func (hashmap *HashMap) Erase(key interface{}) {
	delete(*hashmap, key)
}

func (hashmap *HashMap) Find(key interface{}) (Pair, bool) {
	value, ok := (*hashmap)[key]
	return Pair{key, value}, ok
}

func (hashmap *HashMap) Contains(key interface{}) bool {
	_, ok := (*hashmap)[key]
	return ok
}

func (hashmap *HashMap) Keys() []interface{} {
	keys := make([]interface{}, 0, len(*hashmap))
	for k := range *hashmap {
		keys = append(keys, k)
	}
	return keys
}

func (hashmap *HashMap) ForEach(fn func(key, value *interface{})) {
	for key, value := range *hashmap {
		fn(&key, &value)
	}
}

func (hashmap *HashMap) ForEachPair(fn func(pair *Pair)) {
	for key, value := range *hashmap {
		fn(&Pair{key, value})
	}
}

func (hashmap *HashMap) String() string {
	str := "HashMap {\n"
	for key, value := range *hashmap {
		str += "    " + fmt.Sprint(key) + ": " + fmt.Sprint(value) + ",\n"
	}
	str += "}"
	return str
}
