package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const memorySize = 30000

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Use: bf path/to/file.bf")
	}

	filename := os.Args[1]
	code, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error al leer el archivo: %v", err)
	}

	interpret(string(code))
}

func interpret(code string) {
	mem := make([]byte, memorySize)
	ptr := 0
	pc := 0 // program counter
	stack := []int{}
	input := bufio.NewReader(os.Stdin)

	for pc < len(code) {
		switch code[pc] {
		case '>':
			ptr++
			if ptr >= memorySize {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = memorySize - 1
			}
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			fmt.Printf("%c", mem[ptr])
		case ',':
			b, _ := input.ReadByte()
			mem[ptr] = b
		case '[':
			if mem[ptr] == 0 {
				// Saltar al final del bucle
				open := 1
				for open > 0 {
					pc++
					if pc >= len(code) {
						log.Fatal("Error: bucle '[' sin su correspondiente ']'")
					}
					if code[pc] == '[' {
						open++
					} else if code[pc] == ']' {
						open--
					}
				}
			} else {
				stack = append(stack, pc)
			}
		case ']':
			if mem[ptr] != 0 {
				if len(stack) == 0 {
					log.Fatal("Error: bucle ']' sin su correspondiente '['")
				}
				pc = stack[len(stack)-1] - 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		pc++
	}
}
