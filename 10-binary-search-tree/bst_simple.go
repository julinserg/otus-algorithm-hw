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
		s.size++
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
		s.size++
	} else {
		node := s.searchNodeAndInsert(s.root, key)
		if node != nil {
			node.value = value
		} else {
			panic("ERROR in searchNodeAndInsert")
		}
	}

}

func (s *SimpleBST) findMaximalNode(nodeCurent *NodeBST, nodeMax *NodeBST) *NodeBST {
	if nodeCurent == nil {
		return nodeMax
	}
	if nodeCurent.key > nodeMax.key {
		nodeMax = nodeCurent
	}
	return s.findMaximalNode(nodeCurent.childRight, nodeMax)
}

func (s *SimpleBST) removeNodeAnalysis(node *NodeBST, isLeft bool) {
	if node.childLeft == nil && node.childRight == nil {
		if node.parent == nil {
			s.root = nil
		} else {
			if isLeft {
				node.parent.childLeft = nil
			} else {
				node.parent.childRight = nil
			}
		}
	} else if node.childLeft == nil || node.childRight == nil {
		var nodeS *NodeBST
		if node.childLeft != nil {
			nodeS = node.childLeft
		} else {
			nodeS = node.childRight
		}
		if isLeft {
			node.parent.childLeft = nodeS
		} else {
			node.parent.childRight = nodeS
		}
	} else {
		nodeForRemove := s.findMaximalNode(node.childLeft, node.childLeft)
		node.key = nodeForRemove.key
		node.value = nodeForRemove.value
		isLeftForRemove := false
		if nodeForRemove.key == node.childLeft.key {
			isLeftForRemove = true
		}
		s.removeNodeAnalysis(nodeForRemove, isLeftForRemove)
	}
}

func (s *SimpleBST) removeNode(node *NodeBST, key int, isLeft bool) {
	if node == nil {
		return
	}
	if key == node.key {
		s.removeNodeAnalysis(node, isLeft)
		s.size--
	} else if key > node.key {
		s.removeNode(node.childRight, key, false)
	} else {
		s.removeNode(node.childLeft, key, true)
	}
}

func (s *SimpleBST) Remove(key int) {
	s.removeNode(s.root, key, false)
}

func (s *SimpleBST) searchNode(node *NodeBST, key int) string {
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

func (s *SimpleBST) Search(key int) string {
	return s.searchNode(s.root, key)
}

func (s *SimpleBST) printNode(node *NodeBST, level int) {
	if node == nil {
		return
	}
	level += 10
	s.printNode(node.childRight, level)
	fmt.Println("")
	for i := 10; i < level; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(node.key)

	s.printNode(node.childLeft, level)

}

func (s *SimpleBST) Print() {
	s.printNode(s.root, 0)
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
