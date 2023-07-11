package main

import "fmt"

func main() {
	var conferenceName = "Armadillo Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string

	userName = "Tom"
	fmt.Println(userName)
}
