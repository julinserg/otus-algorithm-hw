package p21rle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRle(t *testing.T) {
	encode := rleEncode("AABAABAAABA")
	require.Equal(t, "A2B1A2B1A3B1A1", encode)

	decode := rleDecode("A2B1A2B1A3B1A1")
	require.Equal(t, "AABAABAAABA", decode)

	encode1 := rleEncode("A")
	require.Equal(t, "A1", encode1)

	decode1 := rleDecode("A1")
	require.Equal(t, "A", decode1)
}
