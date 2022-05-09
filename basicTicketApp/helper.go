package main

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, ticketsRequired int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isvalidEmail := len(email) > 4 && strings.Contains(email, "@")
	isValidTicketsRequired := ticketsRequired <= RemainingTickets

	return isValidName, isvalidEmail, isValidTicketsRequired
}
