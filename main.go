package main

import (
	"github.com/nel215/ziptrees"
)

func main() {
	tree := ziptrees.New()
	for i := 0; i < 10; i++ {
		tree.Insert(i, "value")
	}
	tree.Traverse()
}
