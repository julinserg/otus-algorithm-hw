package p10bst

import "fmt"

type NodeBST struct {
	key        int
	value      string
	parent     *NodeBST
	childLeft  *NodeBST
	childRight *NodeBST
}

type SimpleBST struct {
	root *NodeBST
	size int
}

func (s *SimpleBST) Size() int {
	return s.size
}

func (s *SimpleBST) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SimpleBST) searchNodeAndInsert(node *NodeBST, key int) *NodeBST {
	var child *NodeBST
	isRight := false
	if key > node.key {
		child = node.childRight
		isRight = true
	} else if key < node.key {
		child = node.childLeft
		isRight = false
	} else {
		return node
	}

	if child != nil {
		return s.searchNodeAndInsert(child, key)
	} else {
		child = &NodeBST{key, "", node, nil, nil}
		if isRight {
			node.childRight = child
		} else {
			node.childLeft = child
		}
		return child
	}
}

func (s *SimpleBST) Insert(key int, value string) {
	if s.root == nil {
		s.root = &NodeBST{key, value, nil, nil, nil}
	} else {
		node := s.searchNodeAndInsert(s.root, key)
		if node != nil {
			node.value = value
		} else {
			panic("ERROR in searchNodeAndInsert")
		}
	}
	s.size++
}

func (s *SimpleBST) Remove(key int) {
	s.size--
}

func (s *SimpleBST) Search(key int) string {
	return ""
}

func (s *SimpleBST) printNode(node *NodeBST) {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	s.printNode(node.childLeft)
	s.printNode(node.childRight)
}

func (s *SimpleBST) Print() {
	s.printNode(s.root)
}

func (s *SimpleBST) collectKey(node *NodeBST, array *[]int) {
	if node == nil {
		return
	}
	*array = append(*array, node.key)
	s.collectKey(node.childLeft, array)
	s.collectKey(node.childRight, array)
}

func (s *SimpleBST) ListKey() []int {
	result := make([]int, 0)
	s.collectKey(s.root, &result)
	return result
}
