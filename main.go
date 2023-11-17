package main

import (
	"fmt"
	"os"
	"slices"
)

var cursor uint16 = 0
var memory [30000]byte

var commandWhiteList = [...]string{
	">", "<", "+", "-", ".", ",", "[", "]",
}

func getCode(path string) string {
	data, err := os.ReadFile(path)
	code := ""

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, command := range data {
		if slices.Contains(commandWhiteList[:], string(command)) {
			code += string(command)
		} else {
			continue
		}
	}

	return code
}

func main() {
	var path string = "main.bf"

	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	var code string = getCode(path)

	for command := 0; command < len(code); {
		switch string(code[command]) {
		case ">":
			cursor++
		case "<":
			cursor--
		case "+":
			memory[cursor]++
		case "-":
			memory[cursor]--
		case ".":
			fmt.Printf(string(memory[cursor]))
		case ",":
			var char string
			fmt.Scan(&char)

			memory[cursor] = char[0]
		case "[":
			if memory[cursor] == 0 {
				for i := 1; i > 0; {
					command++

					switch string(code[command]) {
					case "[":
						i++
					case "]":
						i--
					}
				}
			}
		case "]":
			if memory[cursor] != 0 {
				for i := 1; i > 0; {
					command--

					switch string(code[command]) {
					case "[":
						i--
					case "]":
						i++
					}
				}
			}
		}
		command++
	}
}
