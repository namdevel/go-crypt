package main

import (
	"fmt"
	"github.com/namdevel/go-crypt/crypto"
)

func main() {
    // Example plaintext
    plaintext := []byte("example plaintext")

    // Key for encryption and decryption
    key := []byte("example key 1234")

    // Encrypt plaintext to AES then encode to Hex
    EncB64, err := crypto.EncB64(plaintext, key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("1. - AESBase64 Encrypted:%s\n", EncB64)

    // Decrypt Base64 Cipher to AES then decode to plaintext
    DecB64, err := crypto.DecB64(EncB64, key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("   - AESBase64 Decrypted : %s\n\n", DecB64)

	// Encrypt plaintext to AES then encode to Gzip Compressed
    EncGzip, err := crypto.EncGzip(plaintext, key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("2. - AESGzip Encrypted:%s\n", EncGzip)

    // Decrypt Gzip Compressed Cipher to AES then decode to plaintext
    DecGzip, err := crypto.DecGzip(EncGzip, key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("   - AESGzip Decrypted : %s\n\n", DecGzip)

	// Encrypt plaintext to AES then encode to Hex
	EncHex, err := crypto.EncHex(plaintext, key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("3. - AESHex Encrypted:%s\n", EncHex)

    // Decrypt Hex Cipher to AES then decode to plaintext
    DecHex, err := crypto.DecHex(EncHex, key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("   - AESHex Decrypted : %s\n\n", DecHex)

	
}