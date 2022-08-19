package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings = []string{}

	greetUsers(conferenceName, int(conferenceTickets), int(remainingTickets))

	for remainingTickets > 0 && len(bookings) < int(conferenceTickets) {
		noTicketRemaining := remainingTickets == 0
		if noTicketRemaining {
			fmt.Println("Sorry, No tickets left!")
			break
		}

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidUserTickets && isValidEmail && isValidName {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Tank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("Only %v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
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

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets int) {
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
