package password

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	pass := []byte("user@demo.com")

	hash, err := HashPassword(pass)
	if err != nil {
		t.Error(err)
	}

	require.True(t, CheckPasswordHash(pass, hash))
}
