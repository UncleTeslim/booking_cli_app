package main

import (
	"fmt"
	"time"
)
	

const graduationTickets int = 1000
var eventName string = "UWE Graduation"
var remainingTickets uint = 1000
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	usertickets uint
}


func main() {

	greetUsers()

	for {
		firstName, lastName, email, usertickets := getUserInput()
		isNameValid, isEmailValid, isTicketAmountValid := validateUserInput(firstName, lastName, email, usertickets)

		if isNameValid && isEmailValid && isTicketAmountValid {
			
			bookTickets(usertickets, firstName, lastName, email)
			sendTicket(usertickets, firstName, lastName, email)

			// Print First Names
			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)


			var noMoreTickets bool = remainingTickets == 0
			if noMoreTickets  {
				fmt.Println("Graduation tickets are sold out, sorry!!")
				break
			}
		} else {
			if !isNameValid {
				fmt.Println("Name input is too short, please try again")
			}
			if !isEmailValid {
				fmt.Println("Please inout a valid email address")
			}
			if !isTicketAmountValid {
				fmt.Printf("There are %v tickets left, please choose an amount lesser than or equal to %v instead\n", remainingTickets, remainingTickets)
			}
		}
	}
		
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", eventName)
	fmt.Printf("There are total of %v tickets and %v are still available\n", graduationTickets, remainingTickets)
	fmt.Println("You can get you events tickets here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var usertickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&usertickets)

	return firstName, lastName, email, usertickets
}

func bookTickets(usertickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - usertickets

	//create map for user
	var user = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		usertickets: usertickets,
	}
	
	bookings = append(bookings, user)
	// index := len(bookings) -1
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, confirmation has been sent to %v.\n", firstName, lastName , usertickets, email)
	fmt.Printf("There are %v tickets available for %v\n", remainingTickets, eventName)

}

func sendTicket(usertickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v %v tickets for %v %v", usertickets, eventName, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n%v \nsent to email address %v\n",ticket, email)
	fmt.Println("###################")
}