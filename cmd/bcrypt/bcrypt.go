package main

import (
	"fmt"
	"os"
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
	fmt.Printf("Hashing password: %q\n", pass)
}

func compare(pass, hash string) {
	fmt.Printf("Comparing password: %q with the hash %q\n", pass, hash)
}
