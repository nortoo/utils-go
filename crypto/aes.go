package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// AESCBCEncryptor AES CBC encryptor
func AESCBCEncryptor(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	iv := key[len(key)-blockSize:] // 取密钥的后十六个字节

	origData = pkcs7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encryptedData := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptedData, origData)
	return encryptedData, nil
}

// AESCBCDecrypter AES CBC Decrypter
func AESCBCDecrypter(encryptedData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	iv := key[len(key)-blockSize:] // 取密钥的后十六个字节

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encryptedData))
	blockMode.CryptBlocks(origData, encryptedData)
	return pkcs7UnPadding(origData), nil
}

// AESGCMEncryptor AES GCM encryptor
func AESGCMEncryptor(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := key[len(key)-12:]

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nil, nonce, data, nil), nil
}

// AESGCMDecrypter AES GCM decrypter
func AESGCMDecrypter(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := key[len(key)-12:]
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm.Open(nil, nonce, data, nil)
}
