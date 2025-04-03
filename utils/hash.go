package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	hash := append(salt, hashedPassword...)

	return base64.StdEncoding.EncodeToString(hash), nil
}

func VerifyPassword(storedHash, password string) error {
	hashBytes, err := base64.StdEncoding.DecodeString(storedHash)
	if err != nil {
		return err
	}

	salt := hashBytes[:16]

	storedHashedPassword := hashBytes[16:]

	fmt.Printf("Salt: %x\n", salt)
	fmt.Printf("Stored Hashed Password: %x\n", storedHashedPassword)

	computedHashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	fmt.Printf("Computed Hashed Password: %x\n", computedHashedPassword)

	if !equal(computedHashedPassword, storedHashedPassword) {
		return errors.New("password mismatch")
	}

	return nil
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	result := 0
	for i := 0; i < len(a); i++ {
		result |= int(a[i] ^ b[i])
	}

	return result == 0
}
