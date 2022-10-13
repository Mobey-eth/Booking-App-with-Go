package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Web3 Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50
	var bookings []string
	// bookings := []string{}

	//fmt.Printf("The confrencename is %T, confrence tickets is %T and the remaining tickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// Obtaining user data
		var firstName string
		var lastName string
		var email string
		var userTickets int
		fmt.Println("Gday sire, enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Input your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets would you like to purchase ?: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= (remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - (userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v %v for booking %v tickets today, You will recieve a confirmation mail at %v.\n\nHave a nice trip Sire.\n",
				firstName, lastName, userTickets, email)
			fmt.Printf("*%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var frstName = names[0]
				firstNames = append(firstNames, frstName)
			}
			fmt.Printf("these are all our bookings so far: %v\n\n", bookings)
			fmt.Printf("the firstnames of bookers are: %v\n\n", firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println(" Our Web3 confrence is fully booked out comeback next year! ")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First name or last name you entered is invalid\n")
			}
			if !isValidEmail {
				fmt.Printf("E-mail address format you entered is invalid\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid\n")
			}
		}

	}

}
