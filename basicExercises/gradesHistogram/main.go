package main

import (
	"fmt"
)

func main() {
	fmt.Printf("## Grades Histogram\n")
	var numStudents, grade int
	var histo [10]int
	var grades []int

	fmt.Println("Enter the number of students:")
	fmt.Scanf("%d", &numStudents)

	for i := 0; i < numStudents; i++ {
		fmt.Println("Enter the grade of student:", i)
		fmt.Scanf("%d", &grade)
		grades = append(grades, grade)
	}

	for _, g := range grades {

		if g == 100 {
			histo[9]++
		} else {
			histo[g/10]++
		}
	}

	for i, h := range histo {

		if h != 9 {
			fmt.Printf("%2d - %3d: %d\n", i*10, i*10+9, h)
		} else {
			fmt.Printf("%2d - %3d: %d\n", i*10, i*10+10, h)
		}
	}
}
