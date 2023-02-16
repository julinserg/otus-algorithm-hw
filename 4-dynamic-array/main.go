package main

import "fmt"

type ArrayItemType interface {
	int32 | int64 | float32 | float64 | string
}

type IArray[T ArrayItemType] interface {
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

func NewSingleArrayString(size int) IArray[string] {
	sa := &SingleArray[string]{}
	sa.m_array = make([]string, size)
	return sa
}

func main() {
	sa := NewSingleArrayString(2)
	sa.Set(0, "123")
	sa.Set(1, "456")
	fmt.Printf("SingleArray size %d \n", sa.Size())
	fmt.Printf("SingleArray get %s \n", sa.Get(0))
	fmt.Printf("SingleArray get %s \n", sa.Get(1))
	sa.Add(1, "789")
	fmt.Printf("SingleArray get %s \n", sa.Get(2))
}
