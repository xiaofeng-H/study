package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestLRU(t *testing.T) {
	dl := algorithms.NewDoubleList()
	//fmt.Println(dl.Root)
	node := new(algorithms.Node)
	node2 := algorithms.NewNode(1, 1)
	fmt.Println(node)
	fmt.Println(node2)
	//fmt.Printf("%v,%T\n", node, node)
	dl.AddFirst(node)
	dl.AddFirst(node2)
	node3 := dl.RemoveLast()
	fmt.Printf("node1=%v,node2=%v,node3=%v\n", node, node2, node3)
	fmt.Println(node3.GetKey())
}

func TestLRU2(t *testing.T) {
	lru := algorithms.NewLRUCache(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
}
