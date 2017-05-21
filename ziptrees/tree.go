package ziptrees

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
	Rank  int32
}

func rank(n *Node) int32 {
	if n == nil {
		return 0
	}
	return n.Rank
}

func insert(n *Node, x *Node) *Node {
	if n == nil {
		return x
	}
	if x.Key < n.Key {
		if insert(n.Left, x) == x {
			if x.Rank < n.Rank {
				n.Left = x
			} else {
				n.Left = x.Right
				x.Right = n
				return x
			}
		}
	} else {
		if insert(n.Right, x) == x {
			if x.Rank <= n.Rank {
				n.Right = x
			} else {
				n.Right = x.Left
				x.Left = n
				return x
			}
		}
	}
	return n
}

type ZipTree struct {
	Root *Node
}

func popcount(x int32) int32 {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0F0F0F0F
	x = x + (x >> 8)
	x = x + (x >> 16)
	return x & 0x0000003F
}

func newRank() int32 {
	t := rand.Int31()
	rank := popcount((t & -t) - 1)
	return rank
}

func (zt *ZipTree) Insert(k int, v string) {
	x := &Node{k, v, nil, nil, newRank()}
	zt.Root = insert(zt.Root, x)
}

func New() *ZipTree {
	return &ZipTree{nil}
}

func Traverse(n *Node) {
	if n == nil {
		return
	}
	Traverse(n.Left)
	fmt.Println(n.Key, n.Value)
	Traverse(n.Right)
}

func (t *ZipTree) Traverse() {
	Traverse(t.Root)
}
