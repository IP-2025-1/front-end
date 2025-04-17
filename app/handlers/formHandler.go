package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	util "front-end/app/utils"
)

func FormHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
		return
	}

	username := request.FormValue("username")

	ticketid := util.GenerateTicketID(7)

	fmt.Fprintf(response, "Hello "+username+"\n")
	eventname := request.FormValue("eventname")
	eventlocation := request.FormValue("eventlocation")
	ticketprice := request.FormValue("ticketprice")
	ticketquantity := request.FormValue("ticketquantity")

	price, err := strconv.ParseFloat(ticketprice, 64)
	if err != nil {
		fmt.Fprintf(response, "Error converting ticket price: %v", err)
		return
	}

	// Convert ticketquantity to int
	quantity, err := strconv.Atoi(ticketquantity)
	if err != nil {
		fmt.Fprintf(response, "Error converting ticket quantity: %v", err)
		return
	}

	// Calculate total cost
	totalCost := price * float64(quantity)

	fmt.Fprintf(response, "Event Name = %s\n", eventname)
	fmt.Fprintf(response, "Your Ticket ID is %s\n", ticketid)
	fmt.Fprintf(response, "You're off to %s\n", eventlocation)
	fmt.Fprintf(response, "Ticket Price = %s\n", ticketprice)
	fmt.Fprintf(response, "Ticket Quantity = %s\n", ticketquantity)
	fmt.Fprintf(response, "Total Cost = %.2f\n", totalCost)

}
