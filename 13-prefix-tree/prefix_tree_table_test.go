package p13prefixtree

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func testEq(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPrefixTreeTable(t *testing.T) {
	hashTable := &PrefixTreeTable[string]{}
	hashTable.Insert("mykey1", "mavalue1")
	hashTable.Insert("mykey2", "mavalue2")
	hashTable.Print()

	require.Equal(t, "mavalue2", hashTable.Search("mykey2"))
	require.Equal(t, "mavalue1", hashTable.Search("mykey1"))
	require.Equal(t, "", hashTable.Search("mykey0"))

	require.Equal(t, 2, hashTable.Size())
	require.Equal(t, false, hashTable.Remove("mykey0"))
	require.Equal(t, true, hashTable.Remove("mykey1"))
	require.Equal(t, true, hashTable.Remove("mykey2"))
	require.Equal(t, "", hashTable.Search("mykey2"))
	require.Equal(t, "", hashTable.Search("mykey1"))
	require.Equal(t, "", hashTable.Search("mykey"))
	require.Equal(t, "", hashTable.Search("myke"))
	require.Equal(t, "", hashTable.Search("myk"))
	require.Equal(t, "", hashTable.Search("my"))
	require.Equal(t, "", hashTable.Search("m"))

	require.Equal(t, 0, hashTable.Size())

	hashTable.Insert("mykey1", "mavalue1")
	hashTable.Insert("mykey2", "mavalue2")
	hashTable.Insert("myke", "myke")
	hashTable.Insert("myk", "myk")
	hashTable.Insert("m", "m")
	require.Equal(t, "mavalue2", hashTable.Search("mykey2"))
	require.Equal(t, "mavalue1", hashTable.Search("mykey1"))
	require.Equal(t, "myk", hashTable.Search("myk"))
	require.Equal(t, "myke", hashTable.Search("myke"))
	require.Equal(t, true, hashTable.Remove("mykey1"))
	require.Equal(t, true, hashTable.Remove("mykey2"))
	require.Equal(t, "", hashTable.Search("mykey2"))
	require.Equal(t, "", hashTable.Search("mykey1"))
	require.Equal(t, "myk", hashTable.Search("myk"))
	require.Equal(t, "myke", hashTable.Search("myke"))
	require.Equal(t, "m", hashTable.Search("m"))
}

func TestPrefixTreeTableWithRehash(t *testing.T) {
	hashTable := &PrefixTreeTable[string]{}
	N := 100000
	for i := 0; i < N; i++ {
		r := rand.Intn(N)
		rStr := strconv.Itoa(r)
		hashTable.Insert(rStr, rStr)
		hashTable.Insert(strconv.Itoa(i), strconv.Itoa(i))
		hashTable.Insert(rStr, rStr)
	}
	require.Equal(t, N, hashTable.Size())

	hashTable.Insert("mykey1", "mavalue1")
	hashTable.Insert("mykey2", "mavalue2")
	require.Equal(t, "mavalue2", hashTable.Search("mykey2"))
	require.Equal(t, "mavalue1", hashTable.Search("mykey1"))
	require.Equal(t, "", hashTable.Search("mykey0"))
	require.Equal(t, true, hashTable.Remove("mykey1"))
	require.Equal(t, true, hashTable.Remove("mykey2"))

	require.Equal(t, N, hashTable.Size())

	for i := 0; i < N; i++ {
		require.NotEqual(t, "", hashTable.Search(strconv.Itoa(i)))
	}

	for i := 0; i < N; i++ {
		require.Equal(t, true, hashTable.Remove(strconv.Itoa(i)))
	}
	require.Equal(t, 0, hashTable.Size())
}
