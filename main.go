package main

import "fmt"
import "strings"

var (
	author  = "Evan Rose"
	Version = "0.0.1"
)

func main() {
	fmt.Println(double(5))
	first, _ := parseName("Evan Rose")
	first_non_idiomatic, _ := parseNameNonIdiomatic("Evan Rose")
	fmt.Println(first)
	fmt.Println(first_non_idiomatic)

	//anonymous function:
	greet := func(name string) {
		fmt.Println("Hello", name, ", how are you today?")
	}

	greet(first)
}

func double(n int) int {
	return n + n
}

func parseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}

//naming the outputs is optional

func parseNameNonIdiomatic(name string) (string, string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}
