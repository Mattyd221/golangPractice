package stringUtils

import "strings"

func LeftPad(startString string, padString string, padLength int) string {
	return strings.Repeat(padString, padLength) + startString
}

func RightPad(startString string, padString string, padLength int) string {
	return startString + strings.Repeat(padString, padLength)
}

func LeftPadToLength(startString string, padString string, overallLength int) string {
	var padCount int
	padCount = 1 + ((overallLength - len(padString))/ len(padString))
	ans := strings.Repeat(padString, padCount) + startString
	return ans[(len(ans) - overallLength):]
}

func RightPadToLength(startString string, padString string, overallLength int) string {
	var padCount int
	padCount = 1 + ((overallLength - len(padString))/ len(padString))
	ans :=  startString + strings.Repeat(padString, padCount)
	return ans[:overallLength]
}