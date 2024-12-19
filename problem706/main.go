package main

import "fmt"

type MyHashMap struct {
	data map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: make(map[int]int),
	}
}

func (hash *MyHashMap) Put(key int, value int) {
	hash.data[key] = value
	fmt.Printf("Put %v for %v\n", key, hash.data)
}

func (hash *MyHashMap) Get(key int) int {
	if value, exists := hash.data[key]; exists {
		return value
	}
	return -1
}

func (hash *MyHashMap) Remove(key int) {
	delete(hash.data, key)
	fmt.Printf("Removed %v from %v\n", key, hash.data)
}

func main() {

	obj := Constructor()
	obj.Put(0, 1)
	obj.Put(1, 2)
	out := obj.Get(0)
	fmt.Printf("Get %v key: %v\n", 0, out)
	obj.Remove(0)

}
