package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num1, num2 int
	var operator string

	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("you must type an integer:", err)
	}
	num2, err = strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("you must type an integer", err)
	}

	operator = os.Args[3]
	switch operator {
	case "+":
		fmt.Printf("result: %d\n", num1+num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Error division by zero")
		} else {
			fmt.Printf("result: %.2f\n", float64(float64(num1)/float64(num2)))
		}

	case "-":
		fmt.Printf("result: %d\n", num1-num2)

	case "echo":
		fmt.Printf("result: %d\n", num1*num2)

	default:
		fmt.Println("Error, you must select a valid operation")
	}

}
