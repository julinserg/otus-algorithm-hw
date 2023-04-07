package p10bst2

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

type NodeRandomBST struct {
	key        int
	value      string
	size       int
	childLeft  *NodeRandomBST
	childRight *NodeRandomBST
}

// https://habr.com/ru/post/145388/
type RandomBST struct {
	root *NodeRandomBST
	size int
}

func (s *RandomBST) getSize(p *NodeRandomBST) int {
	if p == nil {
		return 0
	} else {
		return p.size
	}
}

func (s *RandomBST) fixSize(p *NodeRandomBST) {
	p.size = s.getSize(p.childLeft) + s.getSize(p.childRight) + 1
}

func (s *RandomBST) rotateRight(p *NodeRandomBST) *NodeRandomBST {
	q := p.childLeft
	if q == nil {
		return p
	}
	p.childLeft = q.childRight
	q.childRight = p
	q.size = p.size
	s.fixSize(p)
	if p == s.root {
		s.root = q
	}
	return q
}

func (s *RandomBST) rotateLeft(q *NodeRandomBST) *NodeRandomBST {
	p := q.childRight
	if p == nil {
		return q
	}
	q.childRight = p.childLeft
	p.childLeft = q
	p.size = q.size
	s.fixSize(q)
	if q == s.root {
		s.root = p
	}
	return p
}

func (s *RandomBST) insertRoot(node *NodeRandomBST, key int, value string) *NodeRandomBST {
	if node == nil {
		s.size++
		return &NodeRandomBST{key, value, 1, nil, nil}
	}
	if key < node.key {
		node.childLeft = s.insertRoot(node.childLeft, key, value)
		return s.rotateRight(node)
	} else if key > node.key {
		node.childRight = s.insertRoot(node.childRight, key, value)
		return s.rotateLeft(node)
	} else {
		node.value = value
		return node
	}
}

func (s *RandomBST) Size() int {
	return s.size
}

func (s *RandomBST) IsEmpty() bool {
	return s.Size() == 0
}

func (s *RandomBST) searchNodeAndInsert(node *NodeRandomBST, key int, value string) *NodeRandomBST {
	if node == nil {
		s.size++
		return &NodeRandomBST{key, value, 1, nil, nil}
	}

	if rand.Intn(math.MaxInt32)%(node.size+1) == 0 {
		return s.insertRoot(node, key, value)
	}
	if key < node.key {
		node.childLeft = s.searchNodeAndInsert(node.childLeft, key, value)
	} else if key > node.key {
		node.childRight = s.searchNodeAndInsert(node.childRight, key, value)
	} else {
		node.value = value
	}
	s.fixSize(node)
	return node
}

func (s *RandomBST) Insert(key int, value string) {
	if s.root == nil {
		s.root = &NodeRandomBST{key, value, 1, nil, nil}
		s.size++
	} else {
		node := s.searchNodeAndInsert(s.root, key, value)
		if node == nil {
			panic("ERROR in searchNodeAndInsert")
		}
	}
}

func (s *RandomBST) join(p *NodeRandomBST, q *NodeRandomBST) *NodeRandomBST {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	if rand.Intn(math.MaxInt32)%(p.size+q.size) < p.size {
		p.childRight = s.join(p.childRight, q)
		s.fixSize(p)
		return p
	} else {
		q.childLeft = s.join(p, q.childLeft)
		s.fixSize(q)
		return q
	}
}

func (s *RandomBST) removeNode(node *NodeRandomBST, key int) *NodeRandomBST {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.childLeft = s.removeNode(node.childLeft, key)
	} else if key > node.key {
		node.childRight = s.removeNode(node.childRight, key)
	} else {
		q := s.join(node.childLeft, node.childRight)
		if node == s.root {
			s.root = q
		}
		s.size--
		return q
	}
	return node
}

func (s *RandomBST) Remove(key int) {
	s.removeNode(s.root, key)
}

func (s *RandomBST) searchNode(node *NodeRandomBST, key int) string {
	if node == nil {
		return ""
	}
	if key == node.key {
		return node.value
	} else if key > node.key {
		return s.searchNode(node.childRight, key)
	} else {
		return s.searchNode(node.childLeft, key)
	}
}

func (s *RandomBST) Search(key int) string {
	return s.searchNode(s.root, key)
}

func (s *RandomBST) printNode(node *NodeRandomBST, level int) {
	if node == nil {
		return
	}
	level += 10
	s.printNode(node.childRight, level)
	fmt.Println("")
	for i := 10; i < level; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(node.key, "("+strconv.Itoa(node.size)+")")

	s.printNode(node.childLeft, level)

}

func (s *RandomBST) Print() {
	s.printNode(s.root, 0)
}

func (s *RandomBST) collectKey(node *NodeRandomBST, array *[]int) {
	if node == nil {
		return
	}
	*array = append(*array, node.key)
	s.collectKey(node.childLeft, array)
	s.collectKey(node.childRight, array)
}

func (s *RandomBST) ListKey() []int {
	result := make([]int, 0)
	s.collectKey(s.root, &result)
	return result
}
