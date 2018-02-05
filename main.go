package main

import (
    "fmt"
    "io"
    "log"
    "crypto/rand"
    "encoding/base64"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"

    "github.com/rs/xid"
)

// Generate random password
func generatePassword() string {
    guid := xid.New()
    return guid.String()
}

// Generate salt string with 32 bytes of crypto/rand data
func generateSalt() string {
    randomBytes := make([]byte, 32)
    _, err := rand.Read(randomBytes)
    if err != nil {
        log.Panic(err)
    }
    return base64.URLEncoding.EncodeToString(randomBytes)
}

// Hash password with the salt
func hashPassword(plainText string, salt string) string {
    hash := hmac.New(sha256.New, []byte(plainText))
    io.WriteString(hash, plainText + salt)
    hashedValue := hash.Sum(nil)
    return hex.EncodeToString(hashedValue)
}

// Print out results
func printResult(salt string, password string, hashedPass string) {
    fmt.Println("Salt: ")
    fmt.Println(salt)
    fmt.Println("Generated password: ")
    fmt.Println(password)
    fmt.Println("Hashed password: ")
    fmt.Println(hashedPass)
}


func main() {
    // new random string that will be used as temp roop password
    password := generatePassword()
    salt := generateSalt()
    hashedPassword := hashPassword(password, salt)
    printResult(salt, password, hashedPassword)
}
