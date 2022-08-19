package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings = []string{}

	for {
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

		if userTickets <= remainingTickets {
			remainingTickets = conferenceTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Tank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("Only %v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)
		} else {
			fmt.Printf("Sorry! You cant reserve %d tickets, only %d remained\n", userTickets, remainingTickets)
			fmt.Println("Please try again")
		}
	}
}
