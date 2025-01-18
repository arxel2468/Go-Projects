package main

import (
    "fmt"
)

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) float64 {
    if b == 0 {
        fmt.Println("Error: Division by zero is not allowed.")
        return 0
    }
    return a / b
}

func main() {
    var num1, num2 float64
    var operator string

    fmt.Println("🧮 Simple Calculator")
    fmt.Println("--------------------")
    fmt.Println("Available operations: +  -  *  /")
    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)

    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)

    var result float64
    switch operator {
    case "+":
        result = add(num1, num2)
    case "-":
        result = subtract(num1, num2)
    case "*":
        result = multiply(num1, num2)
    case "/":
        result = divide(num1, num2)
    default:
        fmt.Println("Error: Invalid operator. Please use +, -, *, or /.")
        return
    }

    fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
