package main

import (
	"fmt"
	"os"
	"strings"
)

func normalWay() {
	file, err := os.Create("hexList.txt")
	if err != nil {
		fmt.Println("Error creating file[", err, "]")
	}

	final := 150000
	for i := 0; i <= final; i++ {
		_, err := file.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
}

func hexaDecimalRoutines(start, end int, result chan string, done chan struct{}) {
	var sBuilder strings.Builder
	for i := start; i <= end; i++ {
		fmt.Fprint(&sBuilder, fmt.Sprintf("%06x\n", i))
	}
	result <- sBuilder.String()
	done <- struct{}{}
}

func main() {
	//normalWay() //30 sec
	file, err := os.Create("hexList.txt")
	if err != nil {
		fmt.Println("Error creating file[", err, "]")
	}

	outChan := make(chan string)
	doneWrite := make(chan struct{})
	go func() {
		for s := range outChan {
			_, err := file.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{}
	}()

	numGoRoutines := 10
	done := make(chan struct{})

	var final = 150000
	for i := 0; i <= final; i = i + (final / numGoRoutines) + 1 {
		step := i + (final / numGoRoutines)
		if step > final {
			step = final
		}
		go hexaDecimalRoutines(i, step, outChan, done)
	}

	doneNum := 0
	for doneNum < numGoRoutines {
		<-done
		doneNum++
	}
	close(outChan)
	<-doneWrite
	fmt.Println("FINISHED")
}
