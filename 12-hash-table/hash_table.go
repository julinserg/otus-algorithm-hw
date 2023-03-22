package p12hashtable

import "fmt"

type ItemType interface {
	int32 | int64 | float32 | float64 | string
}

const defaultBucketsCount = 11

type KeyElement[K ItemType] interface {
	key() K
	hash() uint32
}

type Element[K, V ItemType] struct {
	key   KeyElement[K]
	value V
	next  *Element[K, V]
}

type HashTable[K, V ItemType] struct {
	buckets  []*Element[K, V]
	size     int
	capacity int
}

func (s *HashTable[K, V]) init() {
	s.buckets = make([]*Element[K, V], defaultBucketsCount)
	s.capacity = defaultBucketsCount
}

func (s *HashTable[K, V]) Size() int {
	return s.size
}

func (s *HashTable[K, V]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *HashTable[K, V]) Insert(key KeyElement[K], value V) {
	if s.IsEmpty() {
		s.init()
	}
	indexBucket := key.hash() % uint32(s.capacity)
	if s.buckets[indexBucket] == nil {
		s.buckets[indexBucket] = &Element[K, V]{key, value, nil}
	} else {
		item := s.buckets[indexBucket]
		for item != nil {
			if item.next == nil {
				item.next = &Element[K, V]{key, value, nil}
			}
			item = item.next
		}
	}
	s.size++

}

func (s *HashTable[K, V]) Remove(key KeyElement[K]) {
	if s.IsEmpty() {
		return
	}
	s.size--
}

func (s *HashTable[K, V]) Search(key KeyElement[K]) V {
	var value V
	if s.IsEmpty() {
		return value
	}
	return value

}

func (s *HashTable[K, V]) Print() {
	for i := 0; i < s.capacity; i++ {
		item := s.buckets[i]
		for item != nil {
			fmt.Println("key =", item.key.key(), ", value =", item.value)
			item = item.next
		}
	}
}

func (s *HashTable[K, V]) ListKey() []int {
	return []int{}
}
