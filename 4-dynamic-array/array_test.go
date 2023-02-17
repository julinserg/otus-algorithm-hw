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

func TestSingleArray(t *testing.T) {
	sa := NewSingleArrayString(2)
	sa.Set(0, "123")
	sa.Set(1, "456")
	require.Equal(t, 2, sa.Size())
	require.Equal(t, "123", sa.Get(0))
	require.Equal(t, "456", sa.Get(1))
	sa.Add(1, "789")
	require.Equal(t, 3, sa.Size())
	require.Equal(t, "789", sa.Get(2))

}
