package main

import "fmt"

func main() {
	var conferenceName = "Armadillo Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application!")
	fmt.Println("We have a total of", conferenceTickets, "and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

}
