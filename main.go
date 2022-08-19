package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < int(conferenceTickets) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		firstName, lastName, email, userTickets = getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidUserTickets && isValidEmail && isValidName {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				fmt.Println("Sorry, No tickets left!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you enterd is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you enterd does not contains @")
			}
			if !isValidUserTickets {
				fmt.Println("Number of tickets is not valid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickets
}

func getUserInput() (string, string, string, uint) {
	// ask user for their name and their number of tickets
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []string {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Tank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only %v tickets remaining for %v.\n", remainingTickets, conferenceName)

	return bookings
}
