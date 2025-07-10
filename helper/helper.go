package helper

import (
	"regexp"
	"strings"
)

// ValidateUserInputs validates user booking information
func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// Trim whitespace for validation
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	email = strings.TrimSpace(email)
	
	// Name validation: at least 2 characters, letters only
	isValidName := IsValidName(firstName) && IsValidName(lastName)
	
	// Email validation: proper email format
	isValidEmail := IsValidEmail(email)
	
	// Ticket validation: positive number within available range
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	
	return isValidName, isValidEmail, isValidTicketNumber
}

// IsValidName checks if a name is valid (at least 2 chars, letters and spaces only)
func IsValidName(name string) bool {
	if len(name) < 2 {
		return false
	}
	
	// Check if name contains only letters and spaces
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == ' ') {
			return false
		}
	}
	
	return true
}

// IsValidEmail validates email format using regex
func IsValidEmail(email string) bool {
	if len(email) < 5 { // minimum: a@b.c
		return false
	}
	
	// Basic email regex pattern
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	
	return matched
}

// IsValidTicketCount checks if the requested ticket count is valid
func IsValidTicketCount(userTickets uint, remainingTickets uint) bool {
	return userTickets > 0 && userTickets <= remainingTickets
}

// ValidateUserInputsDetailed provides detailed validation with error messages
func ValidateUserInputsDetailed(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool, []string) {
	var errors []string
	
	// Trim whitespace
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	email = strings.TrimSpace(email)
	
	// Name validation
	isValidName := true
	if !IsValidName(firstName) {
		isValidName = false
		errors = append(errors, "First name must be at least 2 characters and contain only letters")
	}
	if !IsValidName(lastName) {
		isValidName = false
		errors = append(errors, "Last name must be at least 2 characters and contain only letters")
	}
	
	// Email validation
	isValidEmail := IsValidEmail(email)
	if !isValidEmail {
		errors = append(errors, "Please enter a valid email address (e.g., user@example.com)")
	}
	
	// Ticket validation
	isValidTicketNumber := IsValidTicketCount(userTickets, remainingTickets)
	if !isValidTicketNumber {
		if userTickets == 0 {
			errors = append(errors, "Number of tickets must be greater than 0")
		} else if userTickets > remainingTickets {
			errors = append(errors, "Not enough tickets available")
		}
	}
	
	return isValidName, isValidEmail, isValidTicketNumber, errors
}