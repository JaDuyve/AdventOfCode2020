package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var min, max int
	var password string
	var character uint8
	var numberOfValidPasswords int

	_, inputErr := fmt.Fscanf(f, "%d-%d %c: %s", &min, &max, &character, &password)
	for inputErr == nil {
		if isValid(min, max, character, password) {
			numberOfValidPasswords++
		}

		_, inputErr = fmt.Fscanf(f, "%d-%d %c: %s", &min, &max, &character, &password)
	}

	fmt.Printf("There are [%d] valid passwords.\n", numberOfValidPasswords)
}

func isValid(min int, max int, character uint8, password string) bool {
	return (password[min - 1] == character) != (password[max - 1] == character)
}