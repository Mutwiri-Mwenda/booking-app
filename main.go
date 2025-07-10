package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Global configuration
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()
	
	fmt.Printf("conferenceName is %T, remainingTickets is %T, conferenceTickets is %T.\n", 
		conferenceName, remainingTickets, conferenceTickets)
	
	// Main booking loop
	for {
		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(
			firstName, lastName, email, userTickets, remainingTickets)
		
		if isValidName && isValidEmail && isValidTicketNumber {
			// Process the booking
			bookTickets(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			
			// Display current bookings - FIXED: Extract first names properly
			firstNames := getFirstNamesFromBookings()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			// Check if conference is sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			displayValidationErrors(isValidName, isValidEmail, isValidTicketNumber)
		}
	}
}

// greetUsers displays welcome message and conference information
func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", 
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
	fmt.Println(strings.Repeat("=", 50))
}

// getUserInput collects user information for booking
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("\n📝 Please enter your booking details:")
	
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	
	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTickets)
	
	return firstName, lastName, email, userTickets
}

// displayValidationErrors shows specific error messages for invalid inputs
func displayValidationErrors(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
	fmt.Println("\n❌ Please fix the following errors:")
	
	if !isValidName {
		fmt.Println("   • First name or last name you entered is too short (minimum 2 characters).")
	}
	if !isValidEmail {
		fmt.Println("   • Email address entered is invalid. Please include @ and domain.")
	}
	if !isValidTicketNumber {
		fmt.Println("   • The number of tickets you entered is invalid.")
	}
	fmt.Println()
}

// bookTickets processes the ticket booking and updates global state
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)
	
	fmt.Printf("\n✅ Thank you %v %v for booking %v tickets. "+
		"You will receive a confirmation email at %v.\n", 
		firstName, lastName, userTickets, email)
	fmt.Printf("📊 %v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Println(strings.Repeat("-", 50))
}

// getFirstNamesFromBookings extracts first names from the bookings map slice
func getFirstNamesFromBookings() []string {
	firstNames := []string{}
	
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	
	return firstNames
}

// getFirstNames extracts first names from a string slice (legacy function)
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	
	for _, booking := range bookings {
		names := strings.Fields(booking)
		if len(names) > 0 {
			firstNames = append(firstNames, names[0])
		}
	}
	
	return firstNames
}

// displayBookingSummary shows a summary of all bookings (optional enhancement)
func displayBookingSummary() {
	if len(bookings) == 0 {
		fmt.Println("No bookings yet.")
		return
	}
	
	fmt.Printf("\n📋 Booking Summary for %v:\n", conferenceName)
	fmt.Printf("Total bookings: %v\n", len(bookings))
	fmt.Printf("Tickets sold: %v\n", conferenceTickets-remainingTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	
	fmt.Println("\n👥 Attendees:")
	for i, booking := range bookings {
		fmt.Printf("%v. %v %v (%v tickets)\n", i+1, 
			booking["firstName"], booking["lastName"], booking["numberOfTickets"])
	}
	fmt.Println()
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Printf("#########################\n")
	fmt.Printf("Sending %v tickets to email address %v \n", ticket, email)
	fmt.Printf("#########################\n")
}