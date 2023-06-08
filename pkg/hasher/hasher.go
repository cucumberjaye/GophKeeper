package hasher

import (
	"crypto/sha256"
	"fmt"
)

func HasherSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)

	return fmt.Sprintf("%x", hash)
}
