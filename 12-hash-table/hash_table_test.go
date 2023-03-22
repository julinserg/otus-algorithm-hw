package p12hashtable

import (
	"hash/fnv"
	"testing"
)

type KeyElementString struct {
	keyStr string
}

func (s KeyElementString) key() string {
	return s.keyStr
}

func (s KeyElementString) hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(s.keyStr))
	return h.Sum32()
}

func TestHashTable(t *testing.T) {
	hashTable := &HashTable[string, string]{}
	hashTable.Insert(KeyElementString{"mykey1"}, "mavalue1")
	hashTable.Insert(KeyElementString{"mykey2"}, "mavalue2")
	hashTable.Print()
}
