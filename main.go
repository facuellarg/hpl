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
	reducible = map[rune]bool{
		'👈': true,
		'👉': true,
		'👆': true,
		'👇': true,
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
func exec(input []rune, repetitions []int, index int, memory []byte, indexMemory int, output *[]byte, pairPositions []int) (int, int) {
	cmd := input[index]
	switch cmd {
	case '👉':

		indexMemory += repetitions[index]
	case '👈':
		indexMemory -= repetitions[index]
	case '👆':
		memory[indexMemory] = byte(addWithModule(int(memory[indexMemory]), repetitions[index], Module))
	case '👇':
		memory[indexMemory] = byte(addWithModule(int(memory[indexMemory]), -1*repetitions[index], Module))
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

func reduce(input []rune, reducible map[rune]bool) ([]rune, []int) {
	repetitions := make([]int, 0, len(input))
	reduceInput := make([]rune, 0, len(input))
	for i := 0; i < len(input); i++ {
		reduceInput = append(reduceInput, input[i])
		repetitions = append(repetitions, 1)
		if reducible[input[i]] {
			repetition := 0
			initial := input[i]
			current := input[i+1]
			for current == initial {
				repetition++
				current = input[i+repetition+1]
			}
			i += repetition
			repetitions[len(repetitions)-1] += repetition

		}

	}
	return reduceInput, repetitions

}

// hpl executes the HPL program
func hpl(input []rune) string {
	var indexMemory int
	memory := make([]byte, 1000)
	output := make([]byte, 0)
	// helper memory to store the position of the pairs
	prepro, rep := reduce(input, reducible)
	positions := make([]int, len(prepro))
	for i := 0; i < len(positions); i++ {
		positions[i] = -1
	}

	for index := 0; index < len(prepro); index++ {
		index, indexMemory = exec(prepro, rep, index, memory, indexMemory, &output, positions)
	}
	return string(output)
}
