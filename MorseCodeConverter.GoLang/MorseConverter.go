package main

import (
	"bufio"
	"os"
	"strings"
)

type IMorseConverter interface {
	tryParseStringToMorse(input string) string
	tryParseCharacterToMorse(input rune) string
}

type ISOBasicLatinConverter struct {
	characterMap map[rune]string
}

func(c ISOBasicLatinConverter) tryParseStringToMorse(input string) string {
	//Default to uppercase
	input = strings.ToUpper(input)

	//Split our string into words, on the assumption each word is single-space separated
	words := strings.Split(input, " ")

	builtWord := ""

	for _, word := range words {

		runes := []rune(word)

		for _, character := range runes {
			parsedMorse := c.tryParseCharacterToMorse(character)

			//println("Parsed morse: " + parsedMorse)

			builtWord = strings.Join([]string { builtWord, parsedMorse }, " ")
		}

		builtWord = strings.Join([]string { builtWord, "   "}, "")

		//println(builtWord)
	}

	//println("Final: " + builtWord)

	return builtWord
}

func(c ISOBasicLatinConverter) tryParseCharacterToMorse(input rune) string {

	_, ok := c.characterMap[input]

	if ok {
		result := c.characterMap[input]
		//println("Character map result: " + result)
		return result
	} else if input == '\n' {
		return ""
	} else {
		panic("Invalid character specified " + string(input))
	}
}

func NewISOBasicLatinConverter() ISOBasicLatinConverter {

	dot := "."
	dash := "-"

	converter := ISOBasicLatinConverter{}

	converter.characterMap = map[rune]string {
		'A': strings.Join([]string {dot, dash}, ""),
		'B': strings.Join([]string {dash, dot, dot, dot}, ""),
		'C': strings.Join([]string {dash, dot, dash, dot}, ""),
		'D': strings.Join([]string {dash, dot, dot}, ""),
		'E': dot,
		'F': strings.Join([]string {dot, dot, dash, dot}, ""),
		'G': strings.Join([]string {dash, dash, dot}, ""),
		'H': strings.Join([]string {dot, dot, dot, dot}, ""),
		'I': strings.Join([]string {dot, dot}, ""),
		'J': strings.Join([]string {dot, dash, dash, dash}, ""),
		'K': strings.Join([]string {dash, dot, dash}, ""),
		'L': strings.Join([]string {dot, dash, dot, dot}, ""),
		'M': strings.Join([]string {dash, dash}, ""),
		'N': strings.Join([]string {dash, dot}, ""),
		'O': strings.Join([]string {dash, dash, dash}, ""),
		'P': strings.Join([]string {dot, dash, dash, dot}, ""),
		'Q': strings.Join([]string {dash, dash, dot, dash}, ""),
		'R': strings.Join([]string {dot, dash, dot}, ""),
		'S': strings.Join([]string {dot, dot, dot}, ""),
		'T': strings.Join([]string {dash}, ""),
		'U': strings.Join([]string {dot, dot, dash}, ""),
		'V': strings.Join([]string {dot, dot, dot, dash}, ""),
		'W': strings.Join([]string {dot, dash, dash}, ""),
		'X': strings.Join([]string {dash, dot, dot, dash}, ""),
		'Y': strings.Join([]string {dash, dot, dash, dash}, ""),
		'Z': strings.Join([]string {dash, dash, dot, dot}, ""),

		'0': strings.Join([]string {dash, dash, dash, dash, dash}, ""),
		'1': strings.Join([]string {dot, dash, dash, dash, dash}, ""),
		'2': strings.Join([]string {dot, dot, dash, dash, dash}, ""),
		'3': strings.Join([]string {dot, dot, dot, dash, dash}, ""),
		'4': strings.Join([]string {dot, dot, dot, dot, dash}, ""),
		'5': strings.Join([]string {dot, dot, dot, dot, dot}, ""),
		'6': strings.Join([]string {dash, dot, dot, dot, dot}, ""),
		'7': strings.Join([]string {dash, dash, dot, dot, dot}, ""),
		'8': strings.Join([]string {dash, dash, dash, dot, dot}, ""),
		'9': strings.Join([]string {dash, dash, dash, dash, dot}, ""),

		'.': strings.Join([]string {dot, dash, dot, dash, dot, dash}, ""),
		'&': strings.Join([]string {dot, dash, dot, dot, dot}, ""),
		//'\', strings.Join([]string {dot, dash, dash, dash, dash, dot}, ""),
		'@': strings.Join([]string {dot, dash, dash, dot, dash, dash}, ""),
		'(': strings.Join([]string {dash, dot, dash, dash, dot }, ""),
		')': strings.Join([]string {dash, dot, dash, dash, dot, dash}, ""),
		':': strings.Join([]string {dash, dash, dot, dot, dash, dash}, ""),
		'=': strings.Join([]string {dash, dot, dot, dot, dash}, ""),
		'-': strings.Join([]string {dash, dot, dot, dot, dot, dash}, ""),
		'+': strings.Join([]string {dot, dash, dot, dash, dot}, ""),
		'"': strings.Join([]string {dot, dash, dot, dot, dash, dot}, ""),
		'?': strings.Join([]string {dot, dot, dash, dash, dot, dot}, ""),
		'/': strings.Join([]string {dash, dot, dot, dash, dot}, ""),
	}

	return converter
}

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
