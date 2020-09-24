package main

func main() {

}

type MyHashSet struct {
	l []bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		l: make([]bool, 1000001),
	}
}

func (this *MyHashSet) Add(key int) {
	this.l[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.l[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.l[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
