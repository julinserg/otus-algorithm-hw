package p04dynamicarray

type Node[T ArrayItemType] struct {
	item T
	next *Node[T]
}

type Queue[T ArrayItemType] struct {
	front *Node[T]
	back  *Node[T]
	size  int
}

func (s *Queue[T]) Size() int {
	return s.size
}

func (s *Queue[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Queue[T]) Push(item T) {
	node := &Node[T]{item, nil}
	if s.IsEmpty() {
		s.front = node
		s.back = node
	} else {
		oldBack := s.back
		s.back = node
		node.next = oldBack
	}
	s.size++
}

func (s *Queue[T]) Pop() T {
	if s.IsEmpty() {
		panic("Queue is empty")
	}
	node := s.back
	var nodePreFront *Node[T]
	for node.next != nil {
		nodePreFront = node
		node = node.next
	}
	frontOld := s.Front()
	if nodePreFront != nil {
		nodePreFront.next = nil
	} else {
		s.back = nil
	}
	s.front = nodePreFront
	s.size--
	return frontOld
}

func (s *Queue[T]) Front() T {
	return s.front.item
}

func (s *Queue[T]) Back() T {
	return s.back.item
}

func (s *Queue[T]) List() []T {
	result := make([]T, 0)
	node := s.back
	for node != nil {
		result = append(result, node.item)
		node = node.next
	}
	return result
}
