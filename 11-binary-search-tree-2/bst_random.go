package p10bst2

import (
	"fmt"
	"strconv"
)

type NodeRandomBST struct {
	key        int
	value      string
	height     int
	parent     *NodeRandomBST
	childLeft  *NodeRandomBST
	childRight *NodeRandomBST
}

// https://habr.com/ru/post/145388/
type RandomBST struct {
	root *NodeRandomBST
	size int
}

func (s *RandomBST) getHeight(p *NodeRandomBST) int {
	if p == nil {
		return 0
	} else {
		return p.height
	}
}

func (s *RandomBST) getBFactor(p *NodeRandomBST) int {
	return s.getHeight(p.childRight) - s.getHeight(p.childLeft)
}

func (s *RandomBST) fixHeight(p *NodeRandomBST) {
	hl := s.getHeight(p.childLeft)
	hr := s.getHeight(p.childRight)
	if hl > hr {
		p.height = hl + 1
	} else {
		p.height = hr + 1
	}
}

func (s *RandomBST) rotateRight(p *NodeRandomBST) *NodeRandomBST {
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

func (s *RandomBST) rotateLeft(q *NodeRandomBST) *NodeRandomBST {
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

func (s *RandomBST) balance(p *NodeRandomBST) *NodeRandomBST {
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

func (s *RandomBST) Size() int {
	return s.size
}

func (s *RandomBST) IsEmpty() bool {
	return s.Size() == 0
}

func (s *RandomBST) searchNodeAndInsert(node *NodeRandomBST, key int, value string) *NodeRandomBST {
	if node == nil {
		s.size++
		return &NodeRandomBST{key, value, 1, node, nil, nil}
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

func (s *RandomBST) Insert(key int, value string) {
	if s.root == nil {
		s.root = &NodeRandomBST{key, value, 1, nil, nil, nil}
		s.size++
	} else {
		node := s.searchNodeAndInsert(s.root, key, value)
		if node == nil {
			panic("ERROR in searchNodeAndInsert")
		}
	}
}

func (s *RandomBST) findmin(node *NodeRandomBST) *NodeRandomBST {
	if node.childLeft != nil {
		return s.findmin(node.childLeft)
	}
	return node
}

func (s *RandomBST) removemin(node *NodeRandomBST) *NodeRandomBST {
	if node.childLeft == nil {
		return node.childRight
	}
	node.childLeft = s.removemin(node.childLeft)
	return s.balance(node)
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
	fmt.Println(node.key, "("+strconv.Itoa(node.height)+")")

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
