package p10bst

import (
	"fmt"
	"strconv"
)

type NodeAVL struct {
	key        int
	value      string
	height     int
	parent     *NodeAVL
	childLeft  *NodeAVL
	childRight *NodeAVL
}

// https://habr.com/ru/post/150732/
type AVLBST struct {
	root *NodeAVL
	size int
}

func (s *AVLBST) getHeight(p *NodeAVL) int {
	if p == nil {
		return 0
	} else {
		return p.height
	}
}

func (s *AVLBST) getBFactor(p *NodeAVL) int {
	return s.getHeight(p.childRight) - s.getHeight(p.childLeft)
}

func (s *AVLBST) fixHeight(p *NodeAVL) {
	hl := s.getHeight(p.childLeft)
	hr := s.getHeight(p.childRight)
	if hl > hr {
		p.height = hl + 1
	} else {
		p.height = hr + 1
	}
}

func (s *AVLBST) rotateRight(p *NodeAVL) *NodeAVL {
	q := p.childLeft
	p.childLeft = q.childRight
	q.childRight = p
	s.fixHeight(p)
	s.fixHeight(q)
	if p == s.root {
		s.root = q
	}
	return q
}

func (s *AVLBST) rotateLeft(q *NodeAVL) *NodeAVL {
	p := q.childRight
	q.childRight = p.childLeft
	p.childLeft = q
	s.fixHeight(q)
	s.fixHeight(p)
	if q == s.root {
		s.root = p
	}
	return p
}

func (s *AVLBST) balance(p *NodeAVL) *NodeAVL {
	s.fixHeight(p)
	if s.getBFactor(p) == 2 {
		if s.getBFactor(p.childRight) < 0 {
			p.childRight = s.rotateRight(p.childRight)
		}
		return s.rotateLeft(p)
	}
	if s.getBFactor(p) == -2 {
		if s.getBFactor(p.childLeft) > 0 {
			p.childLeft = s.rotateLeft(p.childLeft)
		}
		return s.rotateRight(p)
	}
	return p
}

func (s *AVLBST) Size() int {
	return s.size
}

func (s *AVLBST) IsEmpty() bool {
	return s.Size() == 0
}

func (s *AVLBST) searchNodeAndInsert(node *NodeAVL, key int, value string) *NodeAVL {
	if node == nil {
		s.size++
		return &NodeAVL{key, value, 1, node, nil, nil}
	}
	if key < node.key {
		node.childLeft = s.searchNodeAndInsert(node.childLeft, key, value)
	} else if key > node.key {
		node.childRight = s.searchNodeAndInsert(node.childRight, key, value)
	} else {
		node.value = value
	}

	return s.balance(node)
}

func (s *AVLBST) Insert(key int, value string) {
	if s.root == nil {
		s.root = &NodeAVL{key, value, 1, nil, nil, nil}
		s.size++
	} else {
		node := s.searchNodeAndInsert(s.root, key, value)
		if node == nil {
			panic("ERROR in searchNodeAndInsert")
		}
	}
}

func (s *AVLBST) findmin(node *NodeAVL) *NodeAVL {
	if node.childLeft != nil {
		return s.findmin(node.childLeft)
	}
	return node
}

func (s *AVLBST) removemin(node *NodeAVL) *NodeAVL {
	if node.childLeft == nil {
		return node.childRight
	}
	node.childLeft = s.removemin(node.childLeft)
	return s.balance(node)
}

func (s *AVLBST) removeNode(node *NodeAVL, key int) *NodeAVL {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.childLeft = s.removeNode(node.childLeft, key)
	} else if key > node.key {
		node.childRight = s.removeNode(node.childRight, key)
	} else {
		s.size--
		q := node.childLeft
		r := node.childRight
		if r == nil {
			if node == s.root {
				s.root = q
			}
			return q
		}
		if node == s.root {
			s.root = r
		}
		min := s.findmin(r)
		min.childRight = s.removemin(r)
		min.childLeft = q
		return s.balance(min)
	}
	return s.balance(node)
}

func (s *AVLBST) Remove(key int) {
	s.removeNode(s.root, key)
}

func (s *AVLBST) searchNode(node *NodeAVL, key int) string {
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

func (s *AVLBST) Search(key int) string {
	return s.searchNode(s.root, key)
}

func (s *AVLBST) printNode(node *NodeAVL, level int) {
	if node == nil {
		return
	}
	level += 10
	s.printNode(node.childRight, level)
	fmt.Println("")
	for i := 10; i < level; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(node.key, "("+strconv.Itoa(node.height)+")")

	s.printNode(node.childLeft, level)

}

func (s *AVLBST) Print() {
	s.printNode(s.root, 0)
}

func (s *AVLBST) collectKey(node *NodeAVL, array *[]int) {
	if node == nil {
		return
	}
	*array = append(*array, node.key)
	s.collectKey(node.childLeft, array)
	s.collectKey(node.childRight, array)
}

func (s *AVLBST) ListKey() []int {
	result := make([]int, 0)
	s.collectKey(s.root, &result)
	return result
}
