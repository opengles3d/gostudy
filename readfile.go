package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if(err != nil) {
		log.Fatal(err)
	}

	scanner:= bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	err = file.Close()
	if(err != nil) {
		log.Fatal(err)
	}

	if(scanner.Err() != nil) {
		log.Fatal(scanner.Err())
	}
}