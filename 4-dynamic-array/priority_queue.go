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

}

func (s *PriorityQueue[T]) Pop() T {
	if s.IsEmpty() {
		panic("Queue is empty")
	}
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
	for node.next != nil {
		result = append(result, node.item)
		node = node.next
	}
	result = append(result, node.item)
	return result
}
