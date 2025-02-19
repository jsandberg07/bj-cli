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

// so outer is the players
// inner is their name and then hand
// we want to print all the names, then all the hands
// so we actually iterate inner first
// then outer wait that's how that works
// goofed it all up so lets write a test
// do we have it as one big string yeah why the fuck not
func printPlayers(beans [][]string) string {
	output := ""
	for j := 0; j < 2; j++ {
		for i := 0; i < len(beans); i++ {
			output += "|" + beans[i][j]
		}
		output += "|\n"
	}
	return output
}
