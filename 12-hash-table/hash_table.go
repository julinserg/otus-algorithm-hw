package p12hashtable

import "fmt"

type ItemType interface {
	int32 | int64 | float32 | float64 | string
}

const defaultBucketsCount = 11
const defaultBucketsFill = 0.75

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
	if s.capacity == 0 {
		s.capacity = defaultBucketsCount
	}
	s.buckets = make([]*Element[K, V], s.capacity)
}

func (s *HashTable[K, V]) Size() int {
	return s.size
}

func (s *HashTable[K, V]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *HashTable[K, V]) rehash(capacity int) {
	s.capacity = capacity
	s.size = 0
	bucketsOld := s.buckets

	for i := 0; i < len(bucketsOld); i++ {
		item := bucketsOld[i]
		for item != nil {
			s.Insert(item.key, item.value)
			item = item.next
		}
	}
}

func (s *HashTable[K, V]) Insert(key KeyElement[K], value V) {
	if s.IsEmpty() {
		s.init()
	}
	if float64(s.size) >= float64(s.capacity)*defaultBucketsFill {
		s.rehash(s.capacity * 2)
	}
	indexBucket := key.hash() % uint32(s.capacity)
	if s.buckets[indexBucket] == nil {
		s.buckets[indexBucket] = &Element[K, V]{key, value, nil}
		s.size++
	} else {
		item := s.buckets[indexBucket]
		for item != nil {
			if item.key == key {
				item.value = value
				return
			}
			if item.next == nil {
				item.next = &Element[K, V]{key, value, nil}
				s.size++
				return
			}
			item = item.next
		}
	}
}

func (s *HashTable[K, V]) Remove(key KeyElement[K]) bool {
	if s.IsEmpty() {
		return false
	}
	if s.size*4 < s.capacity {
		s.rehash(s.size * 2)
	}
	indexBucket := key.hash() % uint32(s.capacity)
	if s.buckets[indexBucket] == nil {
		return false
	} else {
		item := s.buckets[indexBucket]
		prevItem := item
		for item != nil {
			if item.key == key {
				if item == prevItem {
					s.buckets[indexBucket] = item.next
				} else {
					prevItem.next = item.next
				}
				s.size--
				return true
			}
			prevItem = item
			item = item.next
		}
	}
	return false
}

func (s *HashTable[K, V]) Search(key KeyElement[K]) V {
	var value V
	if s.IsEmpty() {
		return value
	}
	indexBucket := key.hash() % uint32(s.capacity)
	if s.buckets[indexBucket] == nil {
		return value
	} else {
		item := s.buckets[indexBucket]
		for item != nil {
			if item.key == key {
				return item.value
			}
			item = item.next
		}
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

func (s *HashTable[K, V]) ListKey() []KeyElement[K] {
	result := make([]KeyElement[K], 0)
	for i := 0; i < s.capacity; i++ {
		item := s.buckets[i]
		for item != nil {
			result = append(result, item.key)
			item = item.next
		}
	}
	return result
}
