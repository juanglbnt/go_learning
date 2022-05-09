package main

import (
	"fmt"
	"sync"
	"time"
)

const eventName = "big4concert"
const concertTickets int = 50

var RemainingTickets = 50
var bookings = []UserData{}

type UserData struct {
	FirstName    string
	LastName     string
	Email        string
	NumOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("Welcome to the %s book your ticket: %d tickets available\n", eventName, RemainingTickets)

	for {
		firstName, lastName, email, numOfTickets := getUserInput()
		isValidName, isValidEmail, isValidNumOfTickets := ValidateUserInput(firstName, lastName, email, numOfTickets)

		if isValidName && isValidEmail && isValidNumOfTickets {
			bookTicket(numOfTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(numOfTickets, firstName, lastName, email)
		} else {
			if !isValidName {
				fmt.Println("Your names are wrong, names must be 2 or more length characters")
			}
			if !isValidEmail {
				fmt.Println("Invalid email, please chek")
			}
			if !isValidNumOfTickets {
				fmt.Println("Not enough tickets, tickets remaining: ", RemainingTickets)
			}
		}

		if RemainingTickets == 0 {
			break
		}
	}
	wg.Wait()
	printBook()
}

func getUserInput() (string, string, string, int) {
	var fName, lName, email string
	var tickets int
	fmt.Println("Enter yout first name:")
	fmt.Scan(&fName)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("How many tickets?")
	fmt.Scan(&tickets)

	return fName, lName, email, tickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	RemainingTickets -= userTickets

	user := UserData{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		NumOfTickets: userTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thank you %s %s you have %d tickets now, you will receive soon a email with the tickets\ntickets remaining: %d\n", firstName, lastName, userTickets, RemainingTickets)
}

func printBook() {
	fmt.Printf("Booked tickets\n")
	for i, b := range bookings {
		fmt.Printf("%d. User name: %s %s User email: %s, Num. Tickets: %d\n", i+1, b.FirstName, b.LastName, b.Email, b.NumOfTickets)
	}
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
