package main

import "fmt"

func main() {
	var conferenceName = "Armadillo Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T. \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// USER INPUT
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your first email")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

}
