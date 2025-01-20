package main

import (
	"bufio"
	"file-encryption/encryption"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <encrypt|decrypt> <input_file> <output_file>")
		return
	}

	mode := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encryption key: ")
	key, _ := reader.ReadString('\n')
	key = key[:32] // Use the first 32 characters for AES-256

	switch mode {
	case "encrypt":
		err := encryption.EncryptFile(inputFile, outputFile, key)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully!")
	case "decrypt":
		err := encryption.DecryptFile(inputFile, outputFile, key)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully!")
	default:
		fmt.Println("Invalid mode. Use 'encrypt' or 'decrypt'.")
	}
}
