package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"
)

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GenerateFilename sesuai format custom
func GenerateFilename(userID int, original string) string {
	timestamp := time.Now().Format("20060102150405")
	ext := filepath.Ext(original)
	base := timestamp

	// 5 huruf random
	base += randomString(5)

	// userID + random char biar panjang 4
	uidStr := strconv.Itoa(userID)
	for len(uidStr) < 4 {
		uidStr += randomString(1)
	}
	base += uidStr

	// 5 huruf random lagi
	base += randomString(5)

	return base + ext
}

// GeneratePresigned membuat token HMAC
func GeneratePresigned(secret, folder, filename, action string, expiry int64) string {
	message := fmt.Sprintf("%s/%s|%s|%d", folder, filename, action, expiry)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
