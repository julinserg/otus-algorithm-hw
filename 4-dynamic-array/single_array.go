package p04dynamicarray

type ArrayItemType interface {
	int32 | int64 | float32 | float64 | string
}

type IArray[T ArrayItemType] interface {
	Create(int)
	Size() int
	Get(int) T
	Set(int, T)
	Add(int, T)
	//Remove(int) T
}

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

func (s *SingleArray[T]) resize() {
	arrayNew := make([]T, s.Size()+1)
	for i := 0; i < len(s.m_array); i++ {
		arrayNew[i] = s.m_array[i]
	}
	s.m_array = arrayNew
}

func (s *SingleArray[T]) Add(index int, item T) {
	s.resize()
	s.m_array[s.Size()-1] = item
}

func (s *SingleArray[T]) Create(size int) {
	s.m_array = make([]T, size)
}
