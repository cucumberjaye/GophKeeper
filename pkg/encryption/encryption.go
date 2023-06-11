package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

func randBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Encrypt(data string) (string, error) {
	key, err := randBytes(16)
	if err != nil {
		return "", err
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	enc := aesgcm.Seal(nil, key[:12], []byte(data), nil)
	enc = append(key, enc...)

	return hex.EncodeToString(enc), nil
}

func EncryptBin(data []byte) ([]byte, error) {
	key, err := randBytes(16)
	if err != nil {
		return nil, err
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	enc := aesgcm.Seal(nil, key[:12], data, nil)
	enc = append(key, enc...)

	return enc, nil
}

func Decode(encData string) (string, error) {
	data, err := hex.DecodeString(encData)
	if err != nil {
		return "", err
	}

	key := data[:16]

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	dec, err := aesgcm.Open(nil, key[:12], data[16:], nil)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}

func DecodeBin(data []byte) ([]byte, error) {
	key := data[:16]

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	dec, err := aesgcm.Open(nil, key[:12], data[16:], nil)
	if err != nil {
		return nil, err
	}

	return dec, nil
}
