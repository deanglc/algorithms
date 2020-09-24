package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	rst := obj.Get(1)
	rst2 := obj.Get(3)
	fmt.Println(rst, rst2)
	obj.Remove(1)
}

type MyHashMap struct {
	item []kv
}
type kv struct {
	key int
	val int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	flag := false
	for k := range this.item {
		if this.item[k].key == key {
			this.item[k].val = value
			flag = true
			break
		}
	}
	if !flag {
		this.item = append(this.item, kv{key, value})
	}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	for k := range this.item {
		if this.item[k].key == key {
			return this.item[k].val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for k := range this.item {
		if this.item[k].key == key {
			this.item = append(this.item[:k], this.item[k+1:]...)
			break
		}
	}
}
