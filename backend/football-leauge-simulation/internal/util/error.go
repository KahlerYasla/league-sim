package util

import (
	"log"
)

// CheckError is a helper function to check and log errors.
func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// HandleError is a helper function to handle errors and log them without stopping the program.
func HandleError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
