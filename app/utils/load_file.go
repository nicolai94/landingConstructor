package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

func GenerateRandomFileName() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func CheckOrCreateDirectory(path string) *string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			errorMessage := fmt.Sprintf("Ошибка создания директории: %s", err.Error())
			return &errorMessage
		}
	}
	return nil
}
