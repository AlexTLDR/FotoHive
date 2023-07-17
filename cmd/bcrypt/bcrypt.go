package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		// hash pass
		hash(os.Args[2])
	case "compare":
		// compare pass hash
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(pass string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", pass)
		return
	}
	fmt.Println(string(hashedBytes))
}

func compare(pass, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		fmt.Printf("Password is invalid: %v\n", err)
		return
	}
	fmt.Println("Password is correct!")
}
