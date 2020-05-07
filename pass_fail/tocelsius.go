package main

import (
	"fmt"
	"log"
	"shaocq"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit:")
	fahrenheit, err := shaocq.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
