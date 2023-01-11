# go-crypt

```
go get github.com/namdevel/go-crypt/crypto
```

### Key & String
```go
  plaintext := []byte("example plaintext")
  // Key for encryption and decryption
  key := []byte("example key 1234")
```

### AES Base64
```go
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
```

### AES Gzip Compress
```go
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
```

### AES Hex
```go
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
```

### Example
```
1. - AESBase64 Encrypted:AAAAAAAAAAAAAAAA6rZStyscF5jffw8oeBuVgeop9iPEx9z3tRbF5ZHf1Q4eXg5ZGZV1fjsufRkrrKDE
   - AESBase64 Decrypted : example plaintext

2. - AESGzip Encrypted:�E;�[pr���$�^��1�s���������F�� �����"�}b��������z����M
   - AESGzip Decrypted : example plaintext

3. - AESHex Encrypted:000000000000000000000000eab652b72b1c1798df7f0f28781b9581ea29f623c4c7dcf7b516c5e591dfd50e1e5e0e591995757e3b2e7d192baca0c4
   - AESHex Decrypted : example plaintext
```

License
------------

This open-source software is distributed under the MIT License. See LICENSE.md

Contributing
------------

All kinds of contributions are welcome - code, tests, documentation, bug reports, new features, etc...

* Send feedbacks.
* Submit bug reports.
* Write/Edit the documents.
* Fix bugs or add new features.
