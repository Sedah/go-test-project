package main

import (
	"fmt"
	"sync"
	"time"
)

var projectName = "My first GO project"

const projectTickets = 21

var remainingTickets uint = 21
var bookings = make([]UserdData, 0)

type UserdData struct {
	username  string
	buyTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		username, buyTicket := getUserInput()
		validName, validTicket := validateUserInput(username, buyTicket)
		if validName && validTicket {

			if buyTicket <= remainingTickets {
				bookTicket(buyTicket, username)
				wg.Add(1)
				go bookingSuccess(buyTicket, username)
				user_list := []string{}

				for _, booking := range bookings {
					user_list = append(user_list, booking.username)
				}
				if remainingTickets == 0 {
					break
				}

				fmt.Printf("List of member in this world %v\n", user_list)
				fmt.Printf("Total of tickets %v,ticket remaining %v\n", projectTickets, remainingTickets)
			} else {
				fmt.Printf("Only %v tickets left, you can't buy %v tickets\n", remainingTickets, buyTicket)
			}

		} else {
			if !validName {
				fmt.Println("username should more than 2 char")
			} else {
				fmt.Println("ticket should more than 0 and less than remaining ticket ")
			}
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v\n", projectName)
	fmt.Printf("Total of tickets %v,ticket remaining %v\n", projectTickets, remainingTickets)
	fmt.Println("Get your ticket for visit my project")
}

func getUserInput() (string, uint) {
	var username string
	var buyTicket uint
	fmt.Println("Enter your name")
	fmt.Scan(&username)
	fmt.Println("Enter number of ticket")
	fmt.Scan(&buyTicket)
	return username, buyTicket
}

func bookTicket(buyTicket uint, username string) {
	var userData = UserdData{
		username:  username,
		buyTicket: buyTicket,
	}
	remainingTickets = remainingTickets - buyTicket
	bookings = append(bookings, userData)
}

func bookingSuccess(buyTicket uint, username string) {
	time.Sleep(5 * time.Second)
	fmt.Printf("list of bookings %v\n", bookings)
	fmt.Printf("Welcome to my world! %v, you buy %v ticket from us\n", username, buyTicket)
	fmt.Printf("Tickets available %v\n", remainingTickets)
	wg.Done()
}
