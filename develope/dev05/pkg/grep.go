package pkg

import (
	"fmt"
	"sort"
)

type Node struct {
	Key   int
	Value string
}
type Grep struct {
	Result []Node
}

func NewGrep() *Grep {
	return &Grep{
		Result: []Node{},
	}
}
func (g *Grep) SortResultASC() {
	sort.Slice(g.Result, func(i, j int) bool {
		return g.Result[i].Key < g.Result[j].Key
	})
}
func (g *Grep) Print(indexing bool) {
	g.SortResultASC()
	switch indexing {
	case true:
		for _, v := range g.Result {
			fmt.Printf("%d. %s\n", v.Key, v.Value)
		}
	default:
		for _, v := range g.Result {
			fmt.Printf("%s\n", v.Value)
		}
	}
}
