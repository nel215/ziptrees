package ziptrees

import (
	"math/rand"
)

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
	Rank  int32
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

func zip(x, y *Node) *Node {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	if x.Rank < y.Rank {
		y.Left = zip(x, y.Left)
		return y
	} else {
		x.Right = zip(x.Right, y)
		return x
	}
}

func delete_node(n *Node, k int) *Node {
	if k == n.Key {
		return zip(n.Left, n.Right)
	}
	if k < n.Key {
		if n.Left != nil && k == n.Left.Key {
			n.Left = zip(n.Left.Left, n.Left.Right)
		} else {
			delete_node(n.Left, k)
		}
	} else {
		if n.Right != nil && k == n.Right.Key {
			n.Right = zip(n.Right.Left, n.Right.Right)
		} else {
			delete_node(n.Right, k)
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
	return popcount((t & -t) - 1)
}

func (zt *ZipTree) Insert(k int, v string) {
	x := &Node{k, v, nil, nil, newRank()}
	zt.Root = insert(zt.Root, x)
}

func (zt *ZipTree) Delete(k int) {
	zt.Root = delete_node(zt.Root, k)
}

func New() *ZipTree {
	return &ZipTree{nil}
}
