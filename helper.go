package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, usertickets uint) (bool, bool, bool){
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2 
	isEmailValid := strings.Contains(email, "@")
	isTicketAmountValid := usertickets> 0 && usertickets <= remainingTickets
	return isNameValid, isEmailValid, isTicketAmountValid
}