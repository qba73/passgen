// Package passgen provides functions for generating hashed passwords.
package passgen

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
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

// Password represents generated data: passwd, salt an hash.
type Password struct {
	Salt     string `json:"salt"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

// ToJSON formats output in json string.
func (p Password) ToJSON() string {
	output, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(output)
}

// RunCLI takes output and error writers and runs the program.
func RunCLI(w io.Writer, ew io.Writer) int {
	password := GeneratePassword()
	salt, err := GenerateSalt()
	if err != nil {
		fmt.Fprintln(ew, err)
		return 1
	}
	hashedPassword := HashPassword(password, salt)
	p := Password{
		Password: password,
		Hash:     hashedPassword,
		Salt:     salt,
	}
	fmt.Fprintf(w, "%s\n", p.ToJSON())
	return 0
}

// Main is the entry point to the program.
func Main() int {
	return RunCLI(os.Stdout, os.Stderr)
}
