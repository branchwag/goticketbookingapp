package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for name
	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
