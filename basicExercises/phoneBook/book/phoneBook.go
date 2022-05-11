package book

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"com.juan.phoneBook/utils"
)

var book = map[string][]int{}

func AddItem() {
	reader := bufio.NewReader(os.Stdin)
	var key string
	var numbers []int
	fmt.Println("Enter the key")
	k, _ := reader.ReadString('\n')
	key = strings.TrimRight(k, "\r\n")
	numbers = utils.GetNumbers()

	if len(numbers) == 0 {
		fmt.Println("Error storing the data, phone numbers must be must than 0")
	} else {
		book[key] = numbers
	}
}

func PrintBook() {
	for k, v := range book {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func EditPhoneList() []int {
	var nums []int
	var newNums []int
	var opc int

	var element string
	reader := bufio.NewReader(os.Stdin)

	for !searchInMap(element) {
		fmt.Println("select one of the book")
		PrintBook()
		e, _ := reader.ReadString('\n')
		element = strings.TrimRight(e, "\r\n")
		fmt.Println("you must select one from the list")
	}

	nums = book[element]

	for opc != -1 {
		fmt.Printf("1. Delete number\n2. Add number\n")
		fmt.Scan(&opc)

		if opc == 1 {
			var numToDelete int
			fmt.Println("Select the number to delete")
			for i, n := range nums {
				fmt.Printf("%d. %d\n", i, n)
			}
			fmt.Scan(&numToDelete)
			if numToDelete >= 0 && numToDelete < len(nums) {
				newNums = remove(nums, numToDelete)
				book[element] = newNums
			} else {
				fmt.Println("Error, you must select an element from the list")
			}
		}

		if opc == 2 {
			numToAdd := utils.ValidateNumber()
			newNums = nums
			fmt.Println(newNums)
			newNums = append(newNums, numToAdd)
			book[element] = newNums
		}
	}

	return newNums
}

func searchInMap(e string) bool {
	isIn := false
	for k, _ := range book {
		if e == k {
			isIn = true
		}
	}

	return isIn
}

func remove(n []int, i int) []int {
	n[i] = n[len(n)-1]
	return n[:len(n)-1]
}
