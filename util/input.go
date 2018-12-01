package util

import (
	"bufio"
	"log"
	"os"
)

func ReadInputFromFile(filename string) []string {
	input := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}