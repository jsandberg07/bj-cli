package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		if input == "" {
			fmt.Println("No input found")
			continue
		}
		return strings.TrimSpace(input), nil
	}
}
