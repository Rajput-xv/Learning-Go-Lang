package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!") //new line print
	// fmt.Print("Hello, World!")   //same line print
	// fmt.Printf("some text with variable %s \n", oneVariable) //formatted print

	// fmt.Println("Welcome to the booking app")
	// fmt.Println("Get your tickets here to attend")

	// var conferenceName = "Go Conference"
	// fmt.Println(conferenceName)

	// var conferenceName = "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets = 50

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// conferenceName := "Go Conference" //shorthand way of declaring and initializing variable

	// fmt.Println("Welcome to the", conferenceName, "booking app")
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Println("Get your tickets here to attend")

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	// var userName string //specifying data type
	// var userTickets int //specifying data type

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) //data types

	// userName = "Sanket"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	// fmt.Printf("User name is %v\n", userName)

	var firstName string
	var lastName string
	var email string
	var userName string
	var userTickets uint

	// var bookings = [50]string{"Yash","Tom","Ansh"} //array declaration
	// var bookings = [50]string{} //array declaration
	// var bookings [50]string //array declaration

	// dynamic slice
	var bookings []string
	// var bookings = []string{}
	// bookings := []string{}

	// adding values to dynamic array
	// bookings = append(bookings, "Yash")
	// bookings = append(bookings, firstName + " " + lastName)

	// bookings[0] = "Yash"
	// bookings[1] = "Tom"
	// bookings[2] = "Ansh"
	// bookings[0] = firstName + " " + lastName

	//taking user input
	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName) //& is used to get the memory address of the variable

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your User name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array/slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The array/slice type: %T\n", bookings)
	fmt.Printf("The length of array/slice: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// fmt.Printf("All bookings: %v\n", bookings)

	// loops
	// for {
	// 	//logic
	// }

	firstNames := []string{}
	//always two values when ranging over a slice
	// for index, booking := range bookings {
	// 	var names = strings.Fields(booking)
	// 	firstNames = append(firstNames, names[0])
	// }

	// _ is used to ignore the value(unused variable)
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)

	// if remainingTickets == 0 {
	// 	//end program
	// 	fmt.Println("Our conference is booked out. Come back next year.")
	// 	// break
	// }

	var noTickets bool = remainingTickets == 0
	if noTickets {
		//end program
		fmt.Println("Our conference is booked out. Come back next year.")
		// break
	}

	// if userTickets > remainingTickets {
	// 	// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
	// } else {
	// 	//logic of booking the tickets
	// }

	// if userTickets < remainingTickets {
	// 	//logic of booking the tickets
	// 	fmt.Println("Booking tickets...")
	// } else if userTickets == remainingTickets {
	// 	//logic of booking the tickets
	// 	fmt.Println("You have booked the last tickets...")
	// } else {
	// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
	// }

	// for true{
	// 	//logic
	// }

	// for remainingTickets > 0 && len(bookings) < int(conferenceTickets) {
	// 	//Logic
	// }

	//Input Validations
	// isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// isValidEmail := strings.Contains(email, "@")
	// isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// if isValidName && isValidEmail && isValidTicketNumber {
	// 	//logic for booking the ticket
	// } else {
	// 	if !isValidName {
	// 		fmt.Println("First name or last name you entered is too short")
	// 	}
	// 	if !isValidEmail {
	// 		fmt.Println("Email you entered is not valid")
	// 	}
	// 	if !isValidTicketNumber {
	// 		fmt.Println("Number of tickets you entered is not valid")
	// 	}
	// }

	// switch statement
	// city := "London"

	// switch city {
	// 	case "New York":
	// 		fmt.Println("You are in New York")
	// 	case "Singapore", "Hong Kong":
	// 		fmt.Println("You are in Asia")
	// 	case "London", "Berlin":
	// 		fmt.Println("You are in Europe")
	// 	default:
	// 		fmt.Println("You are somewhere else")
	// }

	//fucntions
	// greetUsers()
	// greetUsers2(conferenceName)
	// fmt.Println(getFirstNames(bookings))
	// firstNames := getFirstNames(bookings)
	// fmt.Printf("The first names are: %v\n", firstNames)

}

// func greetUsers() {
// 	fmt.Println("Welcome to our application")
// }

// func greetUsers2(conName string) {
// 	fmt.Printf("Welcome to %v booking app\n", conName)
// }

// return type []string specifies that this function returns a slice of strings
// func getFirstNames(bookings []string)  []string {
// 	var firstNames []string
// 	for _, booking := range bookings {
// 		var names = strings.Fields(booking)
// 		firstNames = append(firstNames, names[0])
// 	}
// 	return firstNames
// }
