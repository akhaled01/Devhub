package utils

import "fmt"

// THIS IS A DEBUGGING METHOD
// IN THE CASE PANIC HAPPENS. 
// IT ALLOWS US TO RECOVER FROM PANICS
// AND UNDERSTAND THE FAULT.
// DO NOT PUSH INTO PRODUCTION!
func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
		// Additional recovery logic can go here
	}
}
