package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var booking = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const conferenceTickets int = 50

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	//for {
	firstName, lastName, email, userTicket := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateInputValues(firstName, lastName, email, userTicket)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookingLogic(firstName, lastName, email, userTicket)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)
		printFirstName()

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.....")
		}
	} else {
		fmt.Printf("Your input value is in valid. Try Again")

	}
	wg.Wait()

	// }

}

func greetUser() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstName() {
	firstNames := []string{}
	for _, book := range booking {
		firstNames = append(firstNames, book.firstName)
	}
	fmt.Printf("The first name of booking are %v \n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTicket)
	return firstName, lastName, email, userTicket
}

func bookingLogic(firstName string, lastName string, email string, userTicket uint) {

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	booking = append(booking, userData)
	remainingTickets -= userTicket
	fmt.Printf("user %v %v booked %v tickets. and sent confirmation email to %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v to email address %v \n", tickets, email)
	fmt.Println("#################")
	wg.Done()
}
