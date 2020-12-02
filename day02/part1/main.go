package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var min, max int
	var character, password string
	var numberOfValidPasswords int

	_, inputErr := fmt.Fscanf(f, "%d-%d %1s: %s", &min, &max, &character, &password)
	for inputErr == nil {
		if isValid(min, max, character, password) {
			numberOfValidPasswords++
		}

		_, inputErr = fmt.Fscanf(f, "%d-%d %1s: %s", &min, &max, &character, &password)
	}

	fmt.Printf("There are [%d] valid passwords.\n", numberOfValidPasswords)
}

func isValid(min int, max int, character string, password string) bool {
	count := strings.Count(password, character)
	return count >= min && count <= max
}