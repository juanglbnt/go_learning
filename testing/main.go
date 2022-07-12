package main

import (
    "fmt"
    "testing/function"
)

var sq = function.Square {
    X: 10.0,
    Y: 75.6,
}

func main() {
    fmt.Println("hello from Go")
    area := sq.Area()
    perim := sq.Perim()
    fmt.Printf("\nSquare: [%v %v]\nArea: %v\nPerim: %v\n", sq.X, sq.Y, area, perim)

}
