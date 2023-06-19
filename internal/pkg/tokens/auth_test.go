package tokens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTokens - тесты для функции токенов.
func TestTokens(t *testing.T) {
	userID := "test"

	token, err := CreateToken(userID)
	require.NoError(t, err)

	id, err := ParseToken(token)
	require.NoError(t, err)
	require.Equal(t, id, userID)
}
