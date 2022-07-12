package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for _, arg := range os.Args[1:] {
		s += arg + " "
	}
	fmt.Println(s)
}
