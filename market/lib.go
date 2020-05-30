package goinbase

import (
	"fmt"
	"github.com/emirpasic/gods/trees/btree"
)

func NewMarket() {
}

func HelloWorld() {
	tree := btree.NewWithStringComparator(16)
	tree.Put("hello", "world")
	tree.Put("goodbye", "friends")
	tree.Put("yes", "no")
	tree.Put("bye", "hello")
	tree.Put("bye", 12)

	json, err := tree.ToJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
