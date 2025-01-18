package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())          // Seed random number generator
    targetNumber := rand.Intn(100) + 1       // Random number between 1 and 100
    attempts := 0                             // Attempt counter
    reader := bufio.NewReader(os.Stdin)       // For reading user input

    fmt.Println("🎯 Welcome to the Number Guessing Game!")
    fmt.Println("🔢 I'm thinking of a number between 1 and 100.")
    fmt.Println("💡 Can you guess what it is?")

    for {
        fmt.Print("👉 Enter your guess: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("❗ Error reading input, try again.")
            continue
        }

        input = strings.TrimSpace(input)      // Remove spaces/newlines
        guess, err := strconv.Atoi(input)     // Convert input to integer
        if err != nil || guess < 1 || guess > 100 {
            fmt.Println("❗ Please enter a valid number between 1 and 100.")
            continue
        }

        attempts++

        if guess < targetNumber {
            fmt.Println("📉 Too low! Try again.")
        } else if guess > targetNumber {
            fmt.Println("📈 Too high! Try again.")
        } else {
            fmt.Printf("🎉 Congratulations! You guessed the number %d in %d attempts!\n", targetNumber, attempts)
            break
        }
    }
}
