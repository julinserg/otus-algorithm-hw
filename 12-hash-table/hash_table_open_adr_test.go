package p12hashtable

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashTableOpenAdr(t *testing.T) {
	hashTable := &HashTableOpenAdr[string, string]{}
	hashTable.Insert(KeyElementString{"mykey1"}, "mavalue1")
	hashTable.Insert(KeyElementString{"mykey2"}, "mavalue2")
	hashTable.Print()

	require.Equal(t, true, testEq(hashTable.ListKey(), []KeyElementString{{"mykey2"}, {"mykey1"}}))

	require.Equal(t, "mavalue2", hashTable.Search(KeyElementString{"mykey2"}))
	require.Equal(t, "mavalue1", hashTable.Search(KeyElementString{"mykey1"}))
	require.Equal(t, "", hashTable.Search(KeyElementString{"mykey0"}))

	require.Equal(t, 2, hashTable.Size())
	require.Equal(t, true, hashTable.Remove(KeyElementString{"mykey1"}))
	require.Equal(t, true, hashTable.Remove(KeyElementString{"mykey2"}))
	require.Equal(t, false, hashTable.Remove(KeyElementString{"mykey0"}))
	require.Equal(t, 0, hashTable.Size())

	hashTable.Insert(KeyElementString{"mykey1"}, "mavalue1")
	hashTable.Insert(KeyElementString{"mykey2"}, "mavalue2")
	require.Equal(t, true, testEq(hashTable.ListKey(), []KeyElementString{{"mykey2"}, {"mykey1"}}))
}

func TestHashTableOpenAdrWithRehash(t *testing.T) {
	hashTable := &HashTableOpenAdr[string, string]{}
	N := 100000
	for i := 0; i < N; i++ {
		r := rand.Intn(N)
		rStr := strconv.Itoa(r)
		hashTable.Insert(KeyElementString{rStr}, rStr)
		hashTable.Insert(KeyElementString{strconv.Itoa(i)}, strconv.Itoa(i))
		hashTable.Insert(KeyElementString{rStr}, rStr)
	}
	require.Equal(t, N, hashTable.Size())

	hashTable.Insert(KeyElementString{"mykey1"}, "mavalue1")
	hashTable.Insert(KeyElementString{"mykey2"}, "mavalue2")
	require.Equal(t, "mavalue2", hashTable.Search(KeyElementString{"mykey2"}))
	require.Equal(t, "mavalue1", hashTable.Search(KeyElementString{"mykey1"}))
	require.Equal(t, "", hashTable.Search(KeyElementString{"mykey0"}))
	require.Equal(t, true, hashTable.Remove(KeyElementString{"mykey1"}))
	require.Equal(t, true, hashTable.Remove(KeyElementString{"mykey2"}))

	require.Equal(t, N, hashTable.Size())

	for i := 0; i < N; i++ {
		require.NotEqual(t, "", hashTable.Search(KeyElementString{strconv.Itoa(i)}))
	}

	for i := 0; i < N; i++ {
		require.Equal(t, true, hashTable.Remove(KeyElementString{strconv.Itoa(i)}))
	}
	require.Equal(t, 0, hashTable.Size())
}
