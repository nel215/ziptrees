package ziptrees

import (
	"testing"
)

func traverse(n *Node, res []int) {
	if n == nil {
		return
	}
	traverse(n.Left, res)
	res = append(res, n.Key)
	traverse(n.Right, res)
}

func (t *ZipTree) traverse() []int {
	res := make([]int, 0)
	traverse(t.Root, res)
	return res
}

func TestInsert(t *testing.T) {
	tree := New()
	for i := 0; i < 100; i++ {
		tree.Insert(i, "value")
	}
	res := tree.traverse()
	for i, k := range res {
		if i != k {
			t.Errorf("expected key is %d, but got %d", i, k)
		}
	}
}

func TestDelete(t *testing.T) {
	tree := New()
	for i := 0; i < 100; i++ {
		tree.Insert(i, "value")
		tree.Delete(i)
	}
	for i := 0; i < 100; i++ {
		tree.Insert(i, "value")
	}
	for i := 0; i < 100; i++ {
		tree.Delete(i)
	}
	res := tree.traverse()
	if len(res) != 0 {
		t.Errorf("expected length of keys is 0, but got %d", len(res))
	}
}
