package numbers

import (
	"fmt"
)

var oddNumbers = []int{}
var evenNumbers = []int{}

func OddEvenNumbers(num int) {

	for i := 0; i <= num; i++ {
		if i%2 == 0 {
			evenNumbers = append(evenNumbers, i)
		} else {
			oddNumbers = append(oddNumbers, i)
		}
	}

	fmt.Printf("Even Numbers: %v\nOdd Numbers: %v\n", oddNumbers, evenNumbers)
}
