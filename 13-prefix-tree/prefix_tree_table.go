package p13prefixtree

type ItemType interface {
	int32 | int64 | float32 | float64 | string
}

type Element[V ItemType] struct {
	value        V
	isValueExist bool
	next         [128]*Element[V]
}

type Siblings[V ItemType] struct {
	node                *Element[V]
	countNotNilSiblings int
	indexNilSibling     int
}

type PrefixTreeTable[V ItemType] struct {
	root *Element[V]
	size int
}

func (s *PrefixTreeTable[V]) init() {
	s.root = &Element[V]{}
}

func (s *PrefixTreeTable[V]) Size() int {
	return s.size
}

func (s *PrefixTreeTable[V]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *PrefixTreeTable[V]) Insert(key string, value V) {
	if s.IsEmpty() {
		s.init()
	}
	node := s.root
	for i := 0; i < len(key); i++ {
		c := key[i]
		if node.next[c] == nil {
			node.next[c] = &Element[V]{}
		}
		node = node.next[c]
	}
	if node.isValueExist {
		node.value = value
		return
	}
	node.isValueExist = true
	node.value = value
	s.size++
}

func countNotNilElement[V ItemType](array [128]*Element[V]) int {
	countNotNil := 0
	for i := 0; i < len(array); i++ {
		if array[i] != nil {
			countNotNil++
		}
	}
	return countNotNil
}

func (s *PrefixTreeTable[V]) Remove(key string) bool {
	if s.IsEmpty() {
		return false
	}
	node := s.root
	nodePrev := node
	countPrev := 0
	arrayNode := make([]Siblings[V], 0)
	for i := 0; i < len(key); i++ {
		c := key[i]
		node = node.next[c]
		if node == nil {
			return false
		}
		count := countNotNilElement(node.next)
		arrayNode = append(arrayNode, Siblings[V]{nodePrev, countPrev, int(c)})
		nodePrev = node
		countPrev = count

	}
	if !node.isValueExist {
		return false
	}
	count := countNotNilElement(node.next)
	if count > 0 {
		node.isValueExist = false
		s.size--
		return true
	}
	for i := len(arrayNode) - 1; i >= 0; i-- {
		arrayNode[i].node.next[arrayNode[i].indexNilSibling] = nil
		if arrayNode[i].countNotNilSiblings >= 2 {
			break
		}
		if arrayNode[i].node.isValueExist {
			break
		}

	}
	s.size--
	return true
}

func (s *PrefixTreeTable[V]) Search(key string) V {
	var value V
	if s.IsEmpty() {
		return value
	}
	node := s.root
	for i := 0; i < len(key); i++ {
		c := key[i]
		node = node.next[c]
		if node == nil {
			return value
		}
	}
	if node.isValueExist {
		return node.value
	}
	return value
}

func (s *PrefixTreeTable[V]) Print() {

}

func (s *PrefixTreeTable[V]) ListKey() []string {
	result := make([]string, 0)

	return result
}
