package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserInputs()
	isValidUserName, isValidEmail, isValidTicketNumber := helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidUserName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

		if remainingTickets == 0 {
			fmt.Printf("\nThe %v is booked up. Come back next year\n", conferenceName)
			fmt.Println("————————————————————————————————————————")
			fmt.Println()
		} else {
			fmt.Println("————————————————————————————————————————")
			fmt.Println()
		}
	} else {
		fmt.Println("————————————————————————————————————————")
		if !isValidUserName {
			fmt.Println("The first name or last name you entered in too short")
		}
		if !isValidEmail {
			fmt.Println("Email must contain the '@' and '.' signs")
		}
		if !isValidTicketNumber {
			fmt.Printf("Number of tickets should be a positive number that's less or equal to %v\n", remainingTickets)
		}
		fmt.Println("————————————————————————————————————————")
		fmt.Println()
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("————————————————————————————————————————")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Println("————————————————————————————————————————")
	fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("\n%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println()
	wg.Done()
}
