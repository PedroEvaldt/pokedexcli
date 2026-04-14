package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return words
}
