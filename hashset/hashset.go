/*
 * @Description: hashset
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-10 21:51:20
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-12 15:33:36
 */
package hashset

import "fmt"

type HashSet struct {
	elements map[interface{}]bool
}

func NewHashSet() HashSet {
	return HashSet{elements: make(map[interface{}]bool)}
}

func NewHashSetCapacity(capacity int) *HashSet {
	return &HashSet{elements: make(map[interface{}]bool, capacity)}
}

func (hashset *HashSet) Init(arr []interface{}) {
	hashset.Clear()
	for _, elem := range arr {
		hashset.elements[elem] = true
	}
}

func (hashset *HashSet) Empty() bool {
	return len(hashset.elements) == 0
}

func (hashset *HashSet) Size() int {
	return len(hashset.elements)
}

func (hashset *HashSet) Clear() {
	for key := range hashset.elements {
		delete(hashset.elements, key)
	}
}

func (hashset *HashSet) Insert(key interface{}) {
	hashset.elements[key] = true
}

func (hashset *HashSet) Erase(key interface{}) {
	delete(hashset.elements, key)
}

func (hashset *HashSet) Find(key interface{}) (interface{}, bool) {
	_, ok := (hashset.elements)[key]
	return key, ok
}

func (hashset *HashSet) Contains(key interface{}) bool {
	_, ok := hashset.elements[key]
	return ok
}

func (hashset *HashSet) Keys() []interface{} {
	keys := make([]interface{}, 0, len(hashset.elements))
	for key := range hashset.elements {
		keys = append(keys, key)
	}
	return keys
}

func (hashset *HashSet) ForEach(fn func(key *interface{})) {
	for key, value := range hashset.elements {
		if value {
			fn(&key)
		}
	}
}

func (hashset *HashSet) String() string {
	str := "HashSet {\n"
	for key := range hashset.elements {
		str += "    " + fmt.Sprint(key) + ",\n"
	}
	str += "}"
	return str
}
