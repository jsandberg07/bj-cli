package main

import (
	"strings"
)

// makes the input the length requested by adding spaces to the front and back
func centerPad(input string, width int) string {
	padding := width - len(input)
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + input + strings.Repeat(" ", right)
}

// helper function that takes arrays created by players and formats them into one string with breaks
func printPlayers(s [][]string) string {
	output := ""
	for j := 0; j < 2; j++ {
		for i := 0; i < len(s); i++ {
			output += "|" + s[i][j]
		}
		output += "|\n"
	}
	return output
}
