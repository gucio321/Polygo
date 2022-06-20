package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/garnn/Polygo/pkg/factorial"
)

func main() {
	var text string
	fmt.Println("enter number: ")
	fmt.Scanln(&text)
	number, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("error parsing text: %v", err)
	}

	fmt.Printf("Factorial of %d is %d\n", number, factorial.Factorial(number))
}
