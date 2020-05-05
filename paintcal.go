package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}

func main() {
	var amount, total float64
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Println(err)
	}
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)
	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		log.Println(err)
	}
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)
	amount, err = paintNeeded(5.0, -3.3)
	if err != nil {
		log.Println(err)
	}
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)

	fmt.Printf("%.2f total liters needed\n", total)
}
