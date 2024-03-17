/*
 * @Description: hashmap_test
 * @Version:
 * @Autor: dreamy-xay
 * @Date: 2021-10-12 17:01:31
 * @LastEditors: dreamy-xay
 * @LastEditTime: 2021-10-14 12:46:22
 */
package hashmap

import "testing"

func TestInit(t *testing.T) {
	mp := NewHashMap()
	mp.Init([]Pair{
		{
			Key:   "hello",
			Value: 1,
		},
		{
			Key:   "world",
			Value: 3,
		},
	})
	if v, ok := mp["hello"]; v != 1 || !ok {
		t.Error("map[\"hello\"] is not 1 or nonexistent")
	}

	if v, ok := mp["world"]; v != 3 || !ok {
		t.Error("map[\"world\"] is not 3 or nonexistent")
	}
}

func TestSize(t *testing.T) {
	mp := NewHashMap()
	if mp.Size() != 0 {
		t.Error("hashmap is not of size 0")
	}
	mp.Init([]Pair{
		{
			Key:   "hello",
			Value: 1,
		},
		{
			Key:   "world",
			Value: 3,
		},
	})
	if mp.Size() != 2 {
		t.Error("hashmap is not of size 2")
	}
}

func TestEmpty(t *testing.T) {
	mp := NewHashMap()
	if !mp.Empty() {
		t.Error("hashmap is empty")
	}
	mp.Init([]Pair{
		{
			Key:   "hello",
			Value: 1,
		},
		{
			Key:   "world",
			Value: 3,
		},
	})
	if mp.Empty() {
		t.Error("hashmap is not empty")
	}

}

func TestClear(t *testing.T) {
	mp := NewHashMap()
	mp.Init([]Pair{
		{
			Key:   "hello",
			Value: 1,
		},
		{
			Key:   "world",
			Value: 3,
		},
	})
	if mp.Clear(); !mp.Empty() {
		t.Error("hashmap is empty")
	}
	if tmp := NewHashMap(); !tmp.Empty() {
		t.Error("hashmap is empty")
	}
}

func TestInsert(t *testing.T) {
	mp := NewHashMap()
	mp.Insert("hello", 7)
	mp.Insert("world", 5)
	if v, ok := mp["hello"]; v != 7 || !ok {
		t.Error("map[\"hello\"] is not 1 or nonexistent")
	}

	if v, ok := mp["world"]; v != 5 || !ok {
		t.Error("map[\"world\"] is not 3 or nonexistent")
	}
}

func TestErase(t *testing.T) {
	mp := NewHashMap()
	mp.Insert("hello", 7)
	mp.Insert("world", 5)
	mp.Erase("hello")
	if _, ok := mp["hello"]; ok {
		t.Error("map[\"hello\"] is nonexistent")
	}
	mp.Erase("world")
	if _, ok := mp["world"]; ok {
		t.Error("map[\"world\"] is nonexistent")
	}
}

func TestFind(t *testing.T) {
	mp := NewHashMap()
	mp.Insert("hello", 7)
	mp.Insert("world", 5)
	if pair, ok := mp.Find("hello"); pair.Key != "hello" || pair.Value != 7 || !ok {
		t.Errorf("mp.Find(\"hello\") error: {%s: %d}", pair.Key, pair.Value)
	}
	if pair, ok := mp.Find("world"); pair.Key != "world" || pair.Value != 5 || !ok {
		t.Errorf("mp.Find(\"world\") error: {%s: %d}", pair.Key, pair.Value)
	}
	if _, ok := mp.Find("wrld"); ok {
		t.Errorf("map[\"wrld\"] is nonexistent")
	}
}

func TestContains(t *testing.T) {
	mp := NewHashMap()
	mp.Insert("hello", 7)
	mp.Insert("world", 5)
	if !mp.Contains("hello") {
		t.Error("map[\"hello\"] is existent")
	}
	if !mp.Contains("hello") {
		t.Error("map[\"world\"] is existent")
	}
	if mp.Contains("wrld") {
		t.Error("map[\"wrld\"] is nonexistent")
	}
}

func TestKeys(t *testing.T) {
	mp := NewHashMap()
	mp.Insert("hello", 7)
	mp.Insert("world", 5)
	Keys := mp.Keys()
	if len(Keys) != 2 || !(Keys[0] == "hello" || Keys[0] == "world") || !(Keys[1] == "hello" || Keys[1] == "world") || Keys[0] == Keys[1] {
		t.Errorf("an error occurred while getting the Keys: %d  =>  %v", len(Keys), Keys)
	}
}

func TestForEach(t *testing.T) {
	mp := NewHashMap()
	mp.Insert("hello", 7)
	mp.Insert("world", 5)
	index := 0
	var Key0 string
	mp.ForEach(func(Key, Value *interface{}) {
		if index == 0 {
			if !((*Key == "hello" && *Value == 7) || (*Key == "world" && *Value == 5)) {
				t.Errorf("an error occurred while foreach 0: {%s: %d}", *Key, *Value)
			}
			Key0 = (*Key).(string)
		} else if index == 1 {
			if !((*Key == "hello" && *Value == 7) || (*Key == "world" && *Value == 5)) || *Key == Key0 {
				t.Errorf("an error occurred while foreach 1: {%s: %d}", *Key, *Value)
			}
		}
		index += 1
	})
	if index != 2 {
		t.Errorf("traversal count error: %d", index)
	}
}
