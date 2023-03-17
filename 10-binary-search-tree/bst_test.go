package p10bst

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleBST1(t *testing.T) {
	bst := &SimpleBST{}

	bst.Insert(3, "str3")
	bst.Insert(2, "str2")
	bst.Insert(1, "str1")
	bst.Insert(7, "str7")
	bst.Insert(5, "str5")
	bst.Insert(8, "str8")
	fmt.Println(bst.ListKey())
	require.Equal(t, true, reflect.DeepEqual([]int{3, 2, 1, 7, 5, 8}, bst.ListKey()))
	bst.Print()

	require.Equal(t, "str8", bst.Search(8))
	require.Equal(t, "str3", bst.Search(3))
	require.Equal(t, "str2", bst.Search(2))

	require.Equal(t, 6, bst.Size())

	bst.Remove(1)
	bst.Remove(8)
	bst.Print()
	require.Equal(t, "", bst.Search(1))
	require.Equal(t, "", bst.Search(8))
	require.Equal(t, 4, bst.Size())
	bst.Insert(8, "str8")
	bst.Insert(1, "str1")
	bst.Print()
	bst.Remove(2)
	require.Equal(t, "", bst.Search(2))
	bst.Print()
	bst.Insert(2, "str2")
	bst.Print()
	bst.Remove(7)
	require.Equal(t, "", bst.Search(7))
	bst.Print()
}

func TestSimpleBST2(t *testing.T) {
	bst := &SimpleBST{}

	bst.Insert(30, "str30")
	bst.Insert(25, "str25")
	bst.Insert(33, "str33")
	bst.Insert(20, "str20")
	bst.Insert(27, "str27")
	bst.Insert(1, "str1")
	bst.Insert(23, "str23")
	bst.Insert(22, "str22")
	bst.Insert(28, "str28")
	bst.Insert(29, "str29")
	bst.Insert(35, "str35")
	bst.Insert(34, "str34")
	bst.Insert(38, "str38")
	bst.Print()

	require.Equal(t, "str38", bst.Search(38))

	bst.Remove(25)
	bst.Print()
	require.Equal(t, "", bst.Search(25))
}
