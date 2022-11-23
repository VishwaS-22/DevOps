package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	fName           string
	lName           string
	email           string
	numberOfTickets uint
}

var wG = sync.WaitGroup{} // it waits untill all goroutines are executed succussfully.

func main() {

	greetings()
	//ask user for the input.

	firstName, lastName, email, userTickets := UserInput()

	isValidFirstName, isValidLastName, isValidEmail, isValidTicketNumber := validUserInput(firstName, lastName, email, userTickets, remainingTickets)
	if isValidFirstName && isValidLastName && isValidEmail && isValidTicketNumber {

		bookingTicket(userTickets, firstName, lastName, email)
		wG.Add(1)
		go SendingTicket(userTickets, email, firstName, lastName)

		Names := first_Names()
		fmt.Printf("Tickets booked person's list : %v\n", Names)

		if remainingTickets == 0 {
			fmt.Printf("Our tickets for the %v has been sold out. Please come back next year.", conferenceName)
		}
	} else {
		if !isValidFirstName {
			fmt.Println("The firstName you entered is too short. Please enter more than or equal to two chars.")
		}
		if !isValidLastName {
			fmt.Println("The lastName you entered is too short. Please enter more than or equal to two chars.")
		}
		if !isValidEmail {
			fmt.Println("The Email you entered is wrong. Please enter a valid one.")
		}
		if !isValidTicketNumber {
			fmt.Println("Please enter a Valid Number.")
		}
	}
	wG.Wait()

}

func greetings() {
	fmt.Println("Welcome Message!")
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v Booking Application.\n", conferenceName)
	fmt.Printf("We've total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can get your", conferenceName, "tickets here.")
}

func first_Names() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.fName)
	}
	return firstNames
}

func UserInput() (string, string, string, uint) {
	fmt.Println("User Details:")
	var fName string
	var lName string
	var Email string
	var userTicks uint

	fmt.Println("Enter your FirstName:")
	fmt.Scan(&fName)

	fmt.Println("Enter your LastName:")
	fmt.Scan(&lName)

	fmt.Println("Enter your Email Address:")
	fmt.Scan(&Email)

	fmt.Println("Enter the number of tickets that you want:")
	fmt.Scan(&userTicks)

	return fName, lName, Email, userTicks
}

func bookingTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData1 = userData{
		fName:           firstName,
		lName:           lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData1)
	fmt.Printf("List of bookings is %v\n", bookings)
	
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation mail at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)
}
func SendingTicket(uTicks uint, email string, Fname string, Lname string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", uTicks, Fname, Lname)
	fmt.Println("* * * * * * * * *")
	fmt.Printf("Sending %v to given email address %v \n", tickets, email)
	fmt.Println("* * * * * * * * *")
	wG.Done()
}
