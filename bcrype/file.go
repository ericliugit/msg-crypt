package bcrype

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

type Encrypt struct {
	key string
}

func Of(key string) *Encrypt {
	return &Encrypt{key: key}
}

func (obj *Encrypt) Encrypt(msg string) string {
	key, _ := hex.DecodeString(obj.key)
	plaintext := []byte(msg)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x\n", ciphertext)
}

func (obj *Encrypt) Decrypt(msg string) string {
	key, _ := hex.DecodeString(obj.key)
	ciphertext, _ := hex.DecodeString(msg)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}
