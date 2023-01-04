package main

import (
	"fmt"
	"strings"
)

var (
	score    int    = 0
	input    string = ""
	errors          = map[string]string{}
	alphabet        = map[string]string{
		"A": "alpha",
		"B": "bravo",
		"C": "charlie",
		"D": "delta",
		"E": "echo",
		"F": "foxtrot",
		"G": "golf",
		"H": "hotel",
		"I": "india",
		"J": "juliet",
		"K": "kilo",
		"L": "lima",
		"M": "mike",
		"N": "november",
		"O": "oscar",
		"P": "papa",
		"Q": "quebec",
		"R": "romeo",
		"S": "sierra",
		"T": "tango",
		"U": "uniform",
		"V": "victor",
		"W": "whiskey",
		"X": "x-ray",
		"Y": "yankee",
		"Z": "zulu",
	}
)

func main() {
	fmt.Println("Game is starting, type \"stop\" to exit")
loop:
	for k, v := range alphabet {
		fmt.Printf("%v: ", k)
		fmt.Scan(&input)
		input = strings.ToLower(input)

		switch input {
		case v:
			score++
		case "stop", "exit":
			break loop
		default:
			errors[k] = input

		}

	}
	fmt.Printf("\nYour score is %v/%v\n", score, len(alphabet))

	if len(errors) > 0 {
		fmt.Println("\nYour errors:")
		for k, wrong := range errors {
			fmt.Printf("%v: %v --> %v\n", k, wrong, alphabet[k])
		}
	}

}
