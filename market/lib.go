package market

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"os"
)

type OrderedTree struct {
	*treemap.Map
}

func NewMarket() Market {
	return Market{}
}

func (tree *OrderedTree) MarshalJSON() ([]byte, error) {
	return tree.ToJSON()
}

func (tree *OrderedTree) UnmarshalJSON(payload []byte) error {
	return tree.FromJSON(payload)
}

type ThingToSerialize struct {
	Something string
	Tree      *OrderedTree
}

func HelloWorld() {
	tree := OrderedTree{treemap.NewWithStringComparator()}
	tree.Put("hello", "world")
	tree.Put("goodbye", "friends")
	tree.Put("yes", "no")
	tree.Put("bye", "hello")
	tree.Put("bye", 123)

	t := ThingToSerialize{"hello world", &tree}

	json_str, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json_str))

	nt := ThingToSerialize{"", &OrderedTree{treemap.NewWithStringComparator()}}
	json.Unmarshal(json_str, &nt)

	fmt.Println(nt.Tree)
	// fmt.Println(nt.Something)

	json_str, err = json.Marshal(Order{ClientOid: "123"})
	os.Stdout.Write(json_str)
	fmt.Println()

	fmt.Println(Order{Type: "Ask"}.Type)
}
