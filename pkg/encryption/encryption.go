package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

const sigingKey = "test"

func Encrypt(data string) (string, error) {
	key := md5.Sum([]byte(sigingKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	enc := aesgcm.Seal(nil, key[:12], []byte(data), nil)

	return hex.EncodeToString(enc), nil
}

func EncryptBin(data []byte) (string, error) {
	key := md5.Sum([]byte(sigingKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	enc := aesgcm.Seal(nil, key[:12], data, nil)

	return hex.EncodeToString(enc), nil
}

func Decode(encData string) (string, error) {
	data, err := hex.DecodeString(encData)
	if err != nil {
		return "", err
	}

	key := md5.Sum([]byte(sigingKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	dec, err := aesgcm.Open(nil, key[:12], data, nil)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}

func DecodeBin(encData string) ([]byte, error) {
	data, err := hex.DecodeString(encData)
	if err != nil {
		return nil, err
	}

	key := md5.Sum([]byte(sigingKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	dec, err := aesgcm.Open(nil, key[:12], data, nil)
	if err != nil {
		return nil, err
	}

	return dec, nil
}
