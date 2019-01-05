package Interfaces

type IMorseConverter interface {
	tryParseStringToMorse(input string) string
	tryParseCharacterToMorse(input rune) string
}
