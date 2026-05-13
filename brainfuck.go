package main

import (
	"fmt"
	"os"
	"strings"
)

func loadFile(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "")
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Uso errado")
		return
	}

	var memorie [1000]uint8
	pc := 0
	pivot := 0
	instructions := loadFile(args[1])
	opens_brackets := 0
	for {
		if pc == len(instructions) {
			break
		}

		current_istruction := instructions[pc]

		switch current_istruction {
		case "+":
			memorie[pivot] += 1
		case "-":
			memorie[pivot] -= 1
		case ">":
			pivot += 1
		case "<":
			pivot -= 1
		case ".":
			fmt.Printf("%c", memorie[pivot])
		case "[":
			if memorie[pivot] == 0 {
				pc += 1
				for {
					temp := instructions[pc]
					if temp == "[" {
						opens_brackets += 1
					} else if temp == "]" {
						if opens_brackets > 0 {
							opens_brackets -= 1
						} else {
							break
						}
					}
					pc += 1
				}
			}
		case "]":
			if memorie[pivot] != 0 {
				pc -= 1
				for {
					temp := instructions[pc]
					if temp == "]" {
						opens_brackets += 1
					} else if temp == "[" {
						if opens_brackets > 0 {
							opens_brackets -= 1
						} else {
							break
						}
					}
					pc -= 1
				}
			}
		}

		pc += 1
	}
}
