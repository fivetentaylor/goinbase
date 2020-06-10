package market

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/trees/btree"
	"os"
)

func NewMarket() Market {
	return Market{}
}

func HelloWorld() {
	tree := btree.NewWithStringComparator(16)
	tree.Put("hello", "world")
	tree.Put("goodbye", "friends")
	tree.Put("yes", "no")
	tree.Put("bye", "hello")
	tree.Put("bye", 12)

	json_str, err := tree.ToJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json_str))

	json_str, err = json.Marshal(Order{ClientOid: "123"})
	os.Stdout.Write(json_str)
	fmt.Println()

	fmt.Println(Order{Type: "Ask"}.Type)
}
