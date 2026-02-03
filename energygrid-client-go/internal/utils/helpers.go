package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// Generate MD5(url + token + timestamp)
func GenerateSignature(url, token, timestamp string) string {
	hash := md5.Sum([]byte(url + token + timestamp))
	return hex.EncodeToString(hash[:])
}

// Sleep helper
func Sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// Generate SN-000 to SN-499
func GenerateSerialNumbers() []string {
	serials := make([]string, 500)
	for i := 0; i < 500; i++ {
		serials[i] = fmt.Sprintf("SN-%03d", i)
	}
	return serials
}
