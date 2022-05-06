package main

import (
	"fmt"
)

func computeStatistics(grades []float64) {
	min, max := computeMinMax(grades)
	avg := computeAverage(grades)

	fmt.Printf("The average is: %.2f\nThe Maximum grade is: %.2f\nThe minimum grade is: %.2f\n", avg, max, min)
}

func computeAverage(grades []float64) float64 {
	var sum, avg float64

	for _, g := range grades {
		sum += g
	}
	avg = sum / float64(len(grades))

	return avg
}

func computeMinMax(grades []float64) (float64, float64) {
	min, max := grades[0], grades[0]

	for _, g := range grades {

		if g > max {
			max = g
		}

		if g < min {
			min = g
		}
	}

	return min, max
}

func main() {
	fmt.Printf("\n## Grades Statistics\n")
	count := 1
	var grade float64
	var grades []float64

	for grade != -1.0 {
		fmt.Println("Enter the grade of student:", count, "(or -1 to exit)")
		fmt.Scanf("%f", &grade)
		count++
		if grade != -1 {
			grades = append(grades, grade)
		}
	}

	computeStatistics(grades)
}
