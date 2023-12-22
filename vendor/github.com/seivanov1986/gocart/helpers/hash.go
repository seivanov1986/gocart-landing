package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	timeFormat = "20060102150405"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordHashVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateHashSession() string {
	t := time.Now()
	hashBytes := md5.Sum([]byte(t.Format(timeFormat)))
	return hex.EncodeToString(hashBytes[:])
}
