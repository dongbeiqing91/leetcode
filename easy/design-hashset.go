package easy

import "container/list"

const base = 769

type MyHashSet struct {
    data []list.List
}


func Constructor() MyHashSet {
    return MyHashSet{
		data: make([]list.List, base),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Add(key int)  {
    if this.Contains(key) {
		return
	}
	h := this.hash(key)
	this.data[h].PushBack(key)
}


func (this *MyHashSet) Remove(key int)  {
    h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[h].Remove(e)
		}
	}
}


func (this *MyHashSet) Contains(key int) bool {
    h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */