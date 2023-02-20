package p04dynamicarray

type GoArray[T ArrayItemType] struct {
	m_array []T
}

func (s *GoArray[T]) Size() int {
	return len(s.m_array)
}

func (s *GoArray[T]) Get(index int) T {
	return s.m_array[index]
}

func (s *GoArray[T]) Set(index int, item T) {
	s.m_array[index] = item
}

func (s *GoArray[T]) Add(index int, item T) {
	if index < 0 {
		panic("Error index")
	} else if index > len(s.m_array) {
		index = len(s.m_array)
	}
	s.m_array = append(s.m_array[:index], append([]T{item}, s.m_array[index:]...)...)
}

func (s *GoArray[T]) Remove(index int) T {
	if index < 0 {
		panic("Error index")
	} else if index > len(s.m_array) {
		index = len(s.m_array)
	}
	result := s.m_array[index]
	copy(s.m_array[index:], s.m_array[index+1:])
	s.m_array = s.m_array[:len(s.m_array)-1]
	return result
}

func (s *GoArray[T]) Create(size int) {
	s.m_array = make([]T, size)
}
