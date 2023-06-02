package p21rle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRle(t *testing.T) {
	encode := RleEncode([]byte{'A', 'A', 'B', 'A', 'A', 'B', 'A', 'A', 'A', 'B', 'A'})
	require.Equal(t, []byte{191 + 2, 'A', 'B', 191 + 2, 'A', 'B', 191 + 3, 'A', 'B', 'A'}, encode)

	decode := RleDecode([]byte{191 + 2, 'A', 'B', 191 + 2, 'A', 'B', 191 + 3, 'A', 'B', 'A'})
	require.Equal(t, []byte{'A', 'A', 'B', 'A', 'A', 'B', 'A', 'A', 'A', 'B', 'A'}, decode)

	encode1 := RleEncode([]byte{'A'})
	require.Equal(t, []byte{'A'}, encode1)

	decode1 := RleDecode([]byte{'A'})
	require.Equal(t, []byte{'A'}, decode1)

	encode2 := RleEncode([]byte{191 + '1'})
	require.Equal(t, []byte{191 + 1, 191 + '1'}, encode2)

	decode2 := RleDecode([]byte{191 + 1, 191 + '1'})
	require.Equal(t, []byte{191 + '1'}, decode2)
}
