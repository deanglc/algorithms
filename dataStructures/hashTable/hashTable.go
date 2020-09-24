package hashTable

import (
	"fmt"
	"sync"
)

/*
golang.map相关文件 src/runtime/
map.go
map_fast32.go
map_fast64.go
map_faststr.go
*/

/* refer :https://baike.baidu.com/item/%E5%93%88%E5%B8%8C%E8%A1%A8/5981869
散列表（Hash table，也叫哈希表），是根据关键码值(Key value)而直接进行访问的数据结构。
也就是说，它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。
这个映射函数叫做散列函数，存放记录的数组叫做散列表。

给定表M，存在函数f(key)，对任意给定的关键字值key，
代入函数后若能得到包含该关键字的记录在表中的地址，则称表M为哈希(Hash）表，函数f(key)为哈希(Hash) 函数。
*/

// 利用golang.map 实现hashTable的几个方法
type Key interface{}
type Value interface{}

type ValueHashtable struct {
	items map[int]Value
	lock  sync.RWMutex
}

func Hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *ValueHashtable) Put(k Key, v Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := Hash(k)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[i] = v
}

func (ht *ValueHashtable) Remove(k Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := Hash(k)
	delete(ht.items, i)
}

func (ht *ValueHashtable) Get(k Key) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	i := Hash(k)
	return ht.items[i]
}

func (ht *ValueHashtable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)

}
