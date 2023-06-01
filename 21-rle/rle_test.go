package p21rle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRle(t *testing.T) {
	encode := RLEEncode("AABAABAAABA")
	require.Equal(t, "A2B1A2B1A3B1A1", encode)

	decode := RLEDecode("A2B1A2B1A3B1A1")
	require.Equal(t, "AABAABAAABA", decode)

	encode1 := RLEEncode("A")
	require.Equal(t, "A1", encode1)

	decode1 := RLEDecode("A1")
	require.Equal(t, "A", decode1)
}
