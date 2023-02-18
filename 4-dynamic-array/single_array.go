package p04dynamicarray

type SingleArray[T ArrayItemType] struct {
	m_array []T
}

func (s *SingleArray[T]) Size() int {
	return len(s.m_array)
}

func (s *SingleArray[T]) Get(index int) T {
	return s.m_array[index]
}

func (s *SingleArray[T]) Set(index int, item T) {
	s.m_array[index] = item
}

func (s *SingleArray[T]) resizeUp(index int) {
	arrayNew := make([]T, s.Size()+1)
	indexOld := 0
	indexNew := 0
	for indexNew < len(arrayNew) {
		if indexNew == index {
			indexNew++
			continue
		}
		arrayNew[indexNew] = s.m_array[indexOld]
		indexNew++
		indexOld++
	}
	s.m_array = arrayNew
}

func (s *SingleArray[T]) resizeDown(index int) T {
	arrayNew := make([]T, s.Size()-1)
	indexOld := 0
	indexNew := 0
	var item T
	for indexOld < len(s.m_array) {
		if indexOld == index {
			item = s.m_array[index]
			indexOld++
			continue
		}
		arrayNew[indexNew] = s.m_array[indexOld]
		indexNew++
		indexOld++
	}
	s.m_array = arrayNew
	return item
}

func (s *SingleArray[T]) Add(index int, item T) {
	if index < 0 {
		panic("Error index")
	} else if index > len(s.m_array) {
		index = len(s.m_array)
	}

	s.resizeUp(index)
	s.m_array[index] = item
}

func (s *SingleArray[T]) Remove(index int) T {
	if index < 0 {
		panic("Error index")
	} else if index > len(s.m_array) {
		index = len(s.m_array)
	}

	return s.resizeDown(index)
}

func (s *SingleArray[T]) Create(size int) {
	s.m_array = make([]T, size)
}
