package p04dynamicarray

type FactorArray[T ArrayItemType] struct {
	m_array  []T
	m_size   int
	m_factor int
}

func (s *FactorArray[T]) Size() int {
	return s.m_size
}

func (s *FactorArray[T]) Get(index int) T {
	if index > s.Size()-1 {
		panic("Error index")
	}
	return s.m_array[index]
}

func (s *FactorArray[T]) Set(index int, item T) {
	if index > s.Size()-1 {
		panic("Error index")
	}
	s.m_array[index] = item
}

func (s *FactorArray[T]) resizeUp(index int) {
	arrayNew := make([]T, s.Size()*s.m_factor+1)
	indexOld := 0
	indexNew := 0
	for indexOld < s.Size() {
		if indexOld == index && indexOld == indexNew {
			indexNew++
			continue
		}
		arrayNew[indexNew] = s.m_array[indexOld]
		indexNew++
		indexOld++
	}
	s.m_array = arrayNew
}

func (s *FactorArray[T]) resizeDown(index int) T {
	arrayNew := make([]T, len(s.m_array)-1)
	indexOld := 0
	indexNew := 0
	var item T
	for indexOld < s.Size() {
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

func (s *FactorArray[T]) Add(index int, item T) {
	if index < 0 {
		panic("Error index")
	} else if index > s.Size() {
		index = s.Size()
	}
	if !(len(s.m_array) > s.Size() && index == s.Size()) {
		s.resizeUp(index)
	}
	s.m_array[index] = item
	s.m_size++
}

func (s *FactorArray[T]) Remove(index int) T {
	if index < 0 {
		panic("Error index")
	} else if index >= s.Size() {
		index = s.Size() - 1
	}
	var item T
	if index == s.Size()-1 {
		item = s.m_array[index]
	} else {
		item = s.resizeDown(index)
	}
	s.m_size--
	return item
}

func (s *FactorArray[T]) Create(size int) {
	s.m_array = make([]T, size)
	s.m_size += size
	s.m_factor = 2
}
