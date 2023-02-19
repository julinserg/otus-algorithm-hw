package p04dynamicarray

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := &Queue[string]{}
	q.Push("111")
	q.Push("222")
	q.Push("333")
	require.Equal(t, true, reflect.DeepEqual([]string{"333", "222", "111"}, q.List()))
	require.Equal(t, "111", q.Pop())
	require.Equal(t, true, reflect.DeepEqual([]string{"333", "222"}, q.List()))
	q.Push("111")
	require.Equal(t, true, reflect.DeepEqual([]string{"111", "333", "222"}, q.List()))
	q.Push("000")
	require.Equal(t, "222", q.Pop())
	require.Equal(t, "333", q.Pop())
	require.Equal(t, "111", q.Pop())
	require.Equal(t, "000", q.Pop())
	require.Equal(t, true, q.IsEmpty())
	q.Push("111")
	q.Push("222")
	q.Push("333")
	require.Equal(t, true, reflect.DeepEqual([]string{"333", "222", "111"}, q.List()))
}
