package main

import (
	"bufio"
	"os"
)

var (
	Module = 256
	pairs  = map[rune]rune{
		'🤜': '🤛',
		'🤛': '🤜',
	}
)

func main() {

	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	sc.Scan()
	input := []rune(sc.Text())

	println(hpl(input))

}

/*
👉 : moves the memory pointer to the next cell

👈 : moves the memory pointer to the previous cell

👆 : increment the memory cell at the current position

👇 : decreases the memory cell at the current position.

🤜 : if the memory cell at the current position is 0, jump just after the corresponding 🤛

🤛 : if the memory cell at the current position is not 0, jump just after the corresponding 🤜

👊 : Display the current character represented by the ASCII code defined by the current position.
*/

// addWithModule adds two numbers with the module
func addWithModule(value, valueToAdd int, module int) int {
	return (value + valueToAdd + module) % module
}

func lookUpPair(input []rune, index, step int) int {
	initial := input[index]
	count := 1
	for count != 0 {
		index = index + step
		cmd := rune(input[index])
		if cmd == initial {
			count++
		} else if cmd == pairs[initial] {
			count--
		}
	}
	return index
}

// exec executes the command at the current index and returns the new index and the memory index
func exec(input []rune, index int, memory []byte, indexMemory int, output *[]byte, pairPositions []int) (int, int) {
	cmd := input[index]
	switch cmd {
	case '👉':

		indexMemory++
	case '👈':
		indexMemory--
	case '👆':
		memory[indexMemory] = byte(addWithModule(int(memory[indexMemory]), 1, Module))
	case '👇':
		memory[indexMemory] = byte(addWithModule(int(memory[indexMemory]), -1, Module))
	case '🤜':
		if memory[indexMemory] == 0 {
			if pairPositions[index] == -1 {
				pairPositions[index] = lookUpPair(input, index, 1)
			}
			index = pairPositions[index]
		}
	case '🤛':
		if memory[indexMemory] != 0 {
			if pairPositions[index] == -1 {
				pairPositions[index] = lookUpPair(input, index, -1)
			}
			index = pairPositions[index]
		}
	case '👊':
		*output = append(*output, memory[indexMemory])
	}
	return index, indexMemory

}

// hpl executes the HPL program
func hpl(input []rune) string {
	var indexMemory int
	memory := make([]byte, 1000)
	output := make([]byte, 0)
	// helper memory to store the position of the pairs
	positions := make([]int, len(input))
	for i := 0; i < len(positions); i++ {
		positions[i] = -1
	}

	for index := 0; index < len(input); index++ {
		index, indexMemory = exec(input, index, memory, indexMemory, &output, positions)
	}
	return string(output)
}
