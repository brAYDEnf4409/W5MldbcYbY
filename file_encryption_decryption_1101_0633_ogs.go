// 代码生成时间: 2025-11-01 06:33:19
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// encryptionKey must be 16, 24, or 32 bytes long
var encryptionKey = []byte("your_secret_key_here")

// EncryptFile encrypts the contents of a file and writes it to a new file
func EncryptFile(filePath string, outputFilePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Encrypt data
	encryptedData, err := AesEncrypt(data, encryptionKey)
	if err != nil {
		return fmt.Errorf("error encrypting data: %w", err)
	}

	// Write encrypted data to output file
	err = ioutil.WriteFile(outputFilePath, encryptedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing encrypted file: %w", err)
	}

	return nil
}

// DecryptFile decrypts the contents of a file and writes it to a new file
func DecryptFile(filePath string, outputFilePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Decrypt data
	decryptedData, err := AesDecrypt(data, encryptionKey)
	if err != nil {
		return fmt.Errorf("error decrypting data: %w", err)
	}

	// Write decrypted data to output file
	err = ioutil.WriteFile(outputFilePath, decryptedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing decrypted file: %w", err)
	}

	return nil
}

// AesEncrypt encrypts the given data using AES algorithm
func AesEncrypt(data []byte, key []byte) ([]byte, error)
{
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// AesDecrypt decrypts the given data using AES algorithm
func AesDecrypt(data []byte, key []byte) ([]byte, error)
{
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("data too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func main() {
	// Example usage of encryption and decryption
	err := EncryptFile("example.txt", "example_encrypted.txt")
	if err != nil {
		log.Fatalf("Encryption error: %v", err)
	}

	err = DecryptFile("example_encrypted.txt", "example_decrypted.txt")
	if err != nil {
		log.Fatalf("Decryption error: %v", err)
	}
}
