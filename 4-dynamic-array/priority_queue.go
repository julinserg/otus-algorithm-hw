package p04dynamicarray

type NodePriority[T ArrayItemType] struct {
	itemQueue Queue[T]
	priority  int
	next      *NodePriority[T]
}

type PriorityQueue[T ArrayItemType] struct {
	front *NodePriority[T]
	back  *NodePriority[T]
	size  int
}

func (s *PriorityQueue[T]) Size() int {
	return s.size
}

func (s *PriorityQueue[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *PriorityQueue[T]) Push(item T, priority int) {
	node := s.back
	var nodeForInsert *NodePriority[T]
	var nodeForInsertPrev *NodePriority[T]
	isNodeForInsertFind := false
	for node != nil {
		if node.priority == priority {
			node.itemQueue.Push(item)
			return
		} else if node.priority > priority && !isNodeForInsertFind {
			nodeForInsert = nodeForInsertPrev
			isNodeForInsertFind = true
		}
		nodeForInsertPrev = node
		node = node.next
	}
	queueNew := Queue[T]{}
	queueNew.Push(item)
	nodeNew := &NodePriority[T]{queueNew, priority, nil}
	if s.back == nil {
		s.front = nodeNew
		s.back = nodeNew
	} else if !isNodeForInsertFind {
		s.front.next = nodeNew
		s.front = nodeNew
	} else if isNodeForInsertFind && nodeForInsert == nil {
		oldBack := s.back
		s.back = nodeNew
		nodeNew.next = oldBack
	} else if nodeForInsert.next == nil {
		nodeForInsert.next = nodeNew
		s.front = nodeNew
	} else {
		oldNext := nodeForInsert.next
		nodeForInsert.next = nodeNew
		nodeNew.next = oldNext
	}
	s.size++
}

func (s *PriorityQueue[T]) Pop() T {
	if s.IsEmpty() {
		panic("Queue is empty")
	}
	result := s.front.itemQueue.Pop()
	if !s.front.itemQueue.IsEmpty() {
		return result
	}
	node := s.back
	var nodePreFront *NodePriority[T]
	for node.next != nil {
		nodePreFront = node
		node = node.next
	}

	if nodePreFront != nil {
		nodePreFront.next = nil
	} else {
		s.back = nil
	}
	s.front = nodePreFront
	s.size--
	return result
}

func (s *PriorityQueue[T]) Front() T {
	return s.front.itemQueue.Front()
}

func (s *PriorityQueue[T]) Back() T {
	return s.back.itemQueue.Back()
}

func (s *PriorityQueue[T]) List() []T {
	result := make([]T, 0)
	node := s.back
	for node != nil {
		result = append(result, node.itemQueue.List()...)
		node = node.next
	}
	return result
}
