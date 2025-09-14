package main

import "fmt"

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
	var bookings [50]string //array declaration

	// dynamic slice
	// var bookings []string
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
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array/slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The array/slice type: %T\n", bookings)
	fmt.Printf("The length of array/slice: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// fmt.Printf("All bookings: %v\n", bookings)
}
