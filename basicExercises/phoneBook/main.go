package main

import (
	"fmt"

	"com.juan.phoneBook/book"
)

func main() {
	fmt.Printf("## Phone book\n")
	var opc int
	for opc != -1 {
		fmt.Printf("1. Register new\n2. See all data\n3. Edit data\n")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			book.AddItem()
		case 2:
			book.PrintBook()
		case 3:
			book.EditPhoneList()
		default:

		}
	}
}
