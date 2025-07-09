package main

import (
	"fmt"
	"strings"
)
func main(){

	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint= 50
	bookings := [] string{}

 
	fmt.Printf(" conferenceName is %T, remainingTickets is %T, conferenceTicket is %T.\n", conferenceName, remainingTicket, conferenceTicket)

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here to attend.\n")

	for{
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		isValid := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

		if isValid && isValidEmail && isValidTicketNumber{
			remainingTicket = remainingTicket - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			noTicketsRemaining:= remainingTicket == 0
			if noTicketsRemaining {
				fmt.Printf("Our conference is booked out. Come back next year.\n")
				break
			}
		} else{
			if !isValid{
				fmt.Printf("First name or last name you entered is too short.\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address entered doesn't have @ sign. \n")
			}
			if !isValidTicketNumber{
				fmt.Printf("The number of tickets you entered is invalid\n")
			}
		}		
	}
	
}
