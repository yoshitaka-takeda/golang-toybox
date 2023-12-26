package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Input text: ")

	// The Scan method
	var w1, w2, w3 string
	n, err := fmt.Scan(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read text: %s %s %s-\n", w1, w2, w3)
}
