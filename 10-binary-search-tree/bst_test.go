package p10bst

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleBST(t *testing.T) {
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
	require.Equal(t, 4, bst.Size())
	bst.Insert(8, "str8")
	bst.Insert(1, "str1")
	bst.Print()
	bst.Remove(2)
	bst.Print()
	bst.Insert(2, "str2")
	bst.Print()
	bst.Remove(7)
	bst.Print()
}
