package main

import (
	"encoding/base64"
	"fmt"
	"crypto/aes"
	"os"
	"crypto/cipher"
)

// base64

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func base64Crypto() {
	// encode
	hello := "Hello base64!"
	debyte := base64Encode([]byte(hello))
	fmt.Println("base64 encode:", debyte)

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("base64 decode:", string(enbyte))
}

// AES

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
func aesCrypto() {
	//需要加密的字符串
	plaintext := []byte("My name is Frank.")

	// AES密钥串
	key_text := "asdfghjklasdfghjklasdfghjklasdfg"	//32bytes
	fmt.Println("密钥长度：", len(key_text))

	//创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error:NewCipher (%d bytes = %s", len(key_text), err)
		os.Exit(1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s->%x\n", plaintext, ciphertext)

	//解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x->%s\n", ciphertext, plaintextCopy)
}


func main() {
	//base64 test
	base64Crypto()

	//aes test
	aesCrypto()
}