package main

import (
	"fmt"
	"golang/helper"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets
var bookings = make([]userData, 1)

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < int(conferenceTickets) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		firstName, lastName, email, userTickets = getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidUserTickets && isValidEmail && isValidName {
			bookTicket(userTickets, firstName, lastName, email)

			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
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

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Tank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only %v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	ticket := fmt.Sprintf("%v tickets sent for %v %v", userTickets, firstName, lastName)
	fmt.Println("***********")
	fmt.Printf("Sending ticket: %v\n to email %s\n", ticket, email)
	fmt.Println("***********")
}
