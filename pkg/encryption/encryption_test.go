package encryption

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSimple - тесты для Encode и Decode.
func TestSimple(t *testing.T) {
	data := "test"
	enc, err := Encrypt(data)
	require.NoError(t, err)

	dec, err := Decode(enc)
	require.NoError(t, err)
	require.Equal(t, data, dec)
}

// TestBinary - тесты для EncodeBin и DecodeBin.
func TestBinary(t *testing.T) {
	data := []byte("test")
	enc, err := EncryptBin(data)
	require.NoError(t, err)

	dec, err := DecodeBin(enc)
	require.NoError(t, err)
	require.Equal(t, data, dec)
}
