package p10bst

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleBST(t *testing.T) {
	bst := &SimpleBST{}
	bst.Insert(1, "1")
	bst.Insert(2, "2")
	bst.Insert(3, "3")
	fmt.Println(bst.ListKey())
	require.Equal(t, true, reflect.DeepEqual([]int{1, 2, 3}, bst.ListKey()))
	bst.Print()
}
