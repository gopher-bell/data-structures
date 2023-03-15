package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func hash(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}

	return sum % ArraySize
}

func (h *HashTable) Insert(k string) {
	index := hash(k)
	h.array[index].insert(k)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

func (h *HashTable) Search(k string) bool {
	index := hash(k)
	return h.array[index].search(k)
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (h *HashTable) Delete(k string) {
	index := hash(k)
	h.array[index].delete(k)
}

func (b *bucket) delete(k string) {
	prevNode := b.head
	if prevNode == nil {
		fmt.Println("does not exist")
		return
	}
	if prevNode.key == k {
		b.head = prevNode.next
		return
	}

	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
			return
		}

		prevNode = prevNode.next
	}
}

func main() {
	testHashTable := Init()
	testHashTable.Insert("kim")
	testHashTable.Insert("Lee")
	testHashTable.Insert("Park")
	testHashTable.Insert("Lee")
	testHashTable.Delete("z")
	testHashTable.Delete("Lee")
	testHashTable.Insert("Lee")
	fmt.Println(testHashTable.Search("kim"), testHashTable.Search("bae"))
}
