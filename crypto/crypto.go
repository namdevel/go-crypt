package crypto

import (
    "encoding/base64"
    "encoding/hex"
    "compress/gzip"
	"bytes"
    "crypto/aes"
    "crypto/cipher"
    "io/ioutil"
    "fmt"
)

func EncB64(plaintext, key []byte) (string, error) {
    ciphertext, err := AesEncrypt(plaintext, key)
    if err != nil {
        return "", err
    }

    encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

    return encodedCiphertext, nil
}

func DecB64(encodedCiphertext string, key []byte) ([]byte, error) {
    ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
    if err != nil {
        return nil, err
    }

    plaintext, err := AesDecrypt(ciphertext, key)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

func EncGzip(plaintext, key []byte) ([]byte, error) {
    var b bytes.Buffer
    w := gzip.NewWriter(&b)
    _, err := w.Write(plaintext)
    if err != nil {
        return nil, err
    }
    w.Close()
    compressedPlaintext := b.Bytes()

    ciphertext, err := AesEncrypt(compressedPlaintext, key)
    if err != nil {
        return nil, err
    }

    return ciphertext, nil
}

func DecGzip(ciphertext, key []byte) ([]byte, error) {
    compressedPlaintext, err := AesDecrypt(ciphertext, key)
    if err != nil {
        return nil, err
    }

    b := bytes.NewBuffer(compressedPlaintext)
    r, err := gzip.NewReader(b)
    if err != nil {
        return nil, err
    }
    plaintext, err := ioutil.ReadAll(r)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

func EncHex(plaintext, key []byte) (string, error) {
    ciphertext, err := AesEncrypt(plaintext, key)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(ciphertext), nil
}

func DecHex(ciphertext string, key []byte) ([]byte, error) {
    decodedCiphertext, err := hex.DecodeString(ciphertext)
    if err != nil {
        return nil, err
    }

    plaintext, err := AesDecrypt(decodedCiphertext, key)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

func AesEncrypt(plaintext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    plaintext = pad(plaintext)
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func AesDecrypt(ciphertext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    return unpad(plaintext), nil
}

func pad(plaintext []byte) []byte {
    padding := aes.BlockSize - len(plaintext)%aes.BlockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(plaintext, padtext...)
}

func unpad(plaintext []byte) []byte {
    padding := plaintext[len(plaintext)-1]
    return plaintext[:len(plaintext)-int(padding)]
}