package main

import "strings"

func validUserInput(first string, last string, mail string, userTick uint) (bool, bool, bool, bool) {
	isValidFirstName1 := len(first) >= 2
	isValidLastName1 := len(last) >= 2
	isValidEmail1 := strings.Contains(mail, "@") && strings.Contains(mail, ".com")
	isValidTicketNumber1 := userTick > 0 && userTick <= remainingTickets

	return isValidFirstName1, isValidLastName1, isValidEmail1, isValidTicketNumber1
}
