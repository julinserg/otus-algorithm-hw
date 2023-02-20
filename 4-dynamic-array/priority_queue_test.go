package p04dynamicarray

import (
	"fmt"
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

func TestPriorityQueue(t *testing.T) {
	pq := &PriorityQueue[string]{}
	pq.Push("111", 1)
	pq.Push("222", 1)
	pq.Push("333", 1)
	require.Equal(t, true, reflect.DeepEqual([]string{"333", "222", "111"}, pq.List()))
	require.Equal(t, "111", pq.Pop())
	require.Equal(t, true, reflect.DeepEqual([]string{"333", "222"}, pq.List()))
	pq.Push("111", 1)
	require.Equal(t, true, reflect.DeepEqual([]string{"111", "333", "222"}, pq.List()))
	pq.Push("000", 1)
	require.Equal(t, "222", pq.Pop())
	require.Equal(t, "333", pq.Pop())
	require.Equal(t, "111", pq.Pop())
	require.Equal(t, "000", pq.Pop())
	require.Equal(t, true, pq.IsEmpty())
	pq.Push("111", 2)
	pq.Push("222", 2)
	pq.Push("333", 2)
	pq.Push("444", 1)
	pq.Push("555", 1)
	pq.Push("666", 1)
	pq.Push("777", 3)
	pq.Push("888", 3)
	pq.Push("999", 3)
	require.Equal(t, 3, pq.Size())
	fmt.Println(pq.List())
	require.Equal(t, true, reflect.DeepEqual([]string{"666", "555", "444", "333",
		"222", "111", "999", "888", "777"}, pq.List()))
	require.Equal(t, "777", pq.Pop())
	require.Equal(t, "888", pq.Pop())
	require.Equal(t, "999", pq.Pop())
	require.Equal(t, "111", pq.Pop())
	require.Equal(t, "222", pq.Pop())
	require.Equal(t, "333", pq.Pop())
	require.Equal(t, "444", pq.Pop())
	require.Equal(t, "555", pq.Pop())
	require.Equal(t, "666", pq.Pop())
	require.Equal(t, true, pq.IsEmpty())
	pq.Push("111", 0)
	pq.Push("222", 9)
	pq.Push("333", 1)
	pq.Push("444", 8)
	pq.Push("555", 2)
	pq.Push("666", 7)
	pq.Push("777", 3)
	pq.Push("888", 6)
	pq.Push("999", 4)
	require.Equal(t, "222", pq.Pop())
	require.Equal(t, "444", pq.Pop())
	require.Equal(t, "666", pq.Pop())
	require.Equal(t, "888", pq.Pop())
	require.Equal(t, "999", pq.Pop())
	require.Equal(t, "777", pq.Pop())
	require.Equal(t, "555", pq.Pop())
	require.Equal(t, "333", pq.Pop())
	require.Equal(t, "111", pq.Pop())
}
