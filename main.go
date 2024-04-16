package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var bookings []string
		
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			
			remainingTickets = remainingTickets - userTickets 
			bookings = append(bookings, firstName + " " + lastName)
			
			fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining", remainingTickets)
	
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all of our bookings %v\n", firstNames)
	
			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)

		}
		

	}


}
