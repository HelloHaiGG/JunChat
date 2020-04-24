package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

/**
aes 加解密
*/

//加密
func AesEncrypt(body []byte, key string) (string, error) {
	//key 的位数对应三种加密方式 16：aes-128,24：aes-192,32-aes256
	bKey := []byte(key)
	//分组密钥
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	body = PKCS7Padding(body, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, bKey[:blockSize])
	// 创建数组
	cipherBody := make([]byte, len(body))
	// 加密
	blockMode.CryptBlocks(cipherBody, body)
	return base64.StdEncoding.EncodeToString(cipherBody), nil
}

//解密
func AesDecrypt(cipherBody, key string) (string, error) {
	bCipher, err := base64.StdEncoding.DecodeString(cipherBody)
	if err != nil {
		return "", err
	}
	bKey := []byte(key)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	//与加密模式对应的解密模式
	blockMode := cipher.NewCBCDecrypter(block, bKey[:blockSize])
	//创建数组
	body := make([]byte, len(bCipher))
	//解密
	blockMode.CryptBlocks(body, bCipher)
	//去补码
	body = PKCS7UnPadding(body)

	return string(body), nil
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
