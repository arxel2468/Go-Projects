# File Encryption and Decryption Tool

A command-line tool to securely encrypt and decrypt files using AES-256 encryption.

## Features
- Encrypt files with a 256-bit key.
- Decrypt files encrypted by the tool.
- Easy-to-use command-line interface.

## Requirements
- Go 1.20 or higher.

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/arxel2468/Go-Projects.git
   cd Intermediate/file-encryption
   go mod tidy

## Usage
Encryption
   ```bash
   go run main.go encrypt <input_file> <output_file>

- Enter a 32-character encryption key when prompted.
For example: DoAc9Gc6uopORvzgfj3YlkLf/R0AF02Rdg2kTwUDxEd4SXXeL7eF5WPvBI+icyUP


Decryption
   ```bash
   go run main.go decrypt <input_file> <output_file>

Example
1. Encrypt a file:
   ```bash
   go run main.go encrypt test_files/sample.txt test_files/output.enc

2. Decrypt a file
   ```bash
   go run main.go decrypt test_files/output.enc test_files/decrypted.txt

3. Verify the output:
   ```bash
   diff test_files/sample.txt test_files/decrypted.txt

# Notes
Ensure the encryption key is securely stored.
Losing the key will render encrypted files unrecoverable.