// Package passgen provides functions for generating hashed passwords.
package passgen

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/rs/xid"
)

// GeneratePassword generates random password.
func GeneratePassword() string {
	guid := xid.New()
	return guid.String()
}

// GenerateSalt generates a random salt with 32 bytes of crypto/rand data.
func GenerateSalt() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// HashPassword takes passwd string and salt and returns
// a string representing hashed password.
func HashPassword(passwd string, salt string) string {
	hash := hmac.New(sha256.New, []byte(passwd))
	io.WriteString(hash, passwd+salt)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

func runCLI(w io.Writer, ew io.Writer) int {
	password := GeneratePassword()
	salt, err := GenerateSalt()
	if err != nil {
		fmt.Fprintln(ew, err)
		return 1
	}
	hashedPassword := HashPassword(password, salt)
	fmt.Fprintf(w, "Salt: %s\nPasswd: %s\nHashed passwd: %s\n", salt, password, hashedPassword)
	return 0
}

func Main() int {
	return runCLI(os.Stdout, os.Stderr)
}
