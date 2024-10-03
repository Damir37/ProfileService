package streamkey

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateStreamKey() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("не удалось сгенерировать ключ: %v", err)
	}

	return hex.EncodeToString(bytes), nil
}
