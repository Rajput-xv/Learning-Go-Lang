package main

import (
	"fmt"
	"strings"
	// "booking-app/helper"
	// "time"
	// "sync"
)

//package level global varibles
// const conferenceTickets int = 50
// var remainingTickets int = 50
// var conferenceName = "Go Conference"

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

// assign value in main function - isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
//multiple return values
// func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

// 	return isValidName, isValidEmail, isValidTicketNumber
// }

// assign value in main function - firstName, lastName, email, userName, userTickets := getUserInput()
//taking user input in function
// func getUserInput() (string, string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userName string
// 	var userTickets uint

// 	fmt.Println("Enter your First name: ")
// 	fmt.Scan(&firstName) //& is used to get the memory address of the variable

// 	fmt.Println("Enter your Last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter your Email: ")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter your User name: ")
// 	fmt.Scan(&userName)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	return firstName, lastName, email, userName, userTickets
// }

// assign value in main function - bookings, remainingTickets := bookTicket( userTickets, firstName, lastName, email, userName, bookings, remainingTickets)
//updating multiple values in function
// func bookTicket(userTickets uint, firstName string, lastName string, email string, userName string) (uint, []string) {
// 	remainingTickets = remainingTickets - userTickets
// 	bookings = append(bookings, fmt.Sprintf("%v %v %v %v %v", firstName, lastName, email, userName, userTickets))
// 	return remainingTickets, bookings
// }

//map - string to string (only 1 data type)
// var userData = map[string]string{}
// var userData = make(map[string]string)

// userData["firstName"] = firstName
// userData["lastName"] = lastName
// userData["email"] = email
// userData["userName"] = userName
// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) //convert uint to string

// initialize slice of map with 0 size
// var bookings = make([]map[string]string, 0)
// bookings = append(bookings, userData)

// mix data types - structure
// package level global varibles
// type UserData struct{
// 	firsName string
// 	lastName string
// 	email string
// 	userName string
// 	userTickets uint
// }

// assign value
// var UserData = UserData{
// 	firsName:    "" or variable name,
// 	lastName:    "", or variable name,
// 	email:       "", or variable name,
// 	userName:    "", or variable name,
// 	userTickets: 0, or variable name,
// }

// access values by .
// firstNames := []string{}
// firstNames = append(firstNames, UserData.firsName)

// func sendTicket(userTickets uint, firstName string, lastName string, email string) {
// 	time.Sleep(10 * time.Second)
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
// 	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
// }

// concurrency
// new thread
// while calling the function add go in front of the line

// goroutines means new threads
// to continue main thread until the goroutine is completed
// var wg = sync.WaitGroup{}

// wg.Add(1) //add the number of goroutines to wait for (before the new thread)
// go sendTicket(userTickets, firstName, lastName, email) (after the function/thread call)
// wg.Wait() //block until the counter is 0 (last line of function code)
// wg.Done() //decrement the counter by 1 (at the end of the new thread/function)
