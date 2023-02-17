package p04dynamicarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewSingleArrayString(size int) IArray[string] {
	sa := &SingleArray[string]{}
	sa.Create(size)
	return sa
}

func testItemAdd(t *testing.T, array IArray[string]) {
	array.Set(0, "123")
	array.Set(1, "456")
	require.Equal(t, 2, array.Size())
	require.Equal(t, "123", array.Get(0))
	require.Equal(t, "456", array.Get(1))
	array.Add(1, "789")
	require.Equal(t, 3, array.Size())
	require.Equal(t, "123", array.Get(0))
	require.Equal(t, "789", array.Get(1))
	require.Equal(t, "456", array.Get(2))
	array.Add(0, "000")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "456", array.Get(3))
	array.Add(4, "999")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "456", array.Get(3))
	require.Equal(t, "999", array.Get(4))
	array.Add(3, "555")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "555", array.Get(3))
	require.Equal(t, "456", array.Get(4))
	require.Equal(t, "999", array.Get(5))
	array.Add(100, "777")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "555", array.Get(3))
	require.Equal(t, "456", array.Get(4))
	require.Equal(t, "999", array.Get(5))
	require.Equal(t, "777", array.Get(6))
}

func TestSingleArray(t *testing.T) {
	sa := NewSingleArrayString(2)
	testItemAdd(t, sa)
}
