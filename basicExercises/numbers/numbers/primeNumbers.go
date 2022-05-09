package numbers

import (
	"fmt"
)

func PrintPrimeNumbers(num int) {
	var primes = []int{}
	for i := 1; i <= num; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println("Prime numbers:", primes)
}

func isPrime(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum++
		}
	}
	return sum <= 1
}
