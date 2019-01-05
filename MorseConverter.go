package main

import (
	"bufio"
	"os"
)

func main() {
	converter := NewISOBasicLatinConverter()

	println("Welcome to the English > Morse Code translator.")
	println("")
	println("What text do you want to convert? Type it then press enter to see the result:")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	result := converter.tryParseStringToMorse(userInput)

	println(result)
}
