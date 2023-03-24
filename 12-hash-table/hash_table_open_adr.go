package p12hashtable

import (
	"fmt"
)

type ElementOpenAdr[K, V ItemType] struct {
	key      KeyElement[K]
	value    V
	isFill   bool
	isRemove bool
}

type HashTableOpenAdr[K, V ItemType] struct {
	table    []ElementOpenAdr[K, V]
	size     int
	capacity int
}

func (s *HashTableOpenAdr[K, V]) init() {
	if s.capacity == 0 {
		s.capacity = defaultBucketsCount
	}
	s.table = make([]ElementOpenAdr[K, V], s.capacity)
}

func (s *HashTableOpenAdr[K, V]) Size() int {
	return s.size
}

func (s *HashTableOpenAdr[K, V]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *HashTableOpenAdr[K, V]) rehash(capacity int) {
	s.capacity = capacity
	s.size = 0
	tableOld := s.table

	for i := 0; i < len(tableOld); i++ {
		item := tableOld[i]
		if item.isFill && !item.isRemove {
			s.Insert(item.key, item.value)
		}
	}
}

func (s *HashTableOpenAdr[K, V]) index(key KeyElement[K], i int) uint32 {
	return (key.hash() + uint32(i) + uint32(i)*uint32(i)) % uint32(s.capacity)
}

func (s *HashTableOpenAdr[K, V]) Insert(key KeyElement[K], value V) {
	if s.IsEmpty() {
		s.init()
	}
	if float64(s.size) >= float64(s.capacity)*defaultBucketsFill {
		s.rehash(s.capacity * 2)
	}

	for i := 0; i < len(s.table); i++ {
		indexTable := s.index(key, i)
		st := s.table[indexTable]
		if st.isFill && !st.isRemove {
			if st.key == key {
				st.value = value
				return
			} else {
				continue
			}
		}
		s.table[indexTable] = ElementOpenAdr[K, V]{key, value, true, false}
		s.size++
		return
	}
}

func (s *HashTableOpenAdr[K, V]) Remove(key KeyElement[K]) bool {
	if s.IsEmpty() {
		return false
	}
	if s.size*4 < s.capacity {
		s.rehash(s.size * 2)
	}
	for i := 0; i < len(s.table); i++ {
		indexTable := s.index(key, i)
		st := s.table[indexTable]
		if st.isFill && !st.isRemove && st.key == key {
			s.table[indexTable].isRemove = true
			s.size--
			return true
		}
	}

	return false
}

func (s *HashTableOpenAdr[K, V]) Search(key KeyElement[K]) V {
	var value V
	if s.IsEmpty() {
		return value
	}
	for i := 0; i < len(s.table); i++ {
		indexTable := s.index(key, i)
		st := s.table[indexTable]
		if st.isFill && !st.isRemove && st.key == key {
			return st.value
		}
	}
	return value
}

func (s *HashTableOpenAdr[K, V]) Print() {
	for i := 0; i < len(s.table); i++ {
		if !s.table[i].isFill || s.table[i].isRemove {
			continue
		}
		fmt.Println("key =", s.table[i].key.key(), ", value =", s.table[i].value)
	}
}

func (s *HashTableOpenAdr[K, V]) ListKey() []KeyElement[K] {
	result := make([]KeyElement[K], 0)
	for i := 0; i < s.capacity; i++ {
		if !s.table[i].isFill || s.table[i].isRemove {
			continue
		}
		item := s.table[i]
		result = append(result, item.key)
	}
	return result
}
