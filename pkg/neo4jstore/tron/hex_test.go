package tron

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHex(t *testing.T) {
	b, err := DecodeAddress("TRtxgyKrUXmjVks9GFyabPDu6K5952HWXV")
	require.NoError(t, err)
	require.Equal(t, "41aeb466da22b8d280d7c57479c49ff061df311c74", b)
	a, err := EncodeAddress("41aeb466da22b8d280d7c57479c49ff061df311c74")
	require.NoError(t, err)
	require.Equal(t, "TRtxgyKrUXmjVks9GFyabPDu6K5952HWXV", a)
	//c, err := EncodeAddress("41203ebaa47ea3bdda77c1fa07bd542257d7dc35b2")
	//require.NoError(t, err)
	//require.Equal(t, "", c)
}
