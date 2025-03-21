package main

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"

	"github.com/goharbor/harbor/src/lib/log"
	"golang.org/x/crypto/pbkdf2"
)

const (
	// SHA1 is the name of sha1 hash alg
	SHA1 = "sha1"
	// SHA256 is the name of sha256 hash alg
	SHA256 = "sha256"
)

var (
	HashAlg = map[string]func() hash.Hash{
		SHA1:   sha1.New,
		SHA256: sha256.New,
	}

	password string
	salt     string
	alg      string
)

func main() {
	flag.StringVar(&password, "p", "", "password")
	flag.StringVar(&salt, "salt", "", "salt")
	flag.StringVar(&alg, "alg", SHA256, "encrypt algorithm, support sha1 and sha256")
	flag.Parse()

	if salt == "" {
		salt = GenerateRandomString()
	}
	fmt.Printf("password: %x", pbkdf2.Key([]byte(password), []byte(salt), 4096, 16, HashAlg[alg]))
	fmt.Println("\nsalt:", salt)
}

// GenerateRandomStringWithLen generates a random string with length
func GenerateRandomStringWithLen(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	l := len(chars)
	result := make([]byte, length)
	_, err := rand.Read(result)
	if err != nil {
		log.Warningf("Error reading random bytes: %v", err)
	}
	for i := 0; i < length; i++ {
		result[i] = chars[int(result[i])%l]
	}
	return string(result)
}

// GenerateRandomString generate a random string with 32 byte length
func GenerateRandomString() string {
	return GenerateRandomStringWithLen(32)
}
