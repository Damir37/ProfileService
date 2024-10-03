package unit

import (
	"ProfileService/internal/pkg/streamkey"
	"testing"
)

func TestGenerateStreamKey(t *testing.T) {
	key, err := streamkey.GenerateStreamKey()
	if err != nil {
		t.Fatalf("Ожидалось отсутствие ошибки, но получена ошибка: %v", err)
	}

	expectedLength := 64
	if len(key) != expectedLength {
		t.Errorf("Ожидалась длина ключа %d, но получена длина %d", expectedLength, len(key))
	}

	if key == "" {
		t.Errorf("Ожидалась непустая строка, но получена пустая строка")
	}
}
