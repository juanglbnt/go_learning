package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeToFile() {
	file, err := os.OpenFile("./save.txt", os.O_WRONLY|os.O_APPEND, 0644)
	reader := bufio.NewReader(os.Stdin)
	var name string

	for {
		fmt.Println("Enter your name:")
		n, _ := reader.ReadString('\n')
		name = strings.TrimRight(n, "\r\n")
		if name == "" {
			break
		}

		_, err = file.WriteString("\n" + name)
		if err != nil {
			fmt.Println("Error writing name:", err)
		}

	}
}

func readDataFromFile() {
	file, err := os.Open("./save.txt")

	if err != nil {
		fmt.Println("Error opening file data:", err)
	}
	defer file.Close()

	var lines []string
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	fmt.Println("names:", lines)
}

func main() {
	readDataFromFile()
	writeToFile()
	readDataFromFile()
}
