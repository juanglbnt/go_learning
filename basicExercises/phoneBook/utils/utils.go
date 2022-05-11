package utils

import (
	"fmt"
)

func ValidateNumber() int {
	var num int
	for {
		fmt.Println("Enter the number (or -1 to exit)")
		fmt.Scan(&num)
		if (num > 999999 && num < 9999999) || num == -1 {
			return num
		} else {
			fmt.Println("Plese enter a valid number (7 digits)")
		}
	}
}

func GetNumbers() []int {
	var nums []int
	var num int
	for {
		num = ValidateNumber()
		if num == -1 {
			break
		}
		nums = append(nums, num)
	}
	return nums
}
