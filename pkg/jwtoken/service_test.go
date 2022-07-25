package jwtoken

import (
	"encoding/json"
	"github.com/cristalhq/jwt/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var accessDuration = time.Hour * 1
var refreshDuration = time.Hour * 60
var secretKey = []byte("secret")

func TestNew(t *testing.T) {
	s, err := NewService(secretKey, accessDuration, refreshDuration)
	require.NoError(t, err)

	accessToken, err := s.CreateAccessToken("test", time.Now(), time.Now().Add(1*time.Minute))
	require.NoError(t, err)
	require.NotNil(t, accessToken)

	refreshToken, err := s.CreateRefreshToken("test", time.Now(), time.Now().Add(1*time.Minute))
	require.NoError(t, err)
	require.NotNil(t, refreshToken)

	verifier, err := jwt.NewVerifierHS(jwt.HS256, secretKey)
	require.NoError(t, err)

	tokenBytes := accessToken.Bytes()
	newToken, err := jwt.Parse(tokenBytes, verifier)
	require.NoError(t, err)

	err = verifier.Verify(newToken)
	require.NoError(t, err)

	var newClaims jwt.RegisteredClaims
	errClaims := json.Unmarshal(newToken.Claims(), &newClaims)
	require.NoError(t, errClaims)

	errParseClaims := jwt.ParseClaims(tokenBytes, verifier, &newClaims)
	require.NoError(t, errParseClaims)

	require.False(t, newClaims.IsForAudience("admin"))
	require.True(t, newClaims.IsForAudience(AccessAudience))

	require.True(t, newClaims.IsValidAt(time.Now()))
}
