package main

func validateUserInput(username string, buyTicket uint) (bool, bool) {
	validName := len(username) >= 2
	validTicket := buyTicket > 0 && buyTicket <= remainingTickets
	return validName, validTicket
}
