package main

import (
	"fmt"
	"helper" // import helper functions from helper.go
)

func main() {
	// welcome message
	fmt.Println("\n######################################\nWelcome to hetu checker!\n######################################")
	for { // loops until correct hetu is received
		// User types their hetu
		fmt.Println("\nGive your hetu: ")
		// user input -> variable called hetu
		var hetu string
		fmt.Scan(&hetu)

		if !helper.IsHetuValid(hetu) { // if hetu is invalid -> print message
			fmt.Println("Invalid input! Please enter a valid personal identity code \"ddmmyyxnnnt\".")
			continue // continue loop if invalid hetu
		}
		birthDate := helper.GetBirthDate(hetu)                   //get birthday
		gender := helper.GetGender(hetu)                         // get gender
		fmt.Println("Valid hetu.\nYour birthday is:", birthDate) // print bithrday
		fmt.Println("Your identifier sign is", gender)           // print indentifier sign
		break                                                    // break loop when correct hetu is received
	} // for
} // main()
